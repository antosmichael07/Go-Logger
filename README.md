# Go-Logger

A simple library, that makes messages to the console more convenient to debug. The messages are saved to `logs` directory by default. You can customize it in the Logger struct.<hr>
Install with `go get github.com/antosmichael07/Go-Logger`

## Example

```go
package main

import lgr "github.com/antosmichael07/Go-Logger"

func main() {
	// We specify the name, directory where the logs are saved and if the logs are saved in a directory, when creating the logger
	logger := lgr.NewLogger("example", "logs", true)

	// The log function prints into a file and into the console with the level of info, warning or error and the message, with the date, time and the location where it was called
	// You can also put arguments at the end, like in the printf function
	logger.Log(lgr.Info, "test %d", 1)

	// We can customize where to put the logs
	logger.Directory = "other"
	// We can customize if we want to print them to the console or into files
	logger.Output.Console = false
	logger.CloseFileOutput()
	// We can customize which levels do we want to print
	// For example this would print only the logs with the level warning or higher, which means it would log only warnings and errors
	logger.Level = lgr.Warning
}
```
