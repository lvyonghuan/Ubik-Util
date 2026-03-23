package ulog_test

import (
	"errors"
	"testing"

	"github.com/lvyonghuan/Ubik-Util/uerr"
	"github.com/lvyonghuan/Ubik-Util/ulog"
)

func TestLogPrint(t *testing.T) {
	l := ulog.NewULog(5, 5, false, "./")
	l.Debug("test debug")
	l.Info("test info")
	l.Warn("test warn")
	l.Error(uerr.NewError(errors.New("test error")))
	l.Fatal(uerr.NewError(errors.New("test fatal")))
}

func TestLogSave(t *testing.T) {
	l := ulog.NewULog(5, 5, true, "./")
	l.Debug("test debug")
	l.Info("test info")
	l.Warn("test warn")
	l.Error(uerr.NewError(errors.New("test error")))
	l.Fatal(uerr.NewError(errors.New("test fatal")))
}
