package weibo

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter7/assistance"
	"log"
	"testing"
)

func TestParseWeiBo(t *testing.T) {
	content, err := assistance.GetContent(WeiBoRoot)
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(string(content))
	ParseWeiBo(content)
}
