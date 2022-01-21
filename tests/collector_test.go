package tests

import (
	"log"
	"testing"

	"github.com/nebula/nebula-diagnose/pkg/config"
	"github.com/nebula/nebula-diagnose/pkg/info/service"
)

func TestCollector(t *testing.T) {
	serviceConf := &config.ServiceConfig{
		Type:       config.ServiceGraph,
		RuntimeDir: "/home/katz.zhang",
		HTTPPort:   19669,
	}
	sshConf := &config.SSHConfig{
		Address:  "192.168.8.49",
		Port:     22,
		Timeout:  "1s",
		Username: "katz.zhang",
		Password: "nebula",
	}
	nodeConf := &config.NodeConfig{
		SSH: sshConf,
		Services: map[string]*config.ServiceConfig{
			"graph1": serviceConf,
		},
		OutputDirPath: "./tmp/logs",
	}
	ncid := "123"
	collector, err := service.GetServiceCollector(ncid, nodeConf, serviceConf)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = collector.CollectMetricsInfo()
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = collector.CollectStatusInfo()
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = collector.CollectFlagsInfo()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = collector.PackageLogs()
	if err != nil {
		log.Fatal(err.Error())
	}
}
