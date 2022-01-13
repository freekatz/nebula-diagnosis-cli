package errorx

import "errors"

var (
	/*
		Err cli errors
	*/
	ErrPrintAndExit = errors.New("print and exit")
	ErrNoConfig     = errors.New("have no config")
	ErrNoInputDir   = errors.New("have no input dir")

	/*
		Err internal errors
	*/
	ErrConfigInvalid = errors.New("config invalid")
	ErrLogDirInvalid = errors.New("log dir path invalid")
	ErrSSHExecFailed = errors.New("ssh exec failed")
)
