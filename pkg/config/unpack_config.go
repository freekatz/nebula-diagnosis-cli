package config

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"
)

type UnPackConfig struct {
	OutputDirPath string `mapstructure:"outputDirPath,omitempty"` // output location
	InputDirPath  string `mapstructure:"inputDirPath"`            // input location
}

// Complete TODO modify the output dir path
func (c *UnPackConfig) Complete() {
	if c.OutputDirPath == "" {
		c.OutputDirPath = defaultOutputDirPath
	}
}

func (c *UnPackConfig) Validate() bool {
	return c.InputDirPath != ""
}

func NewUnPackConfig(inputDirPath, outputDirPath string) (*UnPackConfig, error) {
	conf := new(UnPackConfig)
	conf.InputDirPath = inputDirPath
	conf.OutputDirPath = outputDirPath
	conf.Complete()
	if ok := conf.Validate(); !ok {
		return nil, errorx.ErrConfigInvalid
	}
	return conf, nil
}
