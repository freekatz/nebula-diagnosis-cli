package config

import (
	"path/filepath"
	"strings"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"

	"github.com/spf13/viper"
)

type PackConfig struct {
	OutputDirPath string     `mapstructure:"outputDirPath,omitempty"` // output location
	TarFilepath   string     `mapstructure:"tarFilepath"`             // input tar filepath
	TarFilename   string     `mapstructure:"tarFilename"`             // output tar filename, will output into OutputDirPath
	RemoteDirPath string     `mapstructure:remoteDirPath`             // remote upload path
	SSH           *SSHConfig `mapstructure:"ssh"`                     // ssh config for upload
}

func (c *PackConfig) Complete() {
	if c.OutputDirPath == "" {
		c.OutputDirPath = defaultOutputDirPath
	}
	if c.TarFilename == "" {
		c.TarFilename = strings.Join([]string{filepath.Base(c.TarFilepath), ".tar.gz"}, "")
	}
	if !strings.HasSuffix(c.TarFilename, ".tar") && !strings.HasSuffix(c.TarFilename, ".gz") {
		c.TarFilename = strings.Join([]string{c.TarFilename, ".tar.gz"}, "")
	}
}

func (c *PackConfig) Validate() bool {
	if c.SSH != nil && c.RemoteDirPath == "" {
		return false
	}
	return c.TarFilepath != ""
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
