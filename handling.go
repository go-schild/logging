package logging

import "os"

func LogWarn(err error) {
	if err != nil {
		Warning(err.Error())
	}
}

func LogErr(err error) {
	if err != nil {
		Error(err.Error())
	}
}

func PanicErr(err error) {
	if err != nil {
		Error(err.Error())
		panic(err)
	}
}

func FatalErr(err error) {
	if err != nil {
		Error(err.Error())
		os.Exit(1)
	}
}

func LogWarnFunc(errFunc func() error) {
	LogWarn(errFunc())
}

func LogErrFunc(errFunc func() error) {
	LogErr(errFunc())
}

func PanicErrFunc(errFunc func() error) {
	PanicErr(errFunc())
}

func FatalErrFunc(errFunc func() error) {
	FatalErr(errFunc())
}
