package country

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/database"
	"fmt"
)

type ControllerCountry struct{}

var Default = ControllerCountry{}

func (C ControllerCountry) GetList(param GetCountryParam) ([]models.CountrySerializer, error) {
	var countries []models.Country
	if err := param.Valid(); err != nil {
		return nil, err
	}
	query := database.MySQL.Where("name like ?", fmt.Sprintf("%s%%", param.Name))
	if param.Short != "" {
		query.Where("short like ?", fmt.Sprintf("%s%%", param.Short))
	}
	if dbError := query.Find(&countries); dbError != nil {
		return nil, dbError
	}
	var result []models.CountrySerializer
	for _, i := range countries {
		result = append(result, i.Serializer())
	}
	return result, nil

}

func (C ControllerCountry) AllList(param GetCountryParam) ([]models.CountrySerializer, error) {
	var result []models.CountrySerializer
	var countries []models.Country
	if dbError := database.MySQL.Find(&countries); dbError != nil {
		return result, dbError
	}
	for _, i := range countries {
		result = append(result, i.Serializer())
	}
	return result, nil
}
