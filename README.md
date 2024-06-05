# Go-Logger

A simple library, that makes messages to the console more convenient to debug. The messages are saved to `logs` directory by default. You can customize it in the Logger struct.<hr>
Install with `go get github.com/antosmichael07/Go-Logger`

## Example

```go
package main

import custom_logger "github.com/antosmichael07/Go-Logger"

func main() {
	logger := custom_logger.NewLogger()

	logger.Log(custom_logger.Info, "test")

	logger.Directory = "other"
	logger.Output.Console = false

	logger.Log(custom_logger.Info, "test %d", 1)
}
```
