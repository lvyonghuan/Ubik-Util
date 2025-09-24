package uerr

import (
	"runtime"
)

type UbikError struct {
	err       error  //error message
	stackInfo string //call stack
	errorInfo string
}

func NewError(err error) UbikError {
	return NewErrorWithInfo(err, "")
}

func NewErrorWithInfo(err error, info string) UbikError {
	if err == nil {
		return UbikError{}
	}

	stackBuf := make([]byte, 1024)
	stackSize := runtime.Stack(stackBuf, true)
	stack := string(stackBuf[:stackSize])

	return UbikError{
		err:       err,
		stackInfo: stack,
		errorInfo: info,
	}
}

func (err UbikError) Error() string {
	return err.err.Error() + "\n" + err.stackInfo
}

func (err UbikError) Stack() string {
	return err.stackInfo
}

func (err UbikError) MetaError() error {
	return err.err
}

func (err UbikError) Info() string {
	return err.errorInfo
}
