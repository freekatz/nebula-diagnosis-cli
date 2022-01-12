package config

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"
	"github.com/spf13/viper"
)

type PackageConfig struct {
	OutputDirPath string    `mapstructure:"outputDirPath,omitempty"` // output location
	InputDirPath  string    `mapstructure:"inputDirPath"`            // input location
	SSH           SSHConfig `mapstructure:"ssh"`                     // ssh config for upload
}

func (c *PackageConfig) Complete() {
	if c.OutputDirPath == "" {
		c.OutputDirPath = defaultOutputDirPath
	}
}

func (c *PackageConfig) Validate() bool {
	return c.InputDirPath != ""
}

func NewPackageConfig(confPath string, configType string) (*PackageConfig, error) {
	var viperConfig = viper.New()
	viperConfig.SetConfigName(confPath)
	viperConfig.SetConfigFile(confPath)
	viperConfig.SetConfigType(configType)
	if err := viperConfig.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := new(PackageConfig)
	err := viperConfig.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	conf.Complete()
	if ok := conf.Validate(); !ok {
		return nil, errorx.ErrConfigInvalid
	}
	return conf, nil
}
