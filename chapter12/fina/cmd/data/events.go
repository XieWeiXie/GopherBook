package data

import (
	"GopherBook/chapter12/fina/configs"
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/assistance"
	"GopherBook/chapter12/fina/pkg/database"
	"fmt"

	"github.com/tidwall/gjson"
)

func RunPostEvent(url string) (bool, error) {
	content, err := assistance.Downloader(url)
	if err != nil {
		return false, err
	}
	for key, v := range urlList(content) {
		fmt.Println(key, v)
		c, err := assistance.Downloader(v)
		if err != nil {
			return false, err
		}
		if err := ParsePostEvent(c, int(key)); err != nil {
			return false, err
		}
	}
	return true, nil
}

func urlList(content []byte) map[int64]string {
	var result = make(map[int64]string)
	list := gjson.ParseBytes(content).Array()
	for _, i := range list {
		year := i.Get("n_Season").Int()
		url := fmt.Sprintf(configs.MatchPostEvent, year)
		result[year] = url
	}
	return result
}
func ParsePostEvent(content []byte, year int) error {
	tx := database.MySQL.NewSession()
	tx.Begin()
	medals := gjson.ParseBytes(content).Get("MedalTableNOC").Array()
	for _, i := range medals {
		var country models.Country
		var countryMedal models.CountryMedal
		name := i.Get("c_NOC").String()
		short := i.Get("c_NOCShort").String()
		if has, dbError := tx.Where("name = ? AND short = ?", name, short).Get(&country); dbError != nil || !has {
			country = models.Country{
				Name:  name,
				Short: short,
			}
			if _, dbError := tx.InsertOne(&country); dbError != nil {
				tx.Rollback()
				return dbError
			}
		}
		gold := i.Get("n_Gold").Int()
		silver := i.Get("n_Silver").Int()
		bronze := i.Get("n_Bronze").Int()
		countryMedal = models.CountryMedal{
			Year:      year,
			CountryId: country.Id,
			Gold:      int(gold),
			Silver:    int(silver),
			Bronze:    int(bronze),
		}
		if _, dbError := tx.InsertOne(&countryMedal); dbError != nil {
			tx.Rollback()
			return dbError
		}

	}
	tx.Commit()
	return nil
}
