package pack

import (
	"path/filepath"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/logger"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/packer"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/remote"
)

var _logger = logger.GetLogger("pack", "")

func Run(conf *config.PackConfig) {
	tgzPacker := packer.NewTgzPacker()
	err := tgzPacker.Pack(conf.TarFilepath, filepath.Join(conf.OutputDirPath, conf.TarFilename))
	if err != nil {
		_logger.Error(false, err)
	}

	if conf.SSH != nil && conf.SSH.Validate() {
		sftpClient, err := remote.GetSFTPClient(conf.SSH.Username, conf.SSH)
		if err != nil {
			_logger.Error(false, "get sftp client failed")
			return
		}
		ok := sftpClient.UploadFile()
		if !ok {
			_logger.Error(false, "upload failed")
		}
	}
}
