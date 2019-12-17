package data

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/configs"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/assistance"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/database"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"

	"github.com/PuerkitoBio/goquery"
)

func RunSports(m map[int]int) (bool, error) {
	tx := database.MySQL.NewSession()
	tx.Begin()
	for _, i := range m {
		payload := strings.NewReader(fmt.Sprintf("sn=%d", i))
		content, err := assistance.PostReturnIOReader(configs.MatchPostDo, payload)
		if err != nil {
			log.Println(err, "sn=", i)
			return false, err
		}
		result, err := ParseSportsByGjson(content, i)
		if err != nil {
			log.Println(err)
			return false, err
		}

		for _, i := range result.Competitions {
			if _, dbError := tx.InsertOne(&i); dbError != nil {
				tx.Rollback()
				return false, dbError
			}
			result.Sports.CompetitionIds = append(result.Sports.CompetitionIds, i.Id)
		}
		if _, dbError := tx.InsertOne(&result.Sports); dbError != nil {
			tx.Rollback()
			return false, dbError
		}

	}
	tx.Commit()
	return true, nil
}

type ResultForSports struct {
	Sports       models.Sports
	Competitions []models.Competitions
}

// Deprecated
func ParseSportsByQuery(reader io.Reader) (ResultForSports, error) {
	var result ResultForSports
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println(err)
		return result, err
	}
	result.Sports.Description = assistance.ReplaceSpace(doc.Find("#intro_S p").Text())
	doc.Find(".content div.games_wrap ul li").Each(func(i int, selection *goquery.Selection) {
		if val, ok := selection.Attr("class"); ok && val == "active" {
			result.Sports.SportName = selection.Find("span").Text()
			result.Sports.SportClass = i
		}
	})
	total := strings.TrimSpace(assistance.SplitBYColon(doc.Find("#games_S span").Text(), ":"))
	result.Sports.Total, _ = strconv.Atoi(total)
	doc.Find("#games_S div div").Each(func(i int, selection *goquery.Selection) {
		class := selection.Find("span").Text()
		classInt := getCompetitionClass(class)
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
	} else {
		return models.TEAM
	}
}

func ParseSportsByGjson(content []byte, sn int) (ResultForSports, error) {
	data := gjson.ParseBytes(content)
	var result ResultForSports
	dataArr := data.Get("dataArr")
	result.Sports.Total = int(dataArr.Get("MEDAL_TOTAL").Int())
	result.Sports.Description = dataArr.Get("INTRO_CHN").String()
	result.Sports.Rule = dataArr.Get("METHOD_CHN").String()
	if dataArr.Get("EVENT_GIRL_CHN").Exists() {
		woman := assistance.SplitBySep(dataArr.Get("EVENT_GIRL_CHN").String(), "|")
		for _, i := range woman {
			var one models.Competitions
			one = models.Competitions{
				CompetitionClass: models.WOMAN,
				Detail:           i,
			}
			result.Competitions = append(result.Competitions, one)
		}
	}
	if dataArr.Get("EVENT_MAN_CHN").Exists() {
		man := assistance.SplitBySep(dataArr.Get("EVENT_MAN_CHN").String(), "|")
		for _, i := range man {
			var one models.Competitions
			one = models.Competitions{
				CompetitionClass: models.MAN,
				Detail:           strings.TrimSpace(i),
			}
			result.Competitions = append(result.Competitions, one)
		}
	}
	if dataArr.Get("COED_TEAM_CHN").Exists() {
		team := assistance.SplitBySep(dataArr.Get("COED_TEAM_CHN").String(), "|")
		for _, i := range team {
			var one models.Competitions
			one = models.Competitions{
				CompetitionClass: models.TEAM,
				Detail:           i,
			}
			result.Competitions = append(result.Competitions, one)

		}
	}
	result.Sports.SportClass = 6 - sn
	result.Sports.SportName = models.SportClass[6-sn]
	return result, nil
}
