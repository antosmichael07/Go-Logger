package custom_logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type Logger struct {
	Directory string
	Output    Output
	Level     Level
	Name      string
}

type Output struct {
	Console bool
	File    bool
}

type Level int

const (
	Info Level = iota
	Warining
	Error
	None
)

func NewLogger() Logger {
	return Logger{
		Directory: "logs",
		Output: Output{
			Console: true,
			File:    true,
		},
		Level: Info,
		Name: "",
	}
}

func (logger Logger) Log(level Level, message string, args ...interface{}) {
	if logger.Level >= level {
		level_as_string := []string{"Info", "Warning", "Error", "None"}

		_, caller, line, _ := runtime.Caller(1)
		for i := len(caller) - 1; i >= 0; i-- {
			if caller[i] == '/' {
				caller = caller[i+1:]
				break
			}
		}

		msg := fmt.Sprintf(message, args...)
		str := fmt.Sprintf("[%s] [%s:%d] [%s/%s] %s\n", time.Now().String()[:19], caller, line, logger.Name, level_as_string[level], msg)

		if logger.Output.Console {
			fmt.Print(str)
		}

		if logger.Output.File {
			if _, err := os.Stat(fmt.Sprintf("./%s", logger.Directory)); os.IsNotExist(err) {
				os.Mkdir(fmt.Sprintf("./%s", logger.Directory), 0755)
			}

			file, _ := os.OpenFile(fmt.Sprintf("./%s/%s.txt", logger.Directory, time.Now().String()[:10]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			file.WriteString(str)
			file.Close()
		}
	}
}
