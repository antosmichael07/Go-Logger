package lgr

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type Logger struct {
	// Directory where logs will be saved
	Directory string
	// Output to console and/or file
	Output Output
	// Prints logs with level equal or higher than this
	Level Level
	// Name of the logger
	Name string
}

type Output struct {
	// Output to console
	Console bool
	// Output to file
	File bool
}

type Level int

// Log levels
const (
	Info Level = iota
	Warning
	Error
	None
)

// Creates a new logger with the given name and default values
func NewLogger(name string) Logger {
	return Logger{
		Directory: "logs",
		Output: Output{
			Console: true,
			File:    true,
		},
		Level: Info,
		Name:  name,
	}
}

// Logs a message with the given level and arguments
func (logger Logger) Log(level Level, message string, args ...interface{}) {
	// If the level is higher than the logger's level, don't log
	if logger.Level <= level {
		level_as_string := []string{"Info", "Warning", "Error", "None"}

		// Get the name of the file that called the logger
		_, caller, line, _ := runtime.Caller(1)
		for i := len(caller) - 1; i >= 0; i-- {
			if caller[i] == '/' {
				caller = caller[i+1:]
				break
			}
		}

		// Format the message
		msg := fmt.Sprintf(message, args...)
		str := fmt.Sprintf("[%s] [%s:%d] [%s/%s] %s\n", time.Now().String()[:19], caller, line, logger.Name, level_as_string[level], msg)

		// Print the message to the console
		if logger.Output.Console {
			fmt.Print(str)
		}

		// Save the message to a file
		if logger.Output.File {
			// Create the directory if it doesn't exist
			if _, err := os.Stat(fmt.Sprintf("./%s", logger.Directory)); os.IsNotExist(err) {
				os.Mkdir(fmt.Sprintf("./%s", logger.Directory), 0755)
			}

			// Open the file and write the message
			file, _ := os.OpenFile(fmt.Sprintf("./%s/%s.txt", logger.Directory, time.Now().String()[:10]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			file.WriteString(str)
			file.Close()
		}
	}
}
