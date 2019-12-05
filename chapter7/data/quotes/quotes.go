package quotes

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter7/assistance"
	"log"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func GetQuotesContentAll() {
	for i := 1; i < 11; i++ {
		GetQuotesContent(fmt.Sprintf("http://quotes.toscrape.com/js/page/%d/", i))
	}
}
func GetQuotesContent(url string) []ResultForQuotes {
	content := assistance.ChromedpGetContent(url)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Println(err)
		return nil
	}
	var results []ResultForQuotes
	doc.Find(`body > div[class="container"] > div[class="quote"]`).
		Each(func(i int, selection *goquery.Selection) {
			var one ResultForQuotes
			tags := func() []string {
				var ts []string
				selection.Find("div > a").
					Each(func(i int, selection *goquery.Selection) {
						ts = append(ts, selection.Text())
					})
				return ts
			}
			one = ResultForQuotes{
				Text:   selection.Find("span").Eq(0).Text(),
				Author: selection.Find("span > small").Text(),
				Tags:   tags(),
			}
			results = append(results, one)
			fmt.Println(one)
		})
	attr, ok := doc.Find(`li[class="next"] > a`).Attr("href")
	if !ok {
		return results
	} else {
		url := fmt.Sprintf("http://quotes.toscrape.com" + attr)
		GetQuotesContent(url)
		return results
	}
}

func GetQuotesContentByClick(url string) []ResultForQuotes {
	content, next := ClickNext(url)
	var results []ResultForQuotes
	getResults := func(content string) []ResultForQuotes {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
		if err != nil {
			log.Println(err)
			return nil
		}
		doc.Find(`body > div[class="container"] > div[class="quote"]`).
			Each(func(i int, selection *goquery.Selection) {
				var one ResultForQuotes
				tags := func() []string {
					var ts []string
					selection.Find("div > a").
						Each(func(i int, selection *goquery.Selection) {
							ts = append(ts, selection.Text())
						})
					return ts
				}
				one = ResultForQuotes{
					Text:   selection.Find("span").Eq(0).Text(),
					Author: selection.Find("span > small").Text(),
					Tags:   tags(),
				}
				results = append(results, one)
				fmt.Println(one)
			})
		return results
	}
	results = append(results, getResults(content)...)
	fmt.Println(next)
	if strings.Contains(next, "http://") {
		GetQuotesContentByClick(next)
	} else {
		return results
	}
	return results
}

func ClickNext(url string) (string, string) {
	var nextPage map[string]string
	var pageSource string
	ctx1, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()
	ctx, cancel1 := context.WithTimeout(ctx1, 30*time.Second)
	defer cancel1()
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(".footer", chromedp.ByQuery),
		chromedp.OuterHTML("body", &pageSource),
		chromedp.Attributes(`li[class="next"] > a`, &nextPage),
		chromedp.Click(`li[class="next"] > a`, chromedp.ByQuery),
	})
	if err != nil {
		log.Println(err)
		err := chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.WaitVisible(".footer", chromedp.ByQuery),
			chromedp.OuterHTML("body", &pageSource),
			chromedp.WaitNotPresent(`li[class="next"]`, chromedp.ByQuery),
		})
		if err != nil {
			return pageSource, ""
		}
	}
	if _, ok := nextPage["href"]; ok {
		return pageSource, fmt.Sprintf("http://quotes.toscrape.com" + nextPage["href"])
	}
	return pageSource, ""
}
