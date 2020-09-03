package main

import (
	"log"
	"os"

	"github.com/go-schild/logging"
)

func init() {
	logging.Outputs = []logging.LoggerOut{
		// LoggerOut defines where the formatted log entry should be written.
		// If this output will be used, depends on the log level in the configuration,
		// the log level of the message and the log level of the  output.
		{os.Stderr, logging.LogLevelAll},
	}

	logging.EntryOutputs = []logging.LoggerEntryOut{
		// This output works like the output above, but it doesn't use a writer. It uses a function instead.
		// This is more useful for more complex handling of log events, like storing it into a database.
		// You get the raw entry data and can handle it like you wish.
		{
			LogLevel: logging.LogLevelWarning,
			HandlerFunc: func(entry logging.Entry) {
				// You could do whatever you want here.
				log.Printf(
					"%s %s %s from EntryOutputs",
					entry.Time.String(),
					entry.LogLevel.String(),
					entry.Message,
				)
			},
		},
	}
}

func main() {
	logging.Debug("this is a debug entry")
	logging.Info("this is an info entry")
	logging.Warning("this is a warning entry")
	logging.Error("this is an error entry")
}
