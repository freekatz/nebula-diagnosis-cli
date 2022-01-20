package pack

import (
	"path/filepath"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/logger"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/packer"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/remote"
)

var _logger = logger.GetLogger("pack", "", false)

func Run(conf *config.PackConfig) {
	tgzPacker := packer.NewTgzPacker()
	err := tgzPacker.Pack(conf.TarFilepath, filepath.Join(conf.OutputDirPath, conf.TarFilename))
	if err != nil {
		_logger.Error(err)
		return
	}

	if conf.SSH != nil && conf.SSH.Validate() {
		sftpClient, err := remote.GetSFTPClient(conf.SSH.Username, conf.SSH)
		if err != nil {
			_logger.Error("get sftp client failed")
			return
		}

		ok := sftpClient.UploadFile(conf.RemoteDirPath, filepath.Join(conf.OutputDirPath, conf.TarFilename))
		if !ok {
			_logger.Error("upload failed")
			return
		}
	}
}
