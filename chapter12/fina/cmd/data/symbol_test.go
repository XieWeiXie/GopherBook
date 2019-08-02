package data

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestHtmlNoe(tests *testing.T) {
	content, _ := os.Open("symbol.html")
	doc, _ := html.Parse(content)
	var f func(node *html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "div" {
			fmt.Println(node.Attr)

		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
