package tests

import (
	"log"
	"testing"

	"github.com/nebula/nebula-diagnose/pkg/config"
	"github.com/nebula/nebula-diagnose/pkg/info/physical"
)

func TestGetInfo(t *testing.T) {
	conf := &config.SSHConfig{
		Address:  "192.168.8.49",
		Port:     22,
		Timeout:  "1s",
		Username: "katz.zhang",
		Password: "nebula",
	}
	info, _ := physical.GetPhyInfo(conf)
	log.Printf("%+v", info)
}
