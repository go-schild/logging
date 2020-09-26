package logging

import "os"

var (
	// Level defines the current log level
	Level = LogLevelAll

	// Outputs contains LoggerOut instances. Those filter messages by the log level and writes them to the defined
	// writer. As default only Stdout is added as output with the loglevel "ALL".
	Outputs = []LoggerOut{{os.Stdout, LogLevelAll}}

	// EntryOutputs defines a list of handlers for handling Entrys instead of the already formatted string.
	// As en example those handlers can be used to store the log entry in a database or as an json object.
	EntryOutputs []LoggerEntryOut

	// TimeFormat defines how the time should be formatted.
	// The time formatting of Go is used.
	// It will be available in the Format string as {{.Time}}
	TimeFormat = "2006-01-02 15:04:05"

	// Format is a string used to format an entry into an string, which will be written to the Outputs.
	Format = "[{{.Time}}] [{{.Level}}] {{.Message}}"
)
