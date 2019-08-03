package data

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"testing"

	"golang.org/x/net/html/charset"

	"github.com/antchfx/htmlquery"

	"golang.org/x/net/html"
)

func TestHtmlNode(tests *testing.T) {
	content, _ := os.Open("symbol.html")
	doc, _ := html.Parse(content)
	var f func(node *html.Node, buf *bytes.Buffer)
	f = func(n *html.Node, buf *bytes.Buffer) {
		switch n.Type {
		case html.TextNode:
			buf.WriteString(n.Data)
			return
		case html.CommentNode:
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, buf)
		}
	}
	var buf bytes.Buffer
	f(doc, &buf)
	fmt.Println(buf.String())
}

func TestHtmlXpath(tests *testing.T) {

	response, err := http.Get("https://www.fina-gwangju2019.com/chn/contentsView.do?pageId=chn4")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader, _ := charset.NewReader(response.Body, response.Header.Get("Content-type"))
	doc, err := htmlquery.Parse(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	list := htmlquery.Find(doc, `//div[@class="content content_wide"]/div`)
	for index, i := range list {
		fmt.Println(index, htmlquery.InnerText(i))
	}
}

func TestParseByXpath(t *testing.T) {
	response, err := http.Get("https://www.fina-gwangju2019.com/chn/contentsView.do?pageId=chn4")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader, _ := charset.NewReader(response.Body, response.Header.Get("Content-type"))

	fmt.Println(ParseSymbolByXpath(reader))
}
