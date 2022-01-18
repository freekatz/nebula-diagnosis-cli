package errorx

import "errors"

var (
	/*
		ErrorX cli errors
	*/
	ErrPrintAndExit    = errors.New("print and exit")
	ErrNoInputConfig   = errors.New("have no input config")
	ErrNoInputFilepath = errors.New("have no input filepath")

	/*
		ErrorX runtime errors
	*/
	ErrPackageLogsFailed = errors.New("package logs failed")

	/*
		ErrorX internal errors
	*/
	ErrConfigInvalid         = errors.New("config invalid")
	ErrRemoteLogDirInvalid   = errors.New("remote log dir path invalid")
	ErrSSHExecFailed         = errors.New("ssh exec failed")
	ErrStatusInfoIncomplete  = errors.New("status info collect incomplete")
	ErrFlagsInfoIncomplete   = errors.New("flags info collect incomplete")
	ErrMetricsInfoIncomplete = errors.New("metrics info collect incomplete")
	ErrPackageLogsIncomplete = errors.New("package logs incomplete")

	/*
		ErrorX file errors
	*/
	ErrFileNotExisted   = errors.New("file not existed")
	ErrFileHasExisted   = errors.New("file has existed")
	ErrFileTypeNotMatch = errors.New("file type not match")
)
