package main

import "github.com/go-schild/logging"

// nilErr is just a dummy func which returns an error.
func nilError() error {
	return nil
}

func main() {
	// This will just create a warning entry with the error inside
	err := nilError()
	logging.LogWarn(err)

	// This will create an error entry instead
	err = nilError()
	logging.LogErr(err)

	// This will log the error and panic, when the error is not nil
	err = nilError()
	logging.PanicErr(err)

	// This will log the error and exit the application
	err = nilError()
	logging.FatalErr(err)
}
