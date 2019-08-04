package data

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/assistance"
	"GopherBook/chapter12/fina/pkg/database"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func RunRecords(url string) (bool, error) {
	content, err := assistance.Downloader(url)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return ParseRecordsJson(content)
}

func ParseRecordsJson(content []byte) (bool, error) {
	data := gjson.ParseBytes(content).Array()
	if len(data) <= 0 {
		return false, fmt.Errorf("is not array")
	}
	tx := database.MySQL.NewSession()
	tx.Begin()
	for _, i := range data {
		var country models.Country
		countryName := i.Get("c_NOC").String()
		if has, dbError := tx.Where("name = ?", countryName).Get(&country); !has || dbError != nil {
			country = models.Country{
				Name:  countryName,
				Short: i.Get("c_NOCShort").String(),
			}
			if _, dbError := tx.InsertOne(&country); dbError != nil {
				tx.Rollback()
				return false, dbError
			}
		}
		var records models.RecordMax
		records = models.RecordMax{
			CountryId:        country.Id,
			Date:             getDate(i.Get("c_Date").String()),
			EventName:        i.Get("c_Event").String(),
			CompetitionClass: getClass(i.Get("c_Gender").String()),
			Location:         i.Get("c_Location").String(),
			Record:           i.Get("c_Result").String(),
			SportClass:       getSportClass(i.Get("c_Sport").String()),
			Name:             i.Get("c_Participant").String(),
		}
		if _, dbError := tx.InsertOne(&records); dbError != nil {
			tx.Rollback()
			return false, dbError
		}

	}
	tx.Commit()
	return true, nil
}

var getClass = func(value string) int {
	v := strings.ToUpper(strings.TrimSpace(value))
	return models.CompetitionClassEn[v]
}

var getSportClass = func(value string) int {
	v := strings.TrimSpace(value)
	return models.SportClassEn[strings.ToUpper(v)]

}

var getDate = func(value string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02T15:04:05", value, time.Local)
	return t

}
