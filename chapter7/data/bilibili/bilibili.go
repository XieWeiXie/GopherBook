package bilibili

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter7/assistance"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetBiliBili(url string) {
	content, err := assistance.GetContent(url)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	if err != nil {
		panic(err)
	}
	doc.Find(`ul[class="rank-list"] li[class="rank-item"] div[class="content"] div[class="info"]`).Each(func(i int, selection *goquery.Selection) {
		var r ResultForBiliBili
		item := selection.Find(`a[class="title"]`)
		text := item.Text()
		r.Title = text
		href, exists := item.Attr("href")
		if exists {
			r.Href = fmt.Sprintf("https:" + href)
		}
		author := selection.Find(`div[class="detail"] a span`)
		r.Author = author.Text()
		authorLink := selection.Find(`div[class="detail"] a`)
		authorHref, exists := authorLink.Attr("href")
		if exists {
			r.AuthorURL = fmt.Sprintf("https:" + authorHref)
		}
		detail := selection.Find(`div[class="detail"] span`)
		play := detail.Eq(0).Text()
		r.Play = play
		view := detail.Eq(1).Text()
		r.View = view
		pts := selection.Find(`div[class="pts"] div`)
		r.Pts = pts.Text()
		fmt.Println(r)
	})
}
