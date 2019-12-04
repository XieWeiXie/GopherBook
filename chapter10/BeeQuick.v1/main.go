package main

import (
	"log"

	"GopherBook/Chapter5/BeeQuick.v1/cmd"
	"GopherBook/Chapter5/BeeQuick.v1/configs"
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
