package ulog

import (
	"fmt"
	"log"
)

type LeaderLogInterface interface {
	// Log It needs to implement a Log interface
	Log
	// RecordFollowerLog It needs to record follower's log message
	// Input follower's UUID, the log message which sent by follower and the level of the log
	RecordFollowerLog(followerUUID, message string, level int)
}

type LeaderLog struct {
	log *LogWithoutPost
}

func NewLeaderLog(level int, isSave bool, logSavePath string) LeaderLog {
	l := &LogWithoutPost{
		Level:       level,
		IsSave:      isSave,
		LogSavePath: logSavePath,
	}

	leaderLog := LeaderLog{
		log: l,
	}
	leaderLog.InitLog()

	return leaderLog
}

func (l LeaderLog) InitLog() {
	l.log.InitLog()
}

func (l LeaderLog) Debug(v string) {
	l.log.Debug(v)
}

func (l LeaderLog) Info(v string) {
	l.log.Info(v)
}

func (l LeaderLog) Warn(v string) {
	l.log.Warn(v)
}

func (l LeaderLog) Error(v error) {
	l.log.Error(v)
}

func (l LeaderLog) Fatal(v error) {
	l.log.Fatal(v)
}

func (l LeaderLog) SaveLogToFile(v string) {
	l.log.SaveLogToFile(v)
}

func (l LeaderLog) RecordFollowerLog(followerUUID, message string, level int) {
	logString := fmt.Sprintf("Follower UUID: %s , message: %s", followerUUID, message)

	switch level {
	case Debug:
		if l.log.Level >= Debug {
			log.Println(green + logString + reset)
			l.SaveLogToFile(logString)
		}
	case Info:
		if l.log.Level >= Info {
			log.Println(logString)
			l.SaveLogToFile(logString)
		}
	case Warn:
		if l.log.Level >= Warn {
			log.Println(yellow + logString + reset)
			l.SaveLogToFile(logString)
		}
	case Error:
		if l.log.Level >= Error {
			log.Println(orange + logString + reset)
			l.SaveLogToFile(logString)
		}
	case Fatal:
		if l.log.Level >= Fatal {
			log.Println(red + logString + reset)
			l.SaveLogToFile(logString)
		}
	default:
		l.log.Warn("Invalid log level for follower log record: " + followerUUID + ", level: " + fmt.Sprint(level) + ", message: " + message)
	}
}
