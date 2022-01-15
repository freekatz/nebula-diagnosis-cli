package service

import (
	"strconv"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
)

func GetMetricsInfo(nodeConfig *config.NodeConfig, serviceConfig *config.ServiceConfig) (*NebulaMetricsInfo, error) {
	ncid := nodeConfig.SSH.Address + ":" + strconv.Itoa(serviceConfig.Port)
	exporter, err := GetServiceCollector(ncid, nodeConfig, serviceConfig)
	if err != nil {
		return nil, err
	}
	metricsInfo, err := exporter.CollectMetricsInfo()
	if err != nil {
		return nil, err
	}
	return &metricsInfo, nil
}
