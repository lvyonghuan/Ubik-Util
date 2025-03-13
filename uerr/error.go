package uerr

import (
	"runtime"
)

type UbikError struct {
	errorMessage string //错误信息
	stackInfo    string //调用堆栈
}

func NewError(err error) UbikError {
	if err == nil {
		return UbikError{}
	}

	stackBuf := make([]byte, 1024)
	stackSize := runtime.Stack(stackBuf, true)
	stack := string(stackBuf[:stackSize])

	return UbikError{
		errorMessage: err.Error(),
		stackInfo:    stack,
	}
}

func (err UbikError) Error() string {
	return err.errorMessage + "\n" + err.stackInfo
}

func (err UbikError) Stack() string {
	return err.stackInfo
}

func (err UbikError) UbikErrorMessage() string {
	return err.errorMessage
}
