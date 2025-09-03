package uerr

import (
	"runtime"
)

type UbikError struct {
	err       error  //error message
	stackInfo string //call stack
}

func NewError(err error) UbikError {
	if err == nil {
		return UbikError{}
	}

	stackBuf := make([]byte, 1024)
	stackSize := runtime.Stack(stackBuf, true)
	stack := string(stackBuf[:stackSize])

	return UbikError{
		err:       err,
		stackInfo: stack,
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
