package data

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/assistance"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/database"
	"log"
	"strings"

	"github.com/tidwall/gjson"
)

func RunRank(url string) (bool, error) {
	content, err := assistance.PostReturnIOReader(url, nil)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return ParseRankByGjson(content)
}

func ParseRankByGjson(content []byte) (bool, error) {
	data := gjson.ParseBytes(content).Get("dataArr").Array()
	if len(data) == 0 {
		return false, fmt.Errorf("not array")
	}
	tx := database.MySQL.NewSession()
	tx.Begin()
	year := 2019
	for _, i := range data {
		nation := i.Get("NATION_NM").String()
		short := strings.TrimSpace(strings.Split(nation, "-")[0])
		var country models.Country
		if has, dbError := tx.Where("short = ?", short).Get(&country); dbError != nil || !has {
			name := strings.TrimSpace(strings.Split(nation, "-")[1])
			country.Short = short
			country.Name = name
			if _, dbError := tx.InsertOne(&country); dbError != nil {
				tx.Rollback()
				return false, dbError
			}
		}
		var countryMedal models.CountryMedal
		countryMedal = models.CountryMedal{
			Year:      year,
			CountryId: country.Id,
			Gold:      int(i.Get("GOLD_MEDAL").Int()),
			Silver:    int(i.Get("SILVER_MEDAL").Int()),
			Bronze:    int(i.Get("BRONZE_MEDAL").Int()),
		}
		if _, dbError := tx.InsertOne(&countryMedal); dbError != nil {
			tx.Rollback()
			return false, dbError
		}

	}
	tx.Commit()
	return true, nil
}
