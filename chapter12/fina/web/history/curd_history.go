package history

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/database"
)

type ControllerHistory struct {
}

var Default = ControllerHistory{}

func (C ControllerHistory) GetHistory(param GetHistoryParam) (models.FiNaHistorySerializer, error) {
	var result models.FiNaHistorySerializer
	if err := param.Valid(); err != nil {
		return result, nil
	}
	var history models.FiNaHistory
	if has, dbError := database.MySQL.Where("year = ?", param.Year).Get(&history); !has || dbError != nil {
		return result, dbError
	}
	result = history.Serializer()
	return result, nil
}

func (C ControllerHistory) GetAll(param GetAllHistoryParam) ([]models.FiNaHistorySerializer, error) {
	var result []models.FiNaHistorySerializer
	var histories []models.FiNaHistory
	if dbError := database.MySQL.OrderBy(param.OrderBy + " desc").Find(&histories); dbError != nil {
		return nil, dbError
	}
	for _, i := range histories {
		result = append(result, i.Serializer())
	}
	return result, nil
}
