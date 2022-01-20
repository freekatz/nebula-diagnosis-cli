package packer

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/utils"
)

type TgzPacker struct {
}

func NewTgzPacker() *TgzPacker {
	return &TgzPacker{}
}

// tarFile write the file into tar.writer
func (tp *TgzPacker) tarFile(tarWriter *tar.Writer, tarFilepath string) error {
	info, err := os.Stat(tarFilepath)
	if err != nil {
		return errorx.ErrFileNotExisted
	}
	if info.IsDir() {
		return errorx.ErrFileTypeNotMatch
	}

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}

	err = tarWriter.WriteHeader(header)
	if err != nil {
		return err
	}

	fr, err := os.Open(tarFilepath)
	if err != nil {
		return err
	}
	defer func() {
		if e := fr.Close(); e != nil && err == nil {
			err = e
		}
	}()
	if _, err = io.Copy(tarWriter, fr); err != nil {
		return err
	}
	return nil
}

// tarFolder sourceFullPath is the source path of folder, baseName is the base name of tar file
func (tp *TgzPacker) tarFolder(tarWriter *tar.Writer, tarFilepath string, baseFilename string) error {
	baseFilepath := tarFilepath
	return filepath.Walk(tarFilepath, func(fileName string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		if fileName == baseFilepath {
			header.Name = baseFilename
		} else {
			header.Name = filepath.Join(baseFilename, strings.TrimPrefix(fileName, baseFilepath))
		}

		log.Println(header.Name)
		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer fr.Close()
		if _, err := io.Copy(tarWriter, fr); err != nil {
			return err
		}
		return nil
	})
}

// Pack sourceFullPath is the source path of file or folder, tarFileName is the tar package
func (tp *TgzPacker) Pack(tarFilepath string, targetFilepath string) (err error) {
	tarFilepath = strings.TrimPrefix(tarFilepath, "./")
	targetFilepath = strings.TrimPrefix(targetFilepath, "./")

	info, err := os.Stat(tarFilepath)
	if err != nil {
		return err
	}

	if utils.IsFileExisted(targetFilepath) {
		return errorx.ErrFileHasExisted
	}

	file, err := os.Create(targetFilepath)
	if err != nil {
		return err
	}
	defer func() {
		if e := file.Close(); e != nil && err == nil {
			err = e
		}
	}()

	gzipWriter := gzip.NewWriter(file)
	defer func() {
		if e := gzipWriter.Close(); e != nil && err == nil {
			err = e
		}
	}()

	// TODO fix `archive/tar: write too long` error
	tarWriter := tar.NewWriter(gzipWriter)
	defer func() {
		if err2 := tarWriter.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()

	if info.IsDir() {
		return tp.tarFolder(tarWriter, tarFilepath, filepath.Base(tarFilepath))
	}
	return tp.tarFile(tarWriter, tarFilepath)
}
