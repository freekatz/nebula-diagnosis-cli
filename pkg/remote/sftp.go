package remote

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/nebula/nebula-diagnose/pkg/config"
	"github.com/nebula/nebula-diagnose/pkg/utils"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SFTPClient struct {
	scid string
	*sftp.Client
}

var (
	sftpClients = make(map[string]*SFTPClient)
	sftpMux     sync.RWMutex
)

// GetSFTPClient TODO refactor the following methods
func GetSFTPClient(scid string, conf *config.SSHConfig) (*SFTPClient, error) {
	sftpMux.Lock()
	if _, ok := sftpClients[scid]; !ok {
		sshClient, err := GetSSHClient(scid, conf)
		if err != nil {
			return nil, err
		}

		c, err := newSFTPClient(sshClient.Client)
		if err != nil {
			return nil, err
		}

		c.scid = scid
		sftpClients[scid] = c
	}
	sftpMux.Unlock()

	sftpMux.RLock()
	c := sftpClients[scid]
	sftpMux.RUnlock()

	return c, nil
}

func newSFTPClient(sshClient *ssh.Client) (*SFTPClient, error) {
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, err
	}
	return &SFTPClient{Client: sftpClient}, nil
}

func (c *SFTPClient) DownloadFile(remotePath string, localPath string) error {
	src, err := c.Client.Open(remotePath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.OpenFile(localPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}

func (c *SFTPClient) DownloadDir(remoteDir string, localDir string) error {
	p, _ := filepath.Abs(localDir)
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return err
		}
	}

	filesInfo, err := c.Client.ReadDir(remoteDir)
	for _, fileInfo := range filesInfo {
		if fileInfo.IsDir() {
			dirName := fileInfo.Name()
			// remoteDir is under the Linux path by default
			remoteSubDir := remoteDir + "/" + dirName
			localSubDir := filepath.Join(localDir, dirName)
			c.DownloadDir(remoteSubDir, localSubDir)
		}

		srcPath := remoteDir + "/" + fileInfo.Name()
		src, err := c.Client.OpenFile(srcPath, os.O_RDONLY)
		if err != nil {
			return err
		}
		defer src.Close()

		dstPath := filepath.Join(localDir, fileInfo.Name())
		dst, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			continue
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			continue
		}
	}
	return nil
}

func (c *SFTPClient) UploadFile(remoteFilepath string, localFilepath string) bool {
	srcFile, err := os.Open(localFilepath)
	if err != nil {
		return false
	}
	defer srcFile.Close()

	remoteFilename := path.Base(localFilepath)
	dstFile, err := c.Client.Create(path.Join(remoteFilepath, remoteFilename))
	if err != nil {
		log.Fatal("remote file create error : ", path.Join(remoteFilepath, remoteFilename))
	}
	defer dstFile.Close()

	srcStat, _ := srcFile.Stat()
	err = utils.PrintWithProcessBar(srcStat.Size(), "uploading", srcFile, dstFile)
	if err != nil {
		return false
	}

	return true
}

func (c *SFTPClient) UploadFileAsync(remoteFilepath string, localFilepath string, ch chan<- bool) {
	ok := c.UploadFile(remoteFilepath, localFilepath)
	ch <- ok
}

func (c *SFTPClient) UploadDir() bool {
	return true
}

func (c *SFTPClient) UploadDirAsync(ch chan<- bool) {
	ok := c.UploadDir()
	ch <- ok
}

func (c *SFTPClient) Close() {
	sftpMux.Lock()
	defer sftpMux.Unlock()
	c.Client.Close()
	delete(sftpClients, c.scid)
}
