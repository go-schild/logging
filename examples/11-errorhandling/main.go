package main

import "github.com/go-schild/logging"

// nilErr is just a dummy func which returns an error.
func nilError() error {
	return nil
}

func main() {
	// You can use those methods to log errors.
	// When the error is nil (like in this example) the methods won't do anything.

	// This will just create a warning entry with the error inside.
	err := nilError()
	logging.LogWarn(err)

	// This will create an error entry instead.
	err = nilError()
	logging.LogErr(err)

	// This will log the error and panic.
	err = nilError()
	logging.PanicErr(err)

	// This will log the error and exit the application with exit code 1.
	err = nilError()
	logging.FatalErr(err)
}
