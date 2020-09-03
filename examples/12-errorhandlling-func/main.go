package main

import (
	"errors"
	"fmt"
	"github.com/go-schild/logging"
)

// Sometimes you want to defer a method - often a close function.
// But those will return an error. This error needs to be considered.
func AnyCloseMethod() error {
	return errors.New(`any error while closing something`)
}

func main() {
	// Here, the error is ignored; not good.
	defer AnyCloseMethod()

	// In this block, the error is logged, but this is too volatile!
	defer func() {
		err := AnyCloseMethod()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// Using the LogErr method, the code is a bit smaller.
	// But still to volatile in my opinion.
	defer func() {
		logging.LogErr(AnyCloseMethod())
	}()

	// This method will call the AnyCloseMethod function and call logging.LogErr with the result.
	// The downside is, it works only with functions which takes nothing and returns only an error.
	defer logging.LogErrFunc(AnyCloseMethod)
}
