package main

import (
	"html/template"
	"log"
	"os"

	"github.com/gobuffalo/packr"
)

type Detail struct {
	Class   string
	Title   string
	Content string
}

var detail Detail

func init() {
	detail = Detail{
		Class:   "test-class",
		Title:   "test-title",
		Content: "test-content",
	}
}

func originMethod() {
	pwd, _ := os.Getwd()
	t, err := template.ParseFiles(pwd + "/chapter5/template/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = t.Execute(os.Stdout, &detail)
	if err != nil {
		log.Println(err)
		return
	}
}

func withPackr() {
	// 相对目录
	box := packr.NewBox(".")
	index, err := box.FindString("index.html")
	if err != nil {
		log.Println(err)
		return
	}
	t, _ := template.New("").Parse(index)
	err = t.Execute(os.Stdout, &detail)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	originMethod()
	withPackr()
}
