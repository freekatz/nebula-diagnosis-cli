package service

import (
	"strconv"

	"github.com/nebula/nebula-diagnose/pkg/config"
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
