package logging

// LogLevel describes how many information should be logged.
type LogLevel int

func (l LogLevel) String() string {
	name := LogLevelNames[l]
	if name != "" {
		return name
	}
	return "UNKNOWN"
}

const (
	LogLevelNone LogLevel = iota
	LogLevelError
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug

	LogLevelAll = LogLevelDebug
)

var (
	// LogLevelNames defines how the log level should be named in the output.
	// You can change this map depending on your needs.
	// If a log level is not mentioned within this map, the log level will be translated to UNKNOWN.
	// The log level name will be available for the format string as {{.Level}}
	LogLevelNames = map[LogLevel]string{
		LogLevelNone:    "NONE",
		LogLevelError:   "ERROR",
		LogLevelWarning: "WARNING",
		LogLevelInfo:    "INFO",
		LogLevelDebug:   "DEBUG",
	}
)
