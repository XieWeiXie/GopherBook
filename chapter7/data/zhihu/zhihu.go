package zhihu

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"

	"github.com/chromedp/chromedp"
)

var res string
var jsonS string

func GetZhiHu(url string) {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Fatalf))
	defer cancel()
	var response string
	err := chromedp.Run(ctx, Tasks(&response))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res)
	fmt.Println(response)
	Parse(jsonS)

}

func Tasks(response *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.zhihu.com/billboard`),
		chromedp.WaitVisible(`main`, chromedp.ByQuery),
		chromedp.Text(`main`, &res, chromedp.NodeVisible),
		chromedp.OuterHTML(`js-initialData`, response),
	}
}

func Parse(content string) []ResultForZhiHu {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Println(err)
		return nil
	}
	docJSON := gjson.Parse(doc.Find("#js-initialData").Text())
	array := docJSON.Get("initialState.topstory.hotList").Array()

	for _, i := range array {
		var r ResultForZhiHu
		r.Metrics = i.Get("target.metricsArea.text").String()
		r.TitleArea = i.Get("target.titleArea.text").String()
		r.ExcerptArea = i.Get("target.excerptArea.text").String()
		r.Link = i.Get("target.link.url").String()
		fmt.Println(r)

	}
	return nil

}
