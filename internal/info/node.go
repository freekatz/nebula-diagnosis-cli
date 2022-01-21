package info

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/nebula/nebula-diagnose/pkg/config"
	"github.com/nebula/nebula-diagnose/pkg/errorx"
	"github.com/nebula/nebula-diagnose/pkg/info/physical"
	"github.com/nebula/nebula-diagnose/pkg/info/service"
	"github.com/nebula/nebula-diagnose/pkg/logger"
	"github.com/nebula/nebula-diagnose/pkg/utils"
)

// NewAllInfo Save the last collected information
var NewAllInfo AllInfo

type AllInfo struct {
	Time        string                                `json:"time,omitempty"`
	PhyInfo     *physical.PhyInfo                     `json:"phy_info,omitempty"`
	StatusInfo  map[string]*service.NebulaStatusInfo  `json:"status_info,omitempty"`
	MetricsInfo map[string]*service.NebulaMetricsInfo `json:"metrics_info,omitempty"`
	FlagsInfo   map[string]*service.NebulaFlagsInfo   `json:"flags_info,omitempty"`
}

func fetchAndSaveInfo(nodeConfig *config.NodeConfig, option config.InfoOption, defaultLogger logger.Logger) {
	allInfo := fetchInfo(nodeConfig, option, defaultLogger)
	NewAllInfo = *allInfo
	marshal, err := json.Marshal(allInfo)
	if err != nil {
		defaultLogger.Errorf("save json data failed: %s", err.Error())
	}

	dir := filepath.Join(nodeConfig.OutputDirPath, nodeConfig.SSH.Address)
	absDir, _ := filepath.Abs(dir)
	if !utils.IsDirExisted(absDir) {
		os.MkdirAll(absDir, os.ModePerm)
	}

	timeUnix := time.Now().Unix()
	filename := fmt.Sprintf("%d%s", timeUnix, ".data")
	filePath := filepath.Join(absDir, filename)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		defaultLogger.Fatal(err)
	}
	_, err = file.Write(marshal)
	_, err = file.Write([]byte("\n"))
	if err != nil {
		defaultLogger.Errorf("save json data failed: %s", err.Error())
	}
}

func fetchInfo(nodeInfo *config.NodeConfig, option config.InfoOption, defaultLogger logger.Logger) *AllInfo {
	allInfo := new(AllInfo)

	phyInfo, err := fetchPhyInfo(option, nodeInfo.SSH)
	if err != nil {
		defaultLogger.Errorf("fetch phy info failed: %s", err.Error())
	} else {
		if phyInfo != nil {
			allInfo.PhyInfo = phyInfo
			defaultLogger.Infof("%s physical info: %+v", nodeInfo.SSH.Address, *phyInfo)
		}
	}

	statusInfos, err := fetchStatusInfo(option, nodeInfo, defaultLogger)
	if err != nil {
		defaultLogger.Errorf("fetch services status failed: %s", err.Error())
	} else {
		if statusInfos != nil {
			for name, statusInfo := range statusInfos {
				defaultLogger.Infof("%s status info: %+v", name, *statusInfo)
			}
			allInfo.StatusInfo = statusInfos
		}
	}

	metricsInfos, err := fetchMetricsInfo(option, nodeInfo, defaultLogger)
	if err != nil {
		defaultLogger.Errorf("fetch services metrics failed: %s", err.Error())
	} else {
		if metricsInfos != nil {
			for name, metricsInfo := range metricsInfos {
				defaultLogger.Infof("%s metrics info: %+v", name, *metricsInfo)
			}
			allInfo.MetricsInfo = metricsInfos
		}
	}

	flagsInfos, err := fetchFlagsInfo(option, nodeInfo, defaultLogger)
	if err != nil {
		defaultLogger.Errorf("fetch services flags failed: %s", err.Error())
	} else {
		if flagsInfos != nil {
			for name, flagsInfo := range flagsInfos {
				defaultLogger.Infof("%s flags info: %+v", name, *flagsInfo)
			}
			allInfo.FlagsInfo = flagsInfos
		}
	}

	err = packageLogs(nodeInfo, option, defaultLogger)
	defaultLogger.Info("packaging service logs...")
	if err != nil {
		defaultLogger.Errorf("service package: failed, %s", err.Error())
	} else {
		defaultLogger.Info(nodeInfo.SSH.Address, " service package: success")
	}

	return allInfo
}

func fetchPhyInfo(option config.InfoOption, sshConfig *config.SSHConfig) (*physical.PhyInfo, error) {
	if option == config.AllInfo || option == config.Physical {
		return physical.GetPhyInfo(sshConfig)
	}
	return nil, nil
}

func fetchStatusInfo(option config.InfoOption, nodeConfig *config.NodeConfig, defaultLogger logger.Logger) (map[string]*service.NebulaStatusInfo, error) {
	if option == config.AllInfo || option == config.Stats {
		serviceConfigs := nodeConfig.Services
		serviceStatusInfos := make(map[string]*service.NebulaStatusInfo, len(serviceConfigs))
		for name, serviceConfig := range serviceConfigs {
			statusInfo, err := service.GetStatusInfo(nodeConfig, serviceConfig)
			if err != nil {
				defaultLogger.Errorf("service %s fetch status info failed: %s", name, err.Error())
				continue
			}
			serviceStatusInfos[name] = statusInfo
		}
		if len(serviceStatusInfos) != len(serviceConfigs) {
			return serviceStatusInfos, errorx.ErrStatusInfoIncomplete
		}
		return serviceStatusInfos, nil
	}
	return nil, nil
}

func fetchFlagsInfo(option config.InfoOption, nodeConfig *config.NodeConfig, defaultLogger logger.Logger) (map[string]*service.NebulaFlagsInfo, error) {
	if option == config.AllInfo || option == config.Flags {
		serviceConfigs := nodeConfig.Services
		serviceFlagsInfos := make(map[string]*service.NebulaFlagsInfo, len(serviceConfigs))
		for name, serviceConfig := range serviceConfigs {
			flagsInfo, err := service.GetFlagsInfo(nodeConfig, serviceConfig)
			if err != nil {
				defaultLogger.Errorf("service %s fetch flags info failed: %s", name, err.Error())
				continue
			}
			serviceFlagsInfos[name] = flagsInfo
		}
		if len(serviceFlagsInfos) != len(serviceConfigs) {
			return serviceFlagsInfos, errorx.ErrFlagsInfoIncomplete
		}
		return serviceFlagsInfos, nil
	}
	return nil, nil
}

func fetchMetricsInfo(option config.InfoOption, nodeConfig *config.NodeConfig, defaultLogger logger.Logger) (map[string]*service.NebulaMetricsInfo, error) {
	if option == config.AllInfo || option == config.Metrics {
		serviceConfigs := nodeConfig.Services
		serviceMetricsInfos := make(map[string]*service.NebulaMetricsInfo, len(serviceConfigs))
		for name, serviceConfig := range serviceConfigs {
			flagsInfo, err := service.GetMetricsInfo(nodeConfig, serviceConfig)
			if err != nil {
				defaultLogger.Errorf("service %s fetch metrics info failed: %s", name, err.Error())
				continue
			}
			serviceMetricsInfos[name] = flagsInfo
		}
		if len(serviceMetricsInfos) != len(serviceConfigs) {
			return serviceMetricsInfos, errorx.ErrMetricsInfoIncomplete
		}
		return serviceMetricsInfos, nil
	}
	return nil, nil
}

func packageLogs(nodeConf *config.NodeConfig, option config.InfoOption, defaultLogger logger.Logger) error {
	var err error
	if option == config.AllInfo || option == config.Logs {
		serviceConfigs := nodeConf.Services
		for _, serviceConfig := range serviceConfigs {
			err = service.GetLogs(nodeConf, serviceConfig)
			if err != nil {
				defaultLogger.Errorf("service %s package logs failed: %s, stop package logs!", err.Error())
			}
		}
	}
	if err != nil {
		return errorx.ErrPackageLogsIncomplete
	}
	return nil
}
