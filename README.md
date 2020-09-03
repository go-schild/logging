# go-schild logging library

This is a small logging library developed with the goals of simplicity, configurability and being dependency free.

## Log levels

There are 4 log levels: `Error`, `Warning`, `Info` and `Debug`.
There are also 2 special log levels: `None` and `All`.

The logging mechanism uses those log level to determine, which log entries should be written.
The log levels take place on the following places:

- The configuration, defines the log level for the whole software
- The output, e.g. when an output should only be used to write error log entries
- The entry, e.g. when you use logging.Warning, you create an entry with the warning log level.

When logging, the log entries are first filtered by the configured log level, then by the output log level.
For example, if you write the output to stdout with the log level Warning and you have set the log level in the
configuration to Info, a Debug entry will not be logged, because it will be filtered out of the configuration and
a Info entry won't be logged because there is no output. An Error log entry will be logged to stdout, because it
fits the output log level (Warning and lower).

The default for the set stdout output and the configuration is `All`, which is equal to `Debug`.
This means, all log entries will be written.

## A small example

    package main
    
    import (
        "github.com/go-schild/logging"
        "io/ioutil"
    )
    
    func main() {
        logging.Info("Example started.")
    
        filename := "reading_some_file.txt"
        data, err := ioutil.ReadFile(filename)
        logging.FatalErr(err)
        logging.InfoF("%s: %d bytes", filename, len(data))
        
        logging.Info("Example stopped.")
    }

## More Examples

I spend some time to document some examples for an easy entry into this library.

- Basics
    - [01-default](/go-schild/logging/blob/master/examples/01-default/main.go)
        A basic example, how you create simple log entries
    - [02-formatting](/go-schild/logging/blob/master/examples/02-formatting/main.go)
        Using the format function to create log entries
- Error logging
    - [11-errorhandling](/go-schild/logging/blob/master/examples/11-errorhandling/main.go)
        An example, how you log error messages 
    - [12-errorhandling-func](/go-schild/logging/blob/master/examples/12-errorhandling-func/main.go)
        This example uses the func methods, designed to be used in defer calls.
- Configuration
    - [21-config](/go-schild/logging/blob/master/examples/21-config/main.go)
        A basic configuration example
    - [22-outputs](/go-schild/logging/blob/master/examples/22-outputs/main.go)
        Defining custom outputs