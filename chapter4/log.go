package chapter4

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func DefaultUsageForLog() {

	log.Print("Hello World, Golang")
	log.Println("Hello World, Golang")
	log.Printf("Hello World, %s", "Golang")

	fmt.Println("log Prefix", log.Prefix())
	fmt.Println("log Flags", log.Flags())
}

func SpecialUsageLog() {

	// 实例化 log.Logger 结构体
	logger := log.New(os.Stdout, "Golang ", log.Lshortfile)
	//
	logger.Println("Hello World, Golang")
}

func SpecialUsageWithBytes() {

	var buf bytes.Buffer
	logger := log.New(&buf, "Hi! ", log.Lshortfile)

	logger.Println("Hello World, Golang")

	fmt.Println(buf.String())
}

func SpecialUsageWithFile() {

	file, _ := os.Create("log.log")
	logger := log.New(file, "Hi!", log.Lshortfile)

	logger.Println("Hello World, Golang")
}
