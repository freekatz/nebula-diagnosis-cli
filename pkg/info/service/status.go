package service

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"strconv"
)

func GetStatusInfo(nodeConfig *config.NodeConfig, serviceConfig *config.ServiceConfig) (*NebulaStatusInfo, error) {
	ncid := nodeConfig.SSH.Address + ":" + strconv.Itoa(serviceConfig.Port)
	exporter, err := GetServiceCollector(ncid, nodeConfig, serviceConfig)
	if err != nil {
		return nil, err
	}
	statusInfo, err := exporter.CollectStatusInfo()
	if err != nil {
		return nil, err
	}
	return &statusInfo, nil
}
