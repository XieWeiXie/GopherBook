package log_for_project

import (
	"fmt"
	"log"
)

func red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

func Println(message string) {
	log.Println(fmt.Sprintf(red(message)))
}
