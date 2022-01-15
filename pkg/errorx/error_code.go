package errorx

import (
	"errors"
	"fmt"
)

type (
	ErrorCode int64

	errorX struct {
		errorCode ErrorCode
		errorMsg  string
	}

	ErrorX interface {
		error
		GetErrorCode() ErrorCode
		GetErrorMsg() string
	}
)

func AsErrorX(err error) (ErrorX, bool) {
	if e := new(errorX); errors.As(err, &e) {
		return e, true
	}
	return nil, false
}

func IsErrorX(err error, c ...ErrorCode) bool {
	errX, ok := AsErrorX(err)
	if !ok {
		return false
	}
	switch len(c) {
	case 0:
		return true
	case 1:
		return errX.GetErrorCode() == c[0]
	default:
		return false
	}
}

func NewErrorX(errorCode ErrorCode, errorMsg string) error {
	return &errorX{
		errorCode: errorCode,
		errorMsg:  errorMsg,
	}
}

func (e *errorX) Error() string {
	return fmt.Sprintf("%d:%s", int64(e.errorCode), e.errorMsg)
}

func (e *errorX) GetErrorCode() ErrorCode {
	return e.errorCode
}

func (e *errorX) GetErrorMsg() string {
	return e.errorMsg
}

/*
	ErrorCode I: Info, W: Warn, E: Error
*/
const (
	ErrorCode_I_SUCCEEDED ErrorCode = 0
	ErrorCode_W_CLI       ErrorCode = -1
	ErrorCode_E_RUNTIME   ErrorCode = -2
	ErrorCode_E_INTERNAL  ErrorCode = -3
	ErrorCode_E_FILE      ErrorCode = -4
)
