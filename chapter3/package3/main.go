package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello World")
	fmt.Print(errors.New(""))
	log.Println()
}
