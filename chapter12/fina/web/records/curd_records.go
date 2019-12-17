package records

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/models"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/database"
)

type ControllerRecord struct {
}

var Default = ControllerRecord{}

func (C ControllerRecord) GetRecords(param GetRecordParam) ([]models.RecordsMaxSerializer, error) {
	var result []models.RecordsMaxSerializer
	var sports []models.RecordMax
	if err := param.Valid(); err != nil {
		return result, err
	}
	query := database.MySQL.NewSession()
	fmt.Println("data")
	if param.All {
		if dbError := query.Find(&sports); dbError != nil {
			return result, dbError
		}
	} else {
		if param.Name != "" {
			query = query.Where("name like ?", fmt.Sprintf("%s%%", param.Name))
		}
		query = query.Where("sport_class = ? AND competition_class = ?", param.SportClass, param.CompetitionClass)
		if dbError := query.Find(&sports); dbError != nil {
			return result, dbError
		}
	}
	for _, i := range sports {
		result = append(result, i.Serializer())
	}
	return result, nil
}
