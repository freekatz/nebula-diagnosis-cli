package tests

import (
	"log"
	"testing"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
)

func TestNewInfoConfig(t *testing.T) {
	confPath := "../config/info_debug.yaml"

	conf, err := config.NewInfoConfig(confPath, "yaml")
	if err != nil {
		t.Error(err.Error())
		return
	}

	log.Println(conf)

	log.Printf("%+v\n", conf.Common)
	for name, node := range conf.Node {
		for _, service := range node.Services {
			log.Printf("%s: %+v\n", name, service)
		}
		log.Printf("%s: %+v\n", name, node)
	}
}

func TestNewDiagConfig(t *testing.T) {
	confPath := "../config/diag_debug.yaml"

	conf, err := config.NewDiagConfig(confPath, "yaml")
	if err != nil {
		t.Error(err.Error())
		return
	}

	log.Println(conf)
}

func TestNewPackageConfig(t *testing.T) {
	confPath := "../config/package_debug.yaml"

	conf, err := config.NewPackageConfig(confPath, "yaml")
	if err != nil {
		t.Error(err.Error())
		return
	}

	log.Println(conf)
}
