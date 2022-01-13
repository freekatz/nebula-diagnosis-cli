package errorx

import "errors"

var (
	ErrConfigInvalid = errors.New("config invalid")
	ErrLogDirInvalid = errors.New("log Dir invalid")
)
