package service

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"strconv"
)

func GetFlagsInfo(nodeConfig *config.NodeConfig, serviceConfig *config.ServiceConfig) (*NebulaFlagsInfo, error) {
	ncid := nodeConfig.SSH.Address + ":" + strconv.Itoa(serviceConfig.Port)
	exporter, err := GetServiceCollector(ncid, nodeConfig, serviceConfig)
	if err != nil {
		return nil, err
	}
	flagsInfo, err := exporter.CollectFlagsInfo()
	if err != nil {
		return nil, err
	}
	return &flagsInfo, nil
}
