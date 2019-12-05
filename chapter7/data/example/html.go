package example

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func ParseByHtmlNode() {
	file, err := os.Open("node.html")
	if err != nil {
		log.Println(err)
		return
	}
	// 构造 DOM 树
	doc, err := html.Parse(file)
	if err != nil {
		log.Println(err)
		return
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		// 遍历节点，根据节点类型和名称判断
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