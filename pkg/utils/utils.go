package utils

import (
	"os"
	"strings"
)

func GetConfigType(confPath string) string {
	if strings.HasSuffix(confPath, "yaml") {
		return "yaml"
	}

	return ""
}

func IsFileExisted(filePath string) bool {
	info, err := os.Stat(filePath)
	return (err == nil || os.IsExist(err)) && !info.IsDir()
}

func IsDirExisted(dirPath string) bool {
	info, err := os.Stat(dirPath)
	return (err == nil || os.IsExist(err)) && info.IsDir()
}
