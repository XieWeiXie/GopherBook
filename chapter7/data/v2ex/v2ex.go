package v2ex

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter7/assistance"
	"strings"

	"github.com/tidwall/gjson"
)

func V2ex(url string) {
	content, err := assistance.GetContent(url)
	if err != nil {
		panic(err)
	}
	doc := gjson.ParseBytes(content).Array()
	for _, i := range doc {
		var r ResultForV2ex
		r = ResultForV2ex{
			Title:       i.Get("title").String(),
			URL:         i.Get("url").String(),
			Description: strings.TrimSpace(i.Get("content").String()),
			Content:     strings.TrimSpace(i.Get("content_rendered").String()),
		}
		fmt.Println(r)
	}
}
