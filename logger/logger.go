package logger

import (
	"time"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

func (l LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[l]
}

type Logger struct {
	sinks    []Sink
	minLevel LogLevel
}

func NewLogger(minLevel LogLevel, sinks ...Sink) *Logger {
	return &Logger{sinks: sinks, minLevel: minLevel}
}

func (l *Logger) Log(level LogLevel, message string, fields map[string]interface{}) {
	if level < l.minLevel {
		return
	}
	entry := LogEntry{
		Level:   level.String(),
		Time:    time.Now().Format(time.RFC3339),
		Message: message,
		Fields:  fields,
	}
	for _, sink := range l.sinks {
		_ = sink.Write(entry)
	}
}

func (l *Logger) Close() {
	for _, sink := range l.sinks {
		_ = sink.Close()
	}
}
