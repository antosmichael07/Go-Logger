package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func Log(message string) {
	_, caller, line, _ := runtime.Caller(1)
	str := fmt.Sprintf("[%s] [%s:%d] %s", time.Now().String()[:19], caller, line, message)

	fmt.Println(str)

	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.Mkdir("./logs", 0755)
	}

	file, _ := os.OpenFile(fmt.Sprintf("./logs/%s.txt", time.Now().String()[:10]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file.WriteString(str + "\n")
	file.Close()
}
