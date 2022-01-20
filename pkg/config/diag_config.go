package config

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"

	"github.com/spf13/viper"
)

type DiagConfig struct {
	OutputDirPath string       `mapstructure:"outputDirPath,omitempty"` // output location
	LogToFile     bool         `mapstructure:"logToFile"`               // logging to file or cmd
	InputDirPath  string       `mapstructure:"inputDirPath"`            // input location
	Options       []DiagOption `mapstructure:"option,omitempty"`        // diag result to analyze, default is all
}

type DiagOption string

const (
	Partition DiagOption = "partition"
	AllDiag   DiagOption = "all"
)

var defaultDiagOptions = []DiagOption{AllDiag}

func (c *DiagConfig) Complete() {
	if c.OutputDirPath == "" {
		c.OutputDirPath = defaultOutputDirPath
	}
	if len(c.Options) == 0 {
		c.Options = defaultDiagOptions
	}
}

func (c *DiagConfig) Validate() bool {
	return c.InputDirPath != ""
}

func NewDiagConfig(confPath string, configType string) (*DiagConfig, error) {
	var viperConfig = viper.New()
	viperConfig.SetConfigName(confPath)
	viperConfig.SetConfigFile(confPath)
	viperConfig.SetConfigType(configType)
	if err := viperConfig.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := new(DiagConfig)
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
