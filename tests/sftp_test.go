package tests

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/remote"
	"log"
	"testing"
)

func TestGetSFTPClient(t *testing.T) {
	path := "/home/katz.zhang/logs"
	localDir := "logs"
	conf := &config.SSHConfig{
		Address:  "192.168.8.49",
		Port:     22,
		Timeout:  "1s",
		Username: "katz.zhang",
		Password: "nebula",
	}

	client, err := remote.GetSFTPClient(conf.Username, conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	//err = client.GetFileInRemotePath(path, localDir)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	err = client.GetFilesInRemoteDir(path, localDir)
	if err != nil {
		log.Fatal(err.Error())
	}
}
