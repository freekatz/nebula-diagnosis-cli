package remote

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/pkg/sftp"
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

func (c *SFTPClient) GetFileInRemotePath(remotePath string, localPath string) error {
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

func (c *SFTPClient) GetFilesInRemoteDir(remoteDir string, localDir string) error {
	p, _ := filepath.Abs(localDir)
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		err := os.Mkdir(p, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	filesInfo, err := c.Client.ReadDir(remoteDir)
	for _, fileInfo := range filesInfo {

		if fileInfo.IsDir() {
			continue
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

func (c *SFTPClient) Close() {
	sftpMux.Lock()
	defer sftpMux.Unlock()
	c.Client.Close()
	delete(sftpClients, c.scid)
}
