package competition

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/database"
)

type ControllerCompetition struct {
}

var Default = ControllerCompetition{}

func (C ControllerCompetition) GetCompetitions(param GetCompetitionParam) ([]models.CompetitionSerializer, error) {
	var result []models.CompetitionSerializer
	if err := param.Valid(); err != nil {
		return result, err
	}
	var competitions []models.Competitions
	if dbError := database.MySQL.Where("competition_class = ?", param.Class).Find(&competitions); dbError != nil {
		return result, dbError
	}
	for _, i := range competitions {
		result = append(result, i.Serializer())
	}
	return result, nil
}
