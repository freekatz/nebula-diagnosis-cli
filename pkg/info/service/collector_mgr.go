package service

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/remote"
	"sync"
)

var collectors = make(map[string]*NebulaCollector)
var mux sync.RWMutex

func GetServiceExporter(ncid string, conf *config.NodeConfig, serviceConfig *config.ServiceConfig) (*NebulaCollector, error) {
	mux.Lock()
	if _, ok := collectors[ncid]; !ok {
		c, err := newServiceExporter(ncid, conf, serviceConfig)
		if err != nil {
			return nil, err
		}

		collectors[ncid] = c
	}
	mux.Unlock()

	mux.RLock()
	defer mux.RUnlock()

	return collectors[ncid], nil
}

func newServiceExporter(ncid string, conf *config.NodeConfig, serviceConfig *config.ServiceConfig) (*NebulaCollector, error) {
	serviceType := serviceConfig.Type
	c := new(NebulaCollector)
	c.NebulaType = serviceType
	c.NodeConfig = conf
	c.ServiceConfig = serviceConfig
	c.Id = ncid
	sshClient, err := remote.GetSFTPClient(c.Id, c.NodeConfig.SSH)
	if err != nil {
		return nil, err
	}
	c.SshClient = sshClient
	return c, nil
}
