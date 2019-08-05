package country_medal

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/database"
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
