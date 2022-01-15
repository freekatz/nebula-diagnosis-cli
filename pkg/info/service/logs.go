package service

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"strconv"
)

func GetLogs(nodeConfig *config.NodeConfig, serviceConfig *config.ServiceConfig) error {
	ncid := nodeConfig.SSH.Address + ":" + strconv.Itoa(serviceConfig.Port)
	exporter, err := GetServiceCollector(ncid, nodeConfig, serviceConfig)
	if err != nil {
		return err
	}
	err = exporter.PackageLogs()
	if err != nil {
		return err
	}
	return err
}
