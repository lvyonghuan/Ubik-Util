package ulog

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/lvyonghuan/Ubik-Util/uconst"
)

type Log interface {
	InitLog()
	Debug(v string)
	Info(v string)
	Warn(v string)
	Error(v error)
	Fatal(v error)
	System(v string)
	SaveLogToFile(v string)
}

type ULog struct {
	Level       int        `json:"level"`       //log level
	WriteLevel  int        `json:"write_level"` //the log level to be written to the file
	IsSave      bool       `json:"is_save"`     //whether to save logs
	LogSavePath string     `json:"save_path"`   //the path where the logs are saved
	fileMutex   sync.Mutex // mutex for file operations
}

// NewULog creates a new ULog instance, initializes it, and returns it.
func NewULog(level int, writeLevel int, isSave bool, logSavePath string) *ULog {
	ULog := &ULog{
		Level:       level,
		WriteLevel:  writeLevel,
		IsSave:      isSave,
		LogSavePath: logSavePath,
	}

	ULog.InitLog()
	return ULog
}

// log levels
const (
	System = uconst.System
	Fatal  = uconst.Fatal
	Error  = uconst.Error
	Warn   = uconst.Warn
	Info   = uconst.Info //Default log level
	Debug  = uconst.Debug
)

// log colors
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	orange = "\033[33m"
	yellow = "\033[93m"
	green  = "\033[32m"
)

// InitLog init log
func (l *ULog) InitLog() {
	if l.IsSave {
		currentTime := time.Now().Format("2006-01-02_15-04-05")
		l.LogSavePath = l.LogSavePath + currentTime + ".log"
	}
}

// Debug print debug level logs
func (l *ULog) Debug(v string) {
	logString := "Debug: " + v
	if l.Level >= Debug {
		log.Println(green + logString + reset)
	}

	if l.WriteLevel >= Debug {
		l.SaveLogToFile(logString)
	}
}

// Info print info level logs
func (l *ULog) Info(v string) {
	logString := "Info: " + v
	if l.Level >= Info {
		log.Println(logString)
	}

	if l.WriteLevel >= Info {
		l.SaveLogToFile(logString)
	}
}

// Warn print the warn level logs
func (l *ULog) Warn(v string) {
	logString := "Warn: " + v
	if l.Level >= Warn {
		log.Println(yellow + logString + reset)
	}

	if l.WriteLevel >= Warn {
		l.SaveLogToFile(logString)
	}
}

// Error print the error level log
func (l *ULog) Error(v error) {
	logString := "Error: " + v.Error()

	if l.Level >= Error {
		log.Println(orange + logString + reset)
	}

	if l.WriteLevel >= Error {
		l.SaveLogToFile(logString)
	}
}

// Fatal print the fatal level logs
func (l *ULog) Fatal(v error) {
	logString := "Fatal: " + v.Error()
	if l.Level >= Fatal {
		log.Println(red + logString + reset)
	}

	if l.WriteLevel >= Fatal {
		l.SaveLogToFile(logString)
	}
}

// System print the system level logs
func (l *ULog) System(v string) {
	logString := "System: " + v
	if l.Level >= System {
		log.Println(logString)
	}

	//if l.WriteLevel >= System {
	//	l.SaveLogToFile(logString)
	//}
}

// SaveLogToFile save the log to a file
func (l *ULog) SaveLogToFile(v string) {
	if l.IsSave {
		l.fileMutex.Lock()
		defer l.fileMutex.Unlock()

		file, err := os.OpenFile(l.LogSavePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer func() {
			err := file.Close()
			if err != nil {
				//唯二不调用error级别却打印error日志的地方
				l.Warn("Close log file failed: " + err.Error())
			}
		}()

		_, err = file.Write([]byte(time.Now().Format("2006-01-02 15:04:05") + ":" + v + "\n"))
		if err != nil {
			//唯二不调用error级别却打印error日志的地方
			l.Warn("Write log to file failed: " + err.Error())
		}
	}
}
