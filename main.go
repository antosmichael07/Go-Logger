package logger

import (
	"fmt"
	"os"
	"time"
)

func Log(location string, message string) {
	str := fmt.Sprintf("[%s] [%s] %s", time.Now().String()[:19], location, message)

	fmt.Println(str)

	file, _ := os.OpenFile(fmt.Sprintf("./logs/%s.txt", time.Now().String()[:10]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file.WriteString(str + "\n")
	file.Close()
}
