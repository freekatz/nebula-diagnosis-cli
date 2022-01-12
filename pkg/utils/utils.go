package utils

import "strings"

func GetConfigType(confPath string) string {
	if strings.HasSuffix(confPath, "yaml") {
		return "yaml"
	}

	return ""
}
