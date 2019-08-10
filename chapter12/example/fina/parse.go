package fina

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/tidwall/gjson"

	"github.com/PuerkitoBio/goquery"

	"github.com/antchfx/htmlquery"
)

func ParseByXpath(reader io.Reader) {
	doc, err := htmlquery.Parse(reader)
	if err != nil {
		log.Println(err)
		return
	}
	text := htmlquery.FindOne(doc, "//div/div/ul")
	fmt.Println(htmlquery.InnerText(text))
}

func ParseByCss(reader io.Reader) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(doc.Find("div div ul").Text())
}

func ParseByJson(reader io.Reader) {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println(err)
		return
	}

	type result struct {
		sport       string
		event       string
		gender      string
		participant string
		national    string
		date        string
	}

	doc := gjson.ParseBytes(content).Array()

	var one result
	for _, i := range doc {
		sport := i.Get("c_Sport").String()
		event := i.Get("c_Event").String()
		gender := i.Get("c_Gender").String()
		participant := i.Get("c_Participant").String()
		national := i.Get("c_ParticipantNatio").String()
		date := i.Get("c_Date").String()
		one = result{
			sport:       sport,
			event:       event,
			gender:      gender,
			participant: participant,
			national:    national,
			date:        date,
		}
		fmt.Println(one)
	}

}
