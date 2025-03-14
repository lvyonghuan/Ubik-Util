package ulog

import (
	"log"
	"os"
	"sync"
	"time"
)

type Log struct {
	Level       int        `json:"level"`     //log level
	IsSave      bool       `json:"is_save"`   //whether to save logs
	LogSavePath string     `json:"save_path"` //the path where the logs are saved
	fileMutex   sync.Mutex //文件互斥访问
}

// 日志等级
const (
	Off = iota
	Fatal
	Error
	Warn
	Info //默认级别
	Debug
)

// 日志颜色
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	orange = "\033[33m"
	yellow = "\033[93m"
	green  = "\033[32m"
)

// InitLog init log
func (l *Log) InitLog() {
	if l.IsSave {
		currentTime := time.Now().Format("_2006-01-02 15-04-05")
		l.LogSavePath = l.LogSavePath + currentTime + ".log"
	}
}

// Debug print debug level logs
func (l *Log) Debug(v string) {
	if l.Level >= Debug {
		logString := "Debug: " + v
		log.Println(green + logString + reset)
		l.SaveLogToFile(logString)
	}
}

// Info print info level logs
func (l *Log) Info(v string) {
	if l.Level >= Info {
		logString := v
		log.Println(logString)
		l.SaveLogToFile(logString)
	}
}

// Warn print the warn level logs
func (l *Log) Warn(v string) {
	if l.Level >= Warn {
		logString := "Warn: " + v
		log.Println(yellow + logString + reset)
		l.SaveLogToFile(logString)
	}
}

// Error print the error level log
func (l *Log) Error(v error) {
	if l.Level >= Error {
		logString := "Error: " + v.Error()
		log.Println(orange + logString + reset)
		l.SaveLogToFile(logString)
	}
}

// Fatal print the fatal level logs
func (l *Log) Fatal(v error) {
	if l.Level >= Fatal {
		logString := "Fatal: " + v.Error()
		log.Println(red + logString + reset)
		l.SaveLogToFile(logString)
	}
}

// SaveLogToFile save the log to a file
func (l *Log) SaveLogToFile(v string) {
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
