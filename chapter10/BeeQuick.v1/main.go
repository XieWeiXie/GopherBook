package main

import (
	"log"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/cmd"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/configs"
)

var ENV string

func main() {
	if ENV == "" {
		configs.ENV = "dev"
	} else {
		configs.ENV = ENV
	}
	log.Println("Running Web Server")
	cmd.Execute()

}
