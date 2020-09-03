package main

import "github.com/go-schild/logging"

func main() {
	// you know this variable from example 01-default.
	// This defines the current log level.
	// The default is ALL, but you can set it depending on you own needs.
	logging.Level = logging.LogLevelAll

	// This variable contains output configurations. They define how to handle logging output.
	// For more information, see example 22-output.
	logging.Outputs = nil

	// Those are the same like logging.Outputs, but they don't handle the already formatted string,
	// they handle the underlying Entry structs.
	// For more information, see example 22-output.
	logging.EntryOutputs = nil

	// Use this variable to configure how the time should be formatted.
	// This uses the golang time format definition.
	// e.g. 2006-05-04 for yyyy-mm-dd
	logging.TimeFormat = "01:02:03"

	// This variable is the format, how Entry structs should be formatted to strings.
	// It uses the golang template syntax and the Entry is the given data struct.
	logging.Format = "[{{.Time}}] [{{.Level}}] {{.Message}}"
}
