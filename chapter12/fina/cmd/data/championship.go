package data

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/assistance"
	"GopherBook/chapter12/fina/pkg/database"
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func RunChampionship(url string) (bool, error) {
	reader, err := assistance.DownloaderReturnIOReader(url)
	if err != nil {
		return false, err
	}
	result, err := ParseChampionshipByJquery(reader)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	tx := database.MySQL.NewSession()
	tx.Begin()
	disciplines := result.Disciplines
	for _, i := range disciplines {
		if _, dbError := tx.InsertOne(&i); dbError != nil {
			tx.Rollback()
			return false, dbError
		}
		result.Championships.DisciplinesIds = append(result.Championships.DisciplinesIds, i.Id)
	}

	venues := result.Venues
	for _, i := range venues {
		if _, dbError := tx.InsertOne(&i); dbError != nil {
			tx.Rollback()
			return false, dbError
		}
		result.Championships.VenuesIds = append(result.Championships.VenuesIds, i.Id)
	}

	if _, dbError := tx.InsertOne(&result.Championships); dbError != nil {
		tx.Rollback()
		return false, dbError
	}
	tx.Commit()
	return true, nil
}

type ResultForChampionship struct {
	Championships models.FiFaChampionships
	Disciplines   []models.Kinds
	Venues        []models.Kinds
}

func ParseChampionshipByJquery(reader io.Reader) (ResultForChampionship, error) {
	var result ResultForChampionship
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return result, err
	}
	doc.Find(".content.content_wide div").Each(func(i int, selection *goquery.Selection) {
		if i >= 4 {
			return
		}
		if i == 0 {
			if val, ok := selection.Find("img").Attr("alt"); ok {
				result.Championships.ShortSlogan = val
				result.Championships.NumberOlympic = 18
			}
			return
		}
		if i == 1 {
			start, end, err := assistance.GetDate(selection.Find("p").Text())
			if err == nil {
				result.Championships.StartDate = start
				result.Championships.EndDate = end
			} else {
				return
			}
			return

		}
		var class int
		if i == 2 {
			class = models.DISCIPLINE
		} else {
			class = models.VENUES
		}
		var kinds []models.Kinds
		results := assistance.GetDisciplines(selection.Find("p").Text())
		for _, i := range results {
			var k models.Kinds
			k = models.Kinds{
				Name:  strings.TrimSpace(i),
				Class: class,
			}
			kinds = append(kinds, k)
		}
		if i == 2 {
			result.Disciplines = kinds
		} else {
			result.Venues = kinds
		}

	})
	return result, nil
}
