package service

import (
	"sync"

	"github.com/nebula/nebula-diagnose/pkg/config"
	"github.com/nebula/nebula-diagnose/pkg/remote"
)

var collectors = make(map[string]*NebulaCollector)
var mux sync.RWMutex

func GetServiceCollector(ncid string, conf *config.NodeConfig, serviceConfig *config.ServiceConfig) (*NebulaCollector, error) {
	mux.Lock()
	if _, ok := collectors[ncid]; !ok {
		c, err := newServiceCollector(ncid, conf, serviceConfig)
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

func newServiceCollector(ncid string, conf *config.NodeConfig, serviceConfig *config.ServiceConfig) (*NebulaCollector, error) {
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
