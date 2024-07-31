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
	// File where logs will be saved
	File *os.File
	// Name of the file where logs will be saved
	OpenedFile string
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
func NewLogger(name string, dir string, file_out bool) (logger Logger, err error) {
	file := os.Stdout
	if file_out {
		// Create the directory if it doesn't exist
		var err error
		if _, err = os.Stat(fmt.Sprintf("./%s", dir)); os.IsNotExist(err) {
			os.Mkdir(fmt.Sprintf("./%s", dir), 0755)
		}
		// Open the file where logs will be saved
		if file, err = os.OpenFile(fmt.Sprintf("./%s/%s.txt", dir, time.Now().String()[:10]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			return Logger{
				Directory: dir,
				Output: Output{
					Console: true,
					File:    file_out,
				},
				Level:      Info,
				Name:       name,
				File:       file,
				OpenedFile: time.Now().String()[:10],
			}, err
		}
	}

	return Logger{
		Directory: dir,
		Output: Output{
			Console: true,
			File:    file_out,
		},
		Level:      Info,
		Name:       name,
		File:       file,
		OpenedFile: time.Now().String()[:10],
	}, nil
}

// Closes the file where logs are saved
func (logger *Logger) CloseFileOutput() (err error) {
	logger.Output.File = false
	logger.OpenedFile = ""
	return logger.File.Close()
}

// Opens the file where logs are saved
func (logger *Logger) OpenFileOutput() (err error) {
	time := time.Now().String()[:19]
	if logger.File, err = os.OpenFile(fmt.Sprintf("./%s/%s.txt", logger.Directory, time[:10]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return err
	}
	logger.OpenedFile = time[:10]
	logger.Output.File = true
	return nil
}

// Logs a message with the given level and arguments
func (logger *Logger) Log(level Level, message string, args ...interface{}) (err error) {
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
		time := time.Now().String()[:19]
		msg := fmt.Sprintf(message, args...)
		str := fmt.Sprintf("[%s] [%s:%d] [%s/%s] %s\n", time, caller, line, logger.Name, level_as_string[level], msg)

		// Print the message to the console
		if logger.Output.Console {
			fmt.Print(str)
		}

		// Save the message to a file
		if logger.Output.File {
			// If the date has changed, close the current file and open a new one
			if logger.OpenedFile != time[:10] {
				logger.File.Close()
				if file, err := os.OpenFile(fmt.Sprintf("./%s/%s.txt", logger.Directory, time[:10]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
					return err
				} else {
					logger.File = file
					logger.OpenedFile = time[:10]
				}
			}
			// Write the message to the file
			_, err := logger.File.WriteString(str)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
