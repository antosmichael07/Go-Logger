package custom_logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type Logger struct {
	Directory string
}

func NewLogger() Logger {
	return Logger{
		Directory: "logs",
	}
}

func (logger Logger) Log(message string) {
	_, caller, line, _ := runtime.Caller(1)
	for i := len(caller) - 1; i >= 0; i-- {
		if caller[i] == '/' {
			caller = caller[i+1:]
			break
		}
	}

	str := fmt.Sprintf("[%s] [%s:%d] %s", time.Now().String()[:19], caller, line, message)

	fmt.Println(str)

	if _, err := os.Stat(fmt.Sprintf("./%s", logger.Directory)); os.IsNotExist(err) {
		os.Mkdir(fmt.Sprintf("./%s", logger.Directory), 0755)
	}

	file, _ := os.OpenFile(fmt.Sprintf("./%s/%s.txt", logger.Directory, time.Now().String()[:10]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file.WriteString(str + "\n")
	file.Close()
}
