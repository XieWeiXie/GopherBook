package data

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/assistance"
	"GopherBook/chapter12/fina/pkg/database"
	"fmt"
	"io"
	"os"

	"github.com/antchfx/htmlquery"
)

func RunFiFaHistory(url string) (bool, error) {
	reader, err := assistance.DownloaderReturnIOReader(url)
	if err != nil {
		fmt.Println(err)
		reader, err = os.Open("fifa_history.html")
		if err != nil {
			fmt.Println(err)
			return false, err
		}
	}
	result, err := ParseFiFaHistoryByXpath(reader)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	tx := database.MySQL.NewSession()
	tx.Begin()
	for _, i := range result {
		if _, dbError := tx.InsertOne(&i); dbError != nil {
			tx.Rollback()
			return false, dbError
		}
	}
	tx.Commit()

	fmt.Println(result)
	return true, nil

}

func ParseFiFaHistoryByXpath(reader io.Reader) ([]models.FiFaHistory, error) {

	doc, error := htmlquery.Parse(reader)
	if error != nil {
		return nil, error
	}
	var result []models.FiFaHistory
	list := htmlquery.Find(doc, `//ul[@class="history mt25"]/li`)
	for _, i := range list {
		var one models.FiFaHistory
		one = models.FiFaHistory{
			Year:   assistance.ToInt(htmlquery.InnerText(htmlquery.FindOne(i, "/strong"))),
			Detail: assistance.ReplaceSpace(htmlquery.InnerText(htmlquery.FindOne(i, "/p"))),
		}
		fmt.Println(one)
		result = append(result, one)
	}
	result = append(result, models.FiFaHistory{
		Year:   2019,
		Detail: "从7月12日到28日，在光州(韩国)举办了FINA世界游泳锦标赛",
	})
	return result, nil
}
