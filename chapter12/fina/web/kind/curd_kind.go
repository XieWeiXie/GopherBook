package kind

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/database"
)

type ControllerKind struct {
}

var Default = ControllerKind{}

func (C ControllerKind) GetKinds(param GetKindParam) ([]models.KindSerializer, error) {
	var result []models.KindSerializer
	if err := param.Valid(); err != nil {
		return result, err
	}
	var kinds []models.Kinds
	if dbError := database.MySQL.Where("class = ?", param.Class).Find(&kinds); dbError != nil {
		return result, dbError
	}
	for _, i := range kinds {
		result = append(result, i.Serializer())
	}
	return result, nil
}
