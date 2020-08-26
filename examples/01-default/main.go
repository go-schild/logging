package main

import "github.com/go-schild/logging"

func main() {
	logging.Level = logging.LogLevelAll

	logging.Debug("This is a debug message")
	logging.Info("This is an info message")
	logging.Warning("This is a warning message")
	logging.Error("This is an error message")

	logging.Level = logging.LogLevelWarning
	logging.Debug("This is not displayed")
	logging.Info("This neither")
	logging.Warning("But this is displayed")
	logging.Error("And this")
}
