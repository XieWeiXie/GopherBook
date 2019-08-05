package blue

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/database"
)

type ControllerBlue struct {
}

var Default = ControllerBlue{}

func (C ControllerBlue) GetBlues() ([]models.BlueSerializer, error) {
	var result []models.BlueSerializer
	var blues []models.Blue
	if dbError := database.MySQL.Find(&blues); dbError != nil {
		return result, nil
	}
	for _, i := range blues {
		result = append(result, i.Serializer())
	}
	return result, nil

}
