package tests

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/info/service"
	"log"
	"testing"
)

func TestCollector(t *testing.T)  {
	serviceConf := &config.ServiceConfig{
		Type: config.ServiceGraph,
		RuntimeDir: "/home/katz.zhang",
		HTTPPort: 19669,
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
		OutputDirPath: "./logs",
	}
	ncid := "123"
	exporter, err := service.GetServiceExporter(ncid, nodeConf, serviceConf)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = exporter.CollectMetricsInfo()
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = exporter.CollectStatusInfo()
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = exporter.CollectFlagsInfo()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = exporter.PackageLogs()
	if err != nil {
		log.Fatal(err.Error())
	}
}