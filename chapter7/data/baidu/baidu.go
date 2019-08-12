package baidu

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/chromedp/chromedp"
)

func GetBaiDu(url string) {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Fatalf))
	defer cancel()
	var response string
	err := chromedp.Run(ctx, Tasks(url, &response))
	if err != nil {
		log.Println(err)
		return
	}
	urls := Parse(response)
	for _, i := range urls {
		fmt.Println(i)
		var childResponse string
		chromedp.Run(ctx, AnotherTasks(i, &childResponse))
		AnotherParse(childResponse)
	}

}

func Tasks(url string, response *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("#flist", chromedp.ByQuery),
		chromedp.OuterHTML("body", response),
	}
}

func Parse(response string) []string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Println(err)
		return nil
	}
	//fmt.Println(doc.Html())
	var urls []string
	doc.Find("#flist div ul li").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		if v, ok := selection.Find("a").Attr("href"); ok {
			urls = append(urls, strings.Replace(v, ".", ROOT, 1))
		}
	})
	return urls

}

func AnotherTasks(url string, response *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("tbody", chromedp.ByQuery),
		chromedp.OuterHTML("body", response),
	}
}

func AnotherParse(response string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(doc.Html())
	doc.Find("tbody tr").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		if v, ok := selection.Attr("class"); ok {
			if v == "item-tr" {
				return
			}
		}
		var r ResultBaiDu
		keyword := selection.Find(`td[class="keyword"] a`).Eq(0)
		r.Keyword = strings.TrimSpace(keyword.Text())
		if v, ok := keyword.Attr("href"); ok {
			r.Href = v
		}
		r.Number, _ = strconv.Atoi(selection.Find(`td[class="last"] span`).Text())
		fmt.Println(r)

	})
}
