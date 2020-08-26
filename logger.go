package logging

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func interfacesToString(parts []interface{}) string {
	var res []string

	for _, part := range parts {
		switch part.(type) {
		case string:
			res = append(res, part.(string))
		case int:
			res = append(res, fmt.Sprintf("%d", part.(int)))
		case float32:
			res = append(res, fmt.Sprintf("%f", part.(float32)))
		case float64:
			res = append(res, fmt.Sprintf("%f", part.(float64)))
		case fmt.Stringer:
			res = append(res, part.(fmt.Stringer).String())
		case error:
			res = append(res, part.(error).Error())
		default:
			res = append(res, reflect.TypeOf(part).Name())
		}
	}

	return strings.Join(res, " ")
}

func Log(level LogLevel, msg ...interface{}) {
	LogEntry(level, Entry{
		Time:     time.Now(),
		LogLevel: level,
		Message:  interfacesToString(msg),
	})
}

func LogEntry(level LogLevel, entry Entry) {
	if Level >= level {
		for _, out := range Outputs {
			out.Write(entry.String(), level)
		}
		for _, out := range EntryOutputs {
			if out.LogLevel >= level {
				out.HandlerFunc(entry)
			}
		}
	}
}

func Error(msg ...interface{}) {
	Log(LogLevelError, msg...)
}

func ErrorF(format string, a ...interface{}) {
	Log(LogLevelError, fmt.Sprintf(format, a...))
}

func Warning(msg ...interface{}) {
	Log(LogLevelWarning, msg...)
}

func WarningF(format string, a ...interface{}) {
	Log(LogLevelWarning, fmt.Sprintf(format, a...))
}

func Info(msg ...interface{}) {
	Log(LogLevelInfo, msg...)
}

func InfoF(format string, a ...interface{}) {
	Log(LogLevelInfo, fmt.Sprintf(format, a...))
}

func Debug(msg ...interface{}) {
	Log(LogLevelDebug, msg...)
}

func DebugF(format string, a ...interface{}) {
	Log(LogLevelDebug, fmt.Sprintf(format, a...))
}
