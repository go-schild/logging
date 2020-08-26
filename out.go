package logging

import "io"

type LoggerOut struct {
	Writer   io.Writer
	LogLevel LogLevel
}

func (l LoggerOut) Write(msg string, level LogLevel) {
	if level <= l.LogLevel {
		_, _ = l.Writer.Write([]byte(msg))
	}
}

type LoggerEntryOut struct {
	HandlerFunc EntryFunc
	LogLevel    LogLevel
}
