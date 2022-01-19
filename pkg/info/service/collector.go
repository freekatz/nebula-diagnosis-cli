package service

import (
	"path/filepath"
	"strings"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/remote"
)

type NebulaCollector struct {
	Id string

	// Save and return the latest collected data,
	NebulaStatusInfo  *NebulaStatusInfo
	NebulaMetricsInfo *NebulaMetricsInfo
	NebulaFlagsInfo   *NebulaFlagsInfo

	NebulaType    config.ComponentType
	NodeConfig    *config.NodeConfig
	ServiceConfig *config.ServiceConfig

	SshClient *remote.SFTPClient
}

type (
	NebulaStatusInfo struct {
		GitInfoSha string `json:"git_info_sha"`
		Status     string `json:"status"`
	}
	NebulaMetricsInfo struct {
		Metrics map[string]QueryValue `json:"metrics"`
	}
	NebulaFlagsInfo struct {
		Flags map[string]QueryValue `json:"configs"`
	}
)

func (c *NebulaCollector) CollectStatusInfo() (NebulaStatusInfo, error) {
	ip := c.NodeConfig.SSH.Address
	port := c.ServiceConfig.HTTPPort
	status, err := remote.GetNebulaComponentStatus(ip, port)
	if err != nil {
		return NebulaStatusInfo{}, err
	}
	c.NebulaStatusInfo = &NebulaStatusInfo{
		GitInfoSha: status[0],
		Status:     status[1],
	}
	return *c.NebulaStatusInfo, nil
}

func (c *NebulaCollector) CollectMetricsInfo() (NebulaMetricsInfo, error) {
	ip := c.NodeConfig.SSH.Address
	port := c.ServiceConfig.HTTPPort
	metrics, err := remote.GetNebulaMetrics(ip, port)
	if err != nil {
		return NebulaMetricsInfo{}, err
	}
	c.NebulaMetricsInfo = &NebulaMetricsInfo{
		Metrics: ConvertToMap(metrics),
	}
	return *c.NebulaMetricsInfo, nil
}

func (c *NebulaCollector) CollectFlagsInfo() (NebulaFlagsInfo, error) {
	ip := c.NodeConfig.SSH.Address
	port := c.ServiceConfig.HTTPPort
	flags, err := remote.GetNebulaFlags(ip, port)
	if err != nil {
		return NebulaFlagsInfo{}, err
	}
	c.NebulaFlagsInfo = &NebulaFlagsInfo{
		Flags: ConvertToMap(flags),
	}
	return *c.NebulaFlagsInfo, nil
}

func (c *NebulaCollector) PackageLogs() error {
	remoteLogDir := ""
	if c.NebulaFlagsInfo == nil {
		return errorx.ErrPackageLogsFailed
	}
	if queryValue, ok := c.NebulaFlagsInfo.Flags["log_dir"]; ok {
		remoteLogDir = queryValue.Value
	}
	if remoteLogDir == "" {
		return errorx.ErrRemoteLogDirInvalid
	}
	if !strings.HasPrefix(remoteLogDir, "/") {
		remoteLogDir = c.ServiceConfig.RuntimeDir + "/" + remoteLogDir
	}
	newDir := c.Id
	localDir := filepath.Join(c.NodeConfig.OutputDirPath, newDir)
	err := c.SshClient.DownloadDir(remoteLogDir, localDir)
	if err != nil {
		return err
	}
	return nil
}
