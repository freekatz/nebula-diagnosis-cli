package tests

import (
	"log"
	"testing"
	"time"

	"github.com/nebula/nebula-diagnose/pkg/config"
	"github.com/nebula/nebula-diagnose/pkg/remote"
)

func TestGetSSHClient(t *testing.T) {
	conf := &config.SSHConfig{
		Address:  "192.168.8.49",
		Port:     22,
		Timeout:  "3s",
		Username: "katz.zhang",
		Password: "nebula",
	}
	go testSSHClient(conf)
	go testSSHClient1(conf)
	testSSHClient2(conf)
	time.Sleep(3 * time.Second)
}

func testSSHClient(conf *config.SSHConfig) {
	c, err := remote.GetSSHClient(conf.Username, conf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", c)

	ch := make(chan remote.ExecuteResult)
	go c.ExecuteAsync("vmstat 1 1", ch)
	// sudo du -sh /home/*
	// df -H | grep -vE '^Filesystem|tmpfs|udev' | awk '{ print $1 " " $2 " " $3 " " $4 " " $5 }'
	for {
		select {
		case res := <-ch:
			log.Println("\n", string(res.StdOut))
			break
		default:
		}
	}
}

func testSSHClient1(conf *config.SSHConfig) {
	c1, err := remote.GetSSHClient(conf.Username, conf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", c1)

	ch1 := make(chan remote.ExecuteResult)
	go c1.ExecuteAsync("ls", ch1)

	for {
		select {
		case res := <-ch1:
			log.Println(string(res.StdOut))
			break
		default:

		}
	}
}

func testSSHClient2(conf *config.SSHConfig) {
	c2, err := remote.GetSSHClient(conf.Username, conf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", c2)

	res2, _ := c2.Execute("ls")

	log.Println(string(res2.StdOut))
}
