package data

import (
	"GopherBook/chapter12/fina/configs"
	"GopherBook/chapter12/fina/pkg/assistance"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/tidwall/gjson"

	"github.com/PuerkitoBio/goquery"
)

func TestParseSportsByQuery(t *testing.T) {
	f, err := os.Open("sports.html")
	if err != nil {
		log.Println(err)
		return
	}

	ParseSportsByQuery(f)
}

func TestParseSportsByQuery2(t *testing.T) {
	for key, i := range configs.MatchSportsMap {
		url := fmt.Sprintf(configs.MatchSports, key, i)
		fmt.Println(url)
		reader, _ := assistance.DownloadByChromeHeadless(url)
		doc, _ := goquery.NewDocumentFromReader(reader)
		fmt.Println(doc.Html())
	}

}

func TestParseSportsByQuery3(t *testing.T) {
	for _, i := range configs.MatchSportsMap {
		payload := strings.NewReader(fmt.Sprintf("sn=%d", i))
		content, err := assistance.PostReturnIOReader(configs.MatchPostDo, payload)
		if err != nil {
			log.Println(err)
			return
		}
		data := gjson.ParseBytes(content)
		fmt.Println(data.Get("dataArr"))

	}
}
