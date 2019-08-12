package douban

import (
	"GopherBook/chapter7/assistance"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/tidwall/gjson"

	"github.com/chromedp/chromedp"
)

func GetDouBan(url string) {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Fatalf))
	defer cancel()
	//ctxTimeOut, _ := context.WithTimeout(ctx, time.Second*10)
	var response string
	chromedp.Run(ctx, Tasks(url, &response))
	//fmt.Println(response)
	results := ParseDouBan(response)
	for _, i := range results {
		fmt.Println(i)

	}
}

func Tasks(url string, response *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.OuterHTML("#content", response),
	}
}

func ParseDouBan(response string) []ResultForDouBan {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Println(err)
		return nil
	}
	var results []ResultForDouBan
	doc.Find(".list-wp a").Each(func(i int, selection *goquery.Selection) {
		var r ResultForDouBan
		href, exists := selection.Attr("href")
		if exists {
			r.Url = href

		}
		r.Title = selection.Find(`p span[class="title"]`).Text()
		r.Rate = selection.Find(`p span[class="rate"]`).Text()
		results = append(results, r)
	})
	return results

}

// GetDouBanByAPI...
func GetDouBanByAPI(url string, start int) {
	fullURL := fmt.Sprintf(url, start)
	fmt.Println(fullURL)
	content, err := assistance.GetContent(fullURL)
	if err != nil {
		log.Println(err)
		return
	}
	doc := gjson.ParseBytes(content).Get("data").Array()
	for _, i := range doc {
		var r ResultForDouBan

		list := func(i gjson.Result) []string {
			var c []string
			for _, j := range i.Array() {
				c = append(c, j.String())
			}
			return c
		}
		r = ResultForDouBan{
			Casts:     list(i.Get("casts")),
			Directors: list(i.Get("directors")),
			Rate:      i.Get("rate").String(),
			Star:      i.Get("star").Str,
			Url:       i.Get("url").String(),
			Title:     i.Get("title").String(),
		}
		fmt.Println(r)
	}
}
