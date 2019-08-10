package test

import (
	"GopherBook/chapter12/example/fina"
	"log"
	"os"
	"testing"
)

func TestParseCss(test *testing.T) {
	file, err := os.Open("content.html")
	if err != nil {
		log.Println(err)
		return

	}
	fina.ParseByCss(file)

}
func TestParseXpath(test *testing.T) {
	file, err := os.Open("content.html")
	if err != nil {
		log.Println(err)
		return

	}
	fina.ParseByXpath(file)
}

func TestParseJson(test *testing.T) {
	file, err := os.OpenFile("content.json", os.O_RDWR, 650)
	if err != nil {
		log.Println(err)
		return
	}
	fina.ParseByJson(file)
}
