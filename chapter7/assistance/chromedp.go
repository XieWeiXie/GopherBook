package assistance

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func ChromedpGetContent(url string) string {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()
	var response string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.OuterHTML("body", &response),
	})
	if err != nil {
		log.Println(err)
		return ""
	}
	return response
}
