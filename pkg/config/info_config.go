package config

import (
	"time"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"
	"github.com/spf13/viper"
)

type (
	InfoConfig struct {
		Common *CommonConfig          `mapstructure:"common"` // common config for each node
		Node   map[string]*NodeConfig `mapstructure:"node"`   // node specific config
	}

	CommonConfig struct {
		OutputDirPath string       `mapstructure:"outputDirPath"`      // output location
		Duration      string       `mapstructure:"duration,omitempty"` // duration of info fetching, default is 0
		Period        string       `mapstructure:"period,omitempty"`   // period of info fetching, default is 0
		Options       []InfoOption `mapstructure:"option,omitempty"`   // info to fetch, default is all
	}

	NodeConfig struct {
		Host          *HostConfig               `mapstructure:"host"`               // node host
		SSH           *SSHConfig                `mapstructure:"ssh"`                // node ssh
		Services      map[string]*ServiceConfig `mapstructure:"services"`           // node service
		OutputDirPath string                    `mapstructure:"outputDirPath"`      // output location
		Duration      string                    `mapstructure:"duration,omitempty"` // duration of info fetching, default is 0
		Period        string                    `mapstructure:"period,omitempty"`   // period of info fetching, default is 0
		Options       []InfoOption              `mapstructure:"option,omitempty"`   // info to fetch, default is all
	}

	HostConfig struct {
		Address string `mapstructure:"address"`
		Port    int    `mapstructure:"port"`
	}

	ServiceConfig struct {
		Type       ComponentType `mapstructure:"type"`
		DeployDir  string        `mapstructure:"deploy_dir"`
		RuntimeDir string        `mapstructure:"runtime_dir"`
		Port       int           `mapstructure:"port"`
		HTTPPort   int           `mapstructure:"http_port"`
		HTTP2Port  int           `mapstructure:"http_2_port"`
	}
)

type InfoOption string

const (
	Metrics  InfoOption = "metrics"
	Physical InfoOption = "physical"
	Stats    InfoOption = "stats"
	AllInfo  InfoOption = "all"
	NoInfo   InfoOption = "no"
)

type ComponentType string

const (
	ServiceGraph   ComponentType = "serviceGraph"
	ServiceMeta    ComponentType = "serviceMeta"
	ServiceStorage ComponentType = "serviceStorage"
)

var (
	defaultDuration    = "0"
	defaultPeriod      = "5s"
	defaultInfoOptions = []InfoOption{AllInfo}
)

func (c *InfoConfig) Complete() {
	if c.Common == nil {
		c.Common = new(CommonConfig)
	}
	c.Common.Complete()
	if c.Node == nil {
		c.Node = map[string]*NodeConfig{}
	}
	for _, node := range c.Node {
		node.Complete(c.Common)
	}
}

func (c *InfoConfig) Validate() bool {
	for _, info := range c.Node {
		if !info.Validate() {
			return false
		}
	}
	return true
}

func (c *CommonConfig) Complete() {
	if c.OutputDirPath == "" {
		c.OutputDirPath = defaultOutputDirPath
	}
	if c.Duration == "" {
		c.Duration = defaultDuration
	}
	if c.Period == "" {
		c.Period = defaultPeriod
	}
	if len(c.Options) == 0 {
		c.Options = defaultInfoOptions
	}
}

func (c *NodeConfig) Complete(common *CommonConfig) {
	if c.SSH != nil {
		c.SSH.Complete()
	}
	if c.OutputDirPath == "" {
		c.OutputDirPath = common.OutputDirPath
	}
	if c.OutputDirPath == "" {
		c.OutputDirPath = defaultOutputDirPath
	}
	if c.Duration == "" {
		c.Duration = common.Duration
	}
	if c.Duration == "" {
		c.Duration = defaultDuration
	}
	if c.Period == "" {
		c.Period = common.Period
	}
	if c.Period == "" {
		c.Period = defaultPeriod
	}
	if len(c.Options) == 0 {
		c.Options = common.Options
	}
	if len(c.Options) == 0 {
		c.Options = defaultInfoOptions
	}
}

func (c *NodeConfig) Validate() bool {
	if c.Host == nil || c.SSH == nil || c.Services == nil || !c.Host.Validate() || !c.SSH.Validate() {
		return false
	}
	for _, s := range c.Services {
		if !s.Validate() {
			return false
		}
	}
	if !ValidateDuration(c.Duration) || !ValidatePeriod(c.Period) {
		return false
	}
	return true
}

func (c *HostConfig) Validate() bool {
	return c.Address != "" && c.Port > 0 // TODO add more exactly verify: address, port
}

func (c *ServiceConfig) Validate() bool {
	return c.HTTPPort > 0 // TODO add more exactly verify: DeployDir, RuntimeDir
}

func ValidateDuration(duration string) bool {
	_, err := time.ParseDuration(duration)
	return duration == "-1" || err == nil
}

func ValidatePeriod(period string) bool {
	d, err := time.ParseDuration(period)
	return err == nil && d > 0
}

func ValidateTimeout(timeout string) bool {
	d, err := time.ParseDuration(timeout)
	return err == nil && d >= 0
}

func NewInfoConfig(confPath string, configType string) (*InfoConfig, error) {
	var viperConfig = viper.New()
	viperConfig.SetConfigName(confPath)
	viperConfig.SetConfigFile(confPath)
	viperConfig.SetConfigType(configType)
	if err := viperConfig.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := new(InfoConfig)
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
