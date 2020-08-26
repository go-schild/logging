package main

import (
	"log"
	"os"

	"github.com/go-schild/logging"
)

func configOutputs() {
	logging.Outputs = []logging.LoggerOut{
		// Log all entries into stderr
		{os.Stderr, logging.LogLevelAll},
	}

	logging.EntryOutputs = []logging.LoggerEntryOut{
		// Log warnings and errors into stdout with a custom handler
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
	configOutputs()

	logging.Debug("this is a debug entry")
	logging.Info("this is an info entry")
	logging.Warning("this is a warning entry")
	logging.Error("this is an error entry")
}
