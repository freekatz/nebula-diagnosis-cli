package config

type (
	SSHConfig struct {
		// for node config, the ssh address equals to service address
		Address  string `mapstructure:"address"`
		Port     int    `mapstructure:"port"`
		Timeout  string `mapstructure:"timeout,omitempty"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`

		// TODO support the private key login
	}
)

const (
	defaultSSHTimeout    = "3s"
	defaultOutputDirPath = "./output"
)

func (c *SSHConfig) Complete() {
	if c.Timeout == "" {
		c.Timeout = defaultSSHTimeout
	}
}

func (c *SSHConfig) Validate() bool {
	return c.Address != "" && c.Port > 0 && ValidateTimeout(c.Timeout) && c.Username != "" && c.Password != "" // TODO add more exactly verify: port
}
