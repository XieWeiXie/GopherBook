package node

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func ParseByHtmlNode() {
	file, err := os.Open("node.html")
	if err != nil {
		log.Println(err)
		return
	}
	doc, err := html.Parse(file)
	if err != nil {
		log.Println(err)
		return
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode && n.Parent.Type == html.ElementNode &&
			n.Parent.Data == "p" {
			fmt.Println(n.Data)
			for _, i := range n.Parent.Attr {
				fmt.Println(i.Key, i.Val)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func replacer(value string) string {
	replacer := strings.NewReplacer("\n", "", "\t", "", " ", "")
	return strings.TrimSpace(replacer.Replace(value))
}

func ParseText() {
	file, err := os.Open("node.html")
	if err != nil {
		log.Println(err)
		return
	}
	z := html.NewTokenizer(file)
	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.TextToken:
			if depth > 0 {
				fmt.Println(string(z.Text()))
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'p' {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}
			}
		}
	}
}
