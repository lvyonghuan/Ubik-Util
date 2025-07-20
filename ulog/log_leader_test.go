package ulog

import (
	"bytes"
	"log"
	"testing"
)

func TestRecordFollowerLogValidDebugLevel(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	leaderLog := NewLeaderLog(Debug, false, "")
	leaderLog.RecordFollowerLog("follower-1", "Debug message", Debug)

	if !bytes.Contains(buf.Bytes(), []byte("Follower UUID: follower-1 , message: Debug message")) {
		t.Errorf("Expected log message not found in output")
	}
}

func TestRecordFollowerLogInvalidLogLevel(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	leaderLog := NewLeaderLog(Debug, false, "")
	leaderLog.RecordFollowerLog("follower-2", "Invalid level message", 999)

	if !bytes.Contains(buf.Bytes(), []byte("Invalid log level for follower log record")) {
		t.Errorf("Expected warning for invalid log level not found in output")
	}
}

func TestRecordFollowerLogLowerLogLevelIgnored(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	leaderLog := NewLeaderLog(Info, false, "")
	leaderLog.RecordFollowerLog("follower-3", "Debug message", Debug)

	if bytes.Contains(buf.Bytes(), []byte("Follower UUID: follower-3 , message: Debug message")) {
		t.Errorf("Unexpected log message found in output for lower log level")
	}
}

func TestRecordFollowerLogValidFatalLevel(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	leaderLog := NewLeaderLog(Fatal, false, "")
	leaderLog.RecordFollowerLog("follower-4", "Fatal message", Fatal)

	if !bytes.Contains(buf.Bytes(), []byte("Follower UUID: follower-4 , message: Fatal message")) {
		t.Errorf("Expected log message not found in output")
	}
}
