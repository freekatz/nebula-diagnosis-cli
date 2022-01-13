package packer

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type TgzPacker struct {
}

func NewTgzPacker() *TgzPacker {
	return &TgzPacker{}
}

// removeTargetFile remove target file
func (tp *TgzPacker) removeTargetFile(fileName string) (err error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return nil
	}
	return os.Remove(fileName)
}

// dirExists is dir existed
func (tp *TgzPacker) dirExists(dir string) bool {
	info, err := os.Stat(dir)
	return (err == nil || os.IsExist(err)) && info.IsDir()
}

// tarFile sourceFullFile is the source path of file
func (tp *TgzPacker) tarFile(sourceFullFile string, writer *tar.Writer) error {
	info, err := os.Stat(sourceFullFile)
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}

	err = writer.WriteHeader(header)
	if err != nil {
		return err
	}

	fr, err := os.Open(sourceFullFile)
	if err != nil {
		return err
	}
	defer func() {
		if err2 := fr.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()
	if _, err = io.Copy(writer, fr); err != nil {
		return err
	}
	return nil
}

// tarFolder sourceFullPath is the source path of folder, baseName is the base name of tar file
func (tp *TgzPacker) tarFolder(writer *tar.Writer, sourceFullPath string, baseName string) error {
	baseFullPath := sourceFullPath
	return filepath.Walk(sourceFullPath, func(fileName string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		if fileName == baseFullPath {
			header.Name = baseName
		} else {
			header.Name = filepath.Join(baseName, strings.TrimPrefix(fileName, baseFullPath))
		}
		if err = writer.WriteHeader(header); err != nil {
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
		if _, err := io.Copy(writer, fr); err != nil {
			return err
		}
		return nil
	})
}

// Pack sourceFullPath is the source path of file or folder, tarFileName is the tar package
func (tp *TgzPacker) Pack(sourceFullPath string, tarFileName string) (err error) {
	sourceInfo, err := os.Stat(sourceFullPath)
	if err != nil {
		return err
	}

	if err = tp.removeTargetFile(tarFileName); err != nil {
		return err
	}

	file, err := os.Create(tarFileName)
	if err != nil {
		return err
	}
	defer func() {
		if err2 := file.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()

	gWriter := gzip.NewWriter(file)
	defer func() {
		if err2 := gWriter.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()

	tarWriter := tar.NewWriter(gWriter)
	defer func() {
		if err2 := tarWriter.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()

	if sourceInfo.IsDir() {
		return tp.tarFolder(tarWriter, sourceFullPath, filepath.Base(sourceFullPath))
	}
	return tp.tarFile(sourceFullPath, tarWriter)
}

// UnPack tarFileName is the tar package，dstDir is the dst dir path
func (tp *TgzPacker) UnPack(tarFileName string, dstDir string) (err error) {
	fr, err := os.Open(tarFileName)
	if err != nil {
		return err
	}
	defer func() {
		if err2 := fr.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()

	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer func() {
		if err2 := gr.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()

	tarReader := tar.NewReader(gr)
	for {
		header, err := tarReader.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}

		targetFullPath := filepath.Join(dstDir, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if exists := tp.dirExists(targetFullPath); !exists {
				if err = os.MkdirAll(targetFullPath, 0755); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			file, err := os.OpenFile(targetFullPath, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			_, err = io.Copy(file, tarReader)
			if err2 := file.Close(); err2 != nil {
				return err2
			}
			if err != nil {
				return err
			}
		}
	}
}
