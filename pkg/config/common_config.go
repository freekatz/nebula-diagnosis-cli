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

	OutputConfig struct {
		SSH     *SSHConfig `mapstructure:"ssh"`               // if the ssh config != nil, will output the data to remote
		DirPath string     `mapstructure:"dirPath,omitempty"` // output dir included logs, info, diag, etc., default is ./output, and will auto create if not existed
	}

	InputConfig struct {
		SSH     *SSHConfig `mapstructure:"ssh"`               // if the ssh config != nil, will read the input data from remote
		DirPath string     `mapstructure:"dirPath,omitempty"` // input dir included logs, info, etc.
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

func (c *OutputConfig) Complete() {
	c.SSH.Complete()
	if c.DirPath == "" {
		c.DirPath = defaultOutputDirPath
	}
}

func (c *OutputConfig) Validate() bool {
	return c.SSH != nil && c.SSH.Validate() && c.DirPath != "" // TODO add more exactly verify: dirPath
}

func (c *InputConfig) Complete() {
	c.SSH.Complete()
	if c.DirPath == "" {
		c.DirPath = defaultOutputDirPath
	}
}

func (c *InputConfig) Validate() bool {
	return c.SSH != nil && c.SSH.Validate() && c.DirPath != "" // TODO add more exactly verify: dirPath
}
