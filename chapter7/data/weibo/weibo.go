package weibo

import (
	"GopherBook/chapter7/assistance"
	"fmt"
	"log"
	"strings"

	"github.com/antchfx/htmlquery"
)

func ParseWeiBo(content []byte) {
	reader := strings.NewReader(string(content))
	doc, err := htmlquery.Parse(reader)
	if err != nil {
		log.Println(err)
		return
	}
	tds := htmlquery.Find(doc, `//div[@id="pl_top_realtimehot"]//tbody/tr/td[2]`)
	for index, i := range tds {
		if index == 0 {
			continue
		}
		a := htmlquery.FindOne(i, "/a")
		if len(a.Attr) > 2 {
			continue
		}
		aText := htmlquery.InnerText(a)
		aHref := htmlquery.InnerText(htmlquery.FindOne(a, "/@href"))
		var result ResultForWeiBo
		result = ResultForWeiBo{
			Title: strings.TrimSpace(aText),
			Url:   fmt.Sprintf("%s%s", HOST, strings.TrimSpace(aHref)),
		}
		span := htmlquery.FindOne(i, "/span")
		if span != nil {
			result.Score = assistance.ToInt(htmlquery.InnerText(span))
		}
		fmt.Println(result)
	}
}
