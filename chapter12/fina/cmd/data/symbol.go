package data

import (
	"GopherBook/chapter12/fina/pkg/assistance"
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func RunForSymbol(url string) (bool, error) {
	content, err := assistance.Downloader(url)
	if err != nil {
		return false, err
	}
	fmt.Println(string(content))
	return true, nil
}

func ParseByHtml(content string) (bool, error) {
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		log.Println("html parse fail", err.Error())
		return false, err
	}
	var f = func(node *html.Node) {
		if node.Type == html.ElementNode {
		}
	}
	f(doc)
	return true, nil

}
