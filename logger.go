package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	LOG_LEVEL_OFF = iota
	LOG_LEVEL_FATAL
	LOG_LEVEL_ERROR
	LOG_LEVEL_WARN
	LOG_LEVEL_INFO
	LOG_LEVEL_DEBUG
	LOG_LEVEL_TRACE
	LOG_LEVEL_ALL
)

var DefaultLevel int = LOG_LEVEL_INFO

type Logger struct {
	level  int
	prefix string
}

func init() {
	if os.Getenv("DEBUG") != "" {
		DefaultLevel = LOG_LEVEL_DEBUG
	}

	switch os.Getenv("LOG_LEVEL") {
	case "OFF":
		DefaultLevel = LOG_LEVEL_OFF
	case "FATAL":
		DefaultLevel = LOG_LEVEL_FATAL
	case "ERROR":
		DefaultLevel = LOG_LEVEL_ERROR
	case "WARN":
		DefaultLevel = LOG_LEVEL_WARN
	case "INFO":
		DefaultLevel = LOG_LEVEL_INFO
	case "DEBUG":
		DefaultLevel = LOG_LEVEL_DEBUG
	case "TRACE":
		DefaultLevel = LOG_LEVEL_TRACE
	case "ALL":
		DefaultLevel = LOG_LEVEL_ALL
	}
}

func NewLogger(level int, prefix string) *Logger {
	return &Logger{
		level:  level,
		prefix: prefix,
	}
}

func NewDefaultLogger(prefix string) *Logger {
	return &Logger{
		level:  DefaultLevel,
		prefix: prefix,
	}
}

func (l *Logger) SetLevel(level int) {
	if level >= LOG_LEVEL_OFF && level <= LOG_LEVEL_ALL {
		l.level = level
	}
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	if l.level >= LOG_LEVEL_FATAL {
		l.Log("[fatal]", format, v...)
	}
	os.Exit(1)
}

func (l *Logger) Error(format string, v ...interface{}) {
	if l.level >= LOG_LEVEL_ERROR {
		l.Log("[error]", format, v...)
	}
}

func (l *Logger) Warn(format string, v ...interface{}) {
	if l.level >= LOG_LEVEL_WARN {
		l.Log("[warn] ", format, v...)
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l.level >= LOG_LEVEL_INFO {
		l.Log("[info] ", format, v...)
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if l.level >= LOG_LEVEL_DEBUG {
		l.Log("[debug]", format, v...)
	}
}

func (l *Logger) Trace(format string, v ...interface{}) {
	if l.level >= LOG_LEVEL_TRACE {
		l.Log("[trace]", format, v...)
	}
}

func (l *Logger) Log(level string, format string, v ...interface{}) {
	tim := time.Now().Format("2006-01-02 15:04:05.000000")
	out := fmt.Sprintf(fmt.Sprintf("%s\n", format), v...)

	if l.prefix != "" {
		fmt.Printf("[%s] %s (%s) %s", tim, level, l.prefix, out)
	} else {
		fmt.Printf("[%s] %s %s", tim, level, out)
	}
}
