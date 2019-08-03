package data

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/assistance"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func RunSports() {}

type ResultForSports struct {
	Sports       models.Sports
	Competitions []models.Competitions
}

func ParseSportsByQuery(reader io.Reader) (ResultForSports, error) {
	var result ResultForSports
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println(err)
		return result, err
	}
	fmt.Println(doc.Html())
	result.Sports.Description = assistance.ReplaceSpace(doc.Find("#intro_S p").Text())
	doc.Find(".content div.games_wrap ul li").Each(func(i int, selection *goquery.Selection) {
		if val, ok := selection.Attr("class"); ok && val == "active" {
			result.Sports.SportName = selection.Find("span").Text()
			result.Sports.SportClass = i
		}
	})
	fmt.Println(doc.Find("#games_S span").Text())
	total := strings.TrimSpace(assistance.SplitBYColon(doc.Find("#games_S span").Text(), ":"))
	result.Sports.Total, _ = strconv.Atoi(total)
	doc.Find("#games_S div div").Each(func(i int, selection *goquery.Selection) {
		class := selection.Find("span").Text()
		classInt := getCompetitionClass(class)
		fmt.Println(classInt, class)
		selection.Find("ul li").Each(func(i int, selection *goquery.Selection) {
			var one models.Competitions
			one.CompetitionClass = classInt
			text := selection.Text()
			if strings.Contains(selection.Text(), ",") {
				text = strings.Replace(text, ",", "、", -1)
			}
			one.Detail = assistance.ReplaceSpace(text)
			fmt.Println(one)
			result.Competitions = append(result.Competitions, one)
		})
	})
	return result, nil
}

var getCompetitionClass = func(value string) int {
	if strings.HasPrefix(value, "男子") {
		return models.MAN
	} else if strings.HasSuffix(value, "女子") {
		return models.WOMAN
	} else if strings.HasSuffix(value, "男女混合") {
		return models.MIX
	} else {
		return models.TEAM
	}
}
