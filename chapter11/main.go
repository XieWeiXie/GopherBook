package main

import (
	"GopherBook/chapter11/cmd"
)

var ENV string

func main() {
	if ENV == "" {
		ENV = "dev"
	}
	cmd.Execute()

}
