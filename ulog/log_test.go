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
	l.Debug("test debug")
	l.Info("test info")
	l.Warn("test warn")
	l.Error(uerr.NewError(errors.New("test error")))
	l.Fatal(uerr.NewError(errors.New("test fatal")))
}

func TestLogSave(t *testing.T) {
	l := ulog.NewLogWithoutPost(5, true, "./")
	l.Debug("test debug")
	l.Info("test info")
	l.Warn("test warn")
	l.Error(uerr.NewError(errors.New("test error")))
	l.Fatal(uerr.NewError(errors.New("test fatal")))
}

// Already tested in local
func TestSendLog(t *testing.T) {
	l := ulog.NewLogWithPost(5, false, "./", "http://localhost:24242", "1")
	l.Debug("test debug")
	time.Sleep(10 * time.Second) // Wait for the log to be sent
}

func TestImplementLogInterface(t *testing.T) {
	l0 := ulog.NewLogWithoutPost(5, false, "./")
	var logInterface ulog.Log = l0
	logInterface.Debug("test debug")
	logInterface.Info("test info")
	logInterface.Warn("test warn")
	logInterface.Error(uerr.NewError(errors.New("test error")))
	logInterface.Fatal(uerr.NewError(errors.New("test fatal")))

	l1 := ulog.NewLogWithPost(5, false, "./", "http://localhost:24242", "1")
	logInterface = l1
	logInterface.Debug("test debug")
	logInterface.Info("test info")
	logInterface.Warn("test warn")
	logInterface.Error(uerr.NewError(errors.New("test error")))
	logInterface.Fatal(uerr.NewError(errors.New("test fatal")))

	l2 := ulog.NewLeaderLog(5, false, "./")
	logInterface = l2
	logInterface.Debug("test debug")
	logInterface.Info("test info")
	logInterface.Warn("test warn")
	logInterface.Error(uerr.NewError(errors.New("test error")))
	logInterface.Fatal(uerr.NewError(errors.New("test fatal")))
}
