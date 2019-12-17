package sports

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/database"
)

type ControllerSports struct {
}

var Default = ControllerSports{}

func (C ControllerSports) GetSports(param GetSportParam) ([]models.SportSerializer, error) {
	var result []models.SportSerializer
	var sports []models.Sports
	if err := param.Valid(); err != nil {
		return result, err
	}
	if dbError := database.MySQL.Where("sport_class = ?", param.Class).Find(&sports); dbError != nil {
		return result, nil
	}
	for _, i := range sports {
		result = append(result, i.Serializer())
	}
	return result, nil
}
