package data

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/assistance"
	"GopherBook/chapter12/fina/pkg/database"
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func RunFiNaBrief(url string) (bool, error) {
	reader, err := assistance.DownloaderReturnIOReader(url)
	if err != nil {
		log.Println(err)
		return false, err
	}
	result, err := ParseFiNaBriefByjQuery(reader)
	if err != nil {
		log.Println(err)
		return false, err
	}
	tx := database.MySQL.NewSession()
	tx.Begin()
	if _, dbError := tx.InsertOne(&result); dbError != nil {
		tx.Rollback()
		return false, dbError
	}
	tx.Commit()
	return true, nil
}

func ParseFiNaBriefByjQuery(reader io.Reader) (models.FiNa, error) {
	var result models.FiNa
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println(err)
		return result, err
	}
	selection := doc.Find(".content_wrap div.content")
	result.Description = assistance.ReplaceSpace(selection.Find("p").Eq(0).Text())
	var detail bytes.Buffer
	selection.Find("p").Each(func(i int, selection *goquery.Selection) {
		if i != 0 {
			detail.WriteString(assistance.ReplaceSpace(selection.Text()))
		}
	})
	ul := selection.Find("ul.list_basic.lt01.mt30 li")
	if ul.Size() != 4 {
		return result, fmt.Errorf("error")
	}
	result.Established = assistance.SplitBYColon(ul.Eq(0).Text(), ":")
	result.Headquarters = assistance.SplitBYColon(ul.Eq(1).Text(), ":")
	result.NationalMember = assistance.SplitBYColon(ul.Eq(2).Text(), ":")
	result.NumberOfDisciplines = assistance.SplitBYColon(ul.Eq(3).Text(), ":")
	return result, nil
}
