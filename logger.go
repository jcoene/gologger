package logger

import (
	"fmt"
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

type Logger struct {
	level  int
	prefix string
}

func NewLogger(level int, prefix string) *Logger {
	return &Logger{
		level:  level,
		prefix: prefix,
	}
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	if l.level >= LOG_LEVEL_FATAL {
		l.Log("[fatal]", format, v...)
	}
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
