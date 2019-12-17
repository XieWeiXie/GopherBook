package country_medal

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/database"
)

type ControllerCountryMedal struct {
}

var Default = ControllerCountryMedal{}

func (C ControllerCountryMedal) GetCountryMedal(param GetCountryMedalParam) (models.CountryMedalSerializer, error) {
	var result models.CountryMedalSerializer
	if err := param.Valid(); err != nil {
		return result, err
	}
	var country models.Country
	if has, dbError := database.MySQL.Where("name = ?", param.Name).Get(&country); dbError != nil || !has {
		return result, dbError
	}
	var countryMedal models.CountryMedal
	if has, dbError := database.MySQL.Where("country_id = ? AND year = ?", country.Id, param.Year).Get(&countryMedal); !has || dbError != nil {
		return result, dbError
	}
	result = countryMedal.Serializer()
	return result, nil
}

func (C ControllerCountryMedal) Rank(param RankCountryMedalParam) ([]models.CountryMedalSerializer, error) {
	var result []models.CountryMedalSerializer
	if err := param.Valid(); err != nil {
		return result, err
	}
	var countryMedals []models.CountryMedal
	if param.SortBy == "total" {
		SQL := fmt.Sprintf("SELECT sum(gold+silver+bronze) as sum,id, gold, silver, bronze, country_id FROM medal WHERE YEAR = %d GROUP BY id ORDER BY sum desc", param.Year)
		if dbError := database.MySQL.SQL(SQL).Find(&countryMedals); dbError != nil {
			return nil, dbError
		}
	} else {
		if dbError := database.MySQL.Where("year = ?", param.Year).OrderBy(param.SortBy + " desc").Find(&countryMedals); dbError != nil {
			return nil, dbError
		}
	}
	for _, i := range countryMedals {
		result = append(result, i.Serializer())
	}
	return result, nil
}
