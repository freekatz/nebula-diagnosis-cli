package config

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"
	"github.com/spf13/viper"
)

type PackConfig struct {
	OutputDirPath string    `mapstructure:"outputDirPath,omitempty"` // output location
	InputDirPath  string    `mapstructure:"inputDirPath"`            // input location
	SSH           SSHConfig `mapstructure:"ssh"`                     // ssh config for upload
}

// Complete TODO modify the output dir path
func (c *PackConfig) Complete() {
	if c.OutputDirPath == "" {
		c.OutputDirPath = defaultOutputDirPath
	}
}

func (c *PackConfig) Validate() bool {
	return c.InputDirPath != ""
}

func NewPackConfig(confPath string, configType string) (*PackConfig, error) {
	var viperConfig = viper.New()
	viperConfig.SetConfigName(confPath)
	viperConfig.SetConfigFile(confPath)
	viperConfig.SetConfigType(configType)
	if err := viperConfig.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := new(PackConfig)
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
