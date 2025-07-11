package ulog

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/lvyonghuan/Ubik-Util/uconst"
	"github.com/lvyonghuan/Ubik-Util/umessenger"
)

type Log interface {
	InitLog()
	Debug(v string)
	Info(v string)
	Warn(v string)
	Error(v error)
	Fatal(v error)
	SaveLogToFile(v string)
}

// LogWithPost will send logs to the leader
type LogWithPost struct {
	Level       int        `json:"level"`     //log level
	IsSave      bool       `json:"is_save"`   //whether to save logs
	LogSavePath string     `json:"save_path"` //the path where the logs are saved
	fileMutex   sync.Mutex // mutex for file operations

	messenger *umessenger.UMessenger //To send logs to the leader
}

// NewLogWithPost creates a new LogWithPost instance, initializes it, and returns it.
func NewLogWithPost(level int, isSave bool, logSavePath, leaderAddr, uuid string) *LogWithPost {
	logWithPost := &LogWithPost{
		Level:       level,
		IsSave:      isSave,
		LogSavePath: logSavePath,
	}

	logWithPost.InitLog(leaderAddr, uuid)
	return logWithPost
}

type LogWithoutPost struct {
	Level       int        `json:"level"`     //log level
	IsSave      bool       `json:"is_save"`   //whether to save logs
	LogSavePath string     `json:"save_path"` //the path where the logs are saved
	fileMutex   sync.Mutex // mutex for file operations
}

// NewLogWithoutPost creates a new LogWithoutPost instance, initializes it, and returns it.
func NewLogWithoutPost(level int, isSave bool, logSavePath string) *LogWithoutPost {
	logWithoutPost := &LogWithoutPost{
		Level:       level,
		IsSave:      isSave,
		LogSavePath: logSavePath,
	}

	logWithoutPost.InitLog()
	return logWithoutPost
}

// log levels
const (
	Off   = 0
	Fatal = uconst.Fatal
	Error = uconst.Error
	Warn  = uconst.Warn
	Info  = uconst.Info //Default log level
	Debug = uconst.Debug
)

// log colors
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	orange = "\033[33m"
	yellow = "\033[93m"
	green  = "\033[32m"
)

const leaderLogPath = "/follower/log/"

// InitLog init log
func (l *LogWithPost) InitLog(leaderAddr, uuid string) {
	if l.IsSave {
		currentTime := time.Now().Format("_2006-01-02 15-04-05")
		l.LogSavePath = l.LogSavePath + currentTime + ".log"
	}

	// Initialize the messenger
	l.messenger = umessenger.NewUMessenger(leaderAddr+leaderLogPath, uuid)
}

// Debug print debug level logs
func (l *LogWithPost) Debug(v string) {
	if l.Level >= Debug {
		logString := "Debug: " + v
		log.Println(green + logString + reset)
		l.SaveLogToFile(logString)
		l.messenger.PostLog(v, umessenger.Debug) // Send debug log to the leader
	}
}

// Info print info level logs
func (l *LogWithPost) Info(v string) {
	if l.Level >= Info {
		logString := v
		log.Println(logString)
		l.SaveLogToFile(logString)
		l.messenger.PostLog(v, umessenger.Info) // Send info log to the leader
	}
}

// Warn print the warn level logs
func (l *LogWithPost) Warn(v string) {
	if l.Level >= Warn {
		logString := "Warn: " + v
		log.Println(yellow + logString + reset)
		l.SaveLogToFile(logString)
		l.messenger.PostLog(v, umessenger.Warn) // Send warn log to the leader
	}
}

// Error print the error level log
func (l *LogWithPost) Error(v error) {
	if l.Level >= Error {
		logString := "Error: " + v.Error()
		log.Println(orange + logString + reset)
		l.SaveLogToFile(logString)
		l.messenger.PostLog(v.Error(), umessenger.Error) // Send error log to the leader
	}
}

// Fatal print the fatal level logs
func (l *LogWithPost) Fatal(v error) {
	if l.Level >= Fatal {
		logString := "Fatal: " + v.Error()
		log.Println(red + logString + reset)
		l.SaveLogToFile(logString)
		l.messenger.PostLog(v.Error(), umessenger.Fatal) // Send fatal log to the leader
	}
}

// SaveLogToFile save the log to a file
func (l *LogWithPost) SaveLogToFile(v string) {
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

// InitLog init log
func (l *LogWithoutPost) InitLog() {
	if l.IsSave {
		currentTime := time.Now().Format("_2006-01-02 15-04-05")
		l.LogSavePath = l.LogSavePath + currentTime + ".log"
	}
}

// Debug print debug level logs
func (l *LogWithoutPost) Debug(v string) {
	if l.Level >= Debug {
		logString := "Debug: " + v
		log.Println(green + logString + reset)
		l.SaveLogToFile(logString)
	}
}

// Info print info level logs
func (l *LogWithoutPost) Info(v string) {
	if l.Level >= Info {
		logString := v
		log.Println(logString)
		l.SaveLogToFile(logString)
	}
}

// Warn print the warn level logs
func (l *LogWithoutPost) Warn(v string) {
	if l.Level >= Warn {
		logString := "Warn: " + v
		log.Println(yellow + logString + reset)
		l.SaveLogToFile(logString)
	}
}

// Error print the error level log
func (l *LogWithoutPost) Error(v error) {
	if l.Level >= Error {
		logString := "Error: " + v.Error()
		log.Println(orange + logString + reset)
		l.SaveLogToFile(logString)
	}
}

// Fatal print the fatal level logs
func (l *LogWithoutPost) Fatal(v error) {
	if l.Level >= Fatal {
		logString := "Fatal: " + v.Error()
		log.Println(red + logString + reset)
		l.SaveLogToFile(logString)
	}
}

// SaveLogToFile save the log to a file
func (l *LogWithoutPost) SaveLogToFile(v string) {
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
