package main

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter11/cmd"
)

var ENV string

func main() {
	if ENV == "" {
		ENV = "dev"
	}
	cmd.Execute()

}
