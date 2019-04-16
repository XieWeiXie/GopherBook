package chapter4

import (
	"fmt"
	"io/ioutil"
	"os"
)

func OsUsage() {
	// 判断文件是否存在，获取文件信息
	fileMode, err := os.Stat("log.log")
	if os.IsNotExist(err) {
		return
	}
	fmt.Println(fileMode.Name(), fileMode.Mode(), fileMode.Size())
}

func OSUsageWith() {

	file, _ := os.OpenFile("os.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fmt.Println(file.Name())
	file.WriteString("Hello")
	file.WriteString("HelloWorld")
	ioutil.ReadFile("os.log")
}
