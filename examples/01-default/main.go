package main

import (
	"github.com/go-schild/logging"
	"io/ioutil"
)

func main() {
	// You can change the log level by overwriting this variable
	logging.Level = logging.LogLevelAll

	// Just write you log message here. It gets filtered out by the log level automatically.
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

func main() {
	logging.Info("Example started.")

	filename := "reading_some_file.txt"
	data, err := ioutil.ReadFile(filename)
	logging.FatalErr(err)
	logging.InfoF("%s: %d bytes", filename, len(data))

	logging.Info("Example stopped.")
}