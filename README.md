# Go-Logger

A simple library, that makes messages to the console more convenient to debug. The messages are saved to `logs` directory by default. You can customize it in the Logger struct.<hr>
Install with `go get github.com/antosmichael07/Go-Logger`

## Example

```go
package main

import logger "github.com/antosmichael07/Go-Logger"

func main() {
	logger.Log("Hello World")
}
```
