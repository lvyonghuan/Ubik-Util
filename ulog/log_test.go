package ulog_test

import (
	"errors"
	"testing"
	"time"

	"github.com/lvyonghuan/Ubik-Util/uerr"
	"github.com/lvyonghuan/Ubik-Util/ulog"
)

func TestLogPrint(t *testing.T) {
	l := ulog.NewLogWithoutPost(5, false, "./")
	l.InitLog()
	l.Debug("test debug")
	l.Info("test info")
	l.Warn("test warn")
	l.Error(uerr.NewError(errors.New("test error")))
	l.Fatal(uerr.NewError(errors.New("test fatal")))
}

func TestLogSave(t *testing.T) {
	l := ulog.NewLogWithoutPost(5, true, "./")
	l.InitLog()
	l.Debug("test debug")
	l.Info("test info")
	l.Warn("test warn")
	l.Error(uerr.NewError(errors.New("test error")))
	l.Fatal(uerr.NewError(errors.New("test fatal")))
}

func TestLogWithoutInitPrint(t *testing.T) {
	l := ulog.NewLogWithoutPost(5, false, "./")
	l.Debug("test debug")
	l.Info("test info")
	l.Warn("test warn")
	l.Error(uerr.NewError(errors.New("test error")))
	l.Fatal(uerr.NewError(errors.New("test fatal")))
}

func TestSendLog(t *testing.T) {
	l := ulog.NewLogWithPost(5, false, "./", "http://localhost:24242", "1")
	l.Debug("test debug")
	time.Sleep(10 * time.Second) // Wait for the log to be sent
}
