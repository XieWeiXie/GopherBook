package data

import (
	"GopherBook/chapter12/fina/models"
	"GopherBook/chapter12/fina/pkg/assistance"
	"GopherBook/chapter12/fina/pkg/database"
	"io"
	"strings"

	"github.com/antchfx/htmlquery"
)

func RunForSymbol(url string) (bool, error) {
	content, err := assistance.Downloader(url)
	if err != nil {
		return false, err
	}
	symbol, blues, err := ParseSymbolByXpath(strings.NewReader(string(content)))
	if err != nil {
		return false, err
	}
	tx := database.MySQL.NewSession()
	tx.Begin()
	for _, i := range blues {
		if _, dbError := tx.InsertOne(&i); dbError != nil {
			tx.Rollback()
			return false, dbError
		}
		symbol.BlueVersions = append(symbol.BlueVersions, i.Id)
	}
	if _, dbError := tx.InsertOne(&symbol); dbError != nil {
		tx.Rollback()
		return false, dbError
	}
	tx.Commit()
	return true, nil
}

func ParseSymbolByXpath(reader io.Reader) (models.Symbol, []models.Blue, error) {
	var symbol models.Symbol
	var blues []models.Blue
	doc, err := htmlquery.Parse(reader)
	if err != nil {
		return symbol, nil, err
	}
	list := htmlquery.Find(doc, `//div[@class="content content_wide"]/div`)

	for index, i := range list {
		img := htmlquery.FindOne(i, "/img")
		if index == 0 {
			symbol.SymbolTextImage = assistance.FullURL(htmlquery.SelectAttr(img, "src"))
		}
		if index == 1 {
			symbol.SymbolTextShort = htmlquery.InnerText(htmlquery.FindOne(i, "/h5"))
			symbol.SymbolText = htmlquery.InnerText(htmlquery.FindOne(i, "/p"))
		}
		if index == 2 {
			symbol.SymbolDescriptionImage = assistance.FullURL(htmlquery.SelectAttr(img, "src"))
		}
		if index == 3 {
			symbol.SymbolDescriptionShort = htmlquery.InnerText(htmlquery.FindOne(i, "/h5"))
			symbol.SymbolDescription = htmlquery.InnerText(htmlquery.FindOne(i, "/p"))
		}
		if index == 4 {
			symbol.SymbolAnimalImage = assistance.FullURL(htmlquery.SelectAttr(img, "src"))
		}
		if index == 5 {
			symbol.SymbolAnimalShort = htmlquery.InnerText(htmlquery.FindOne(i, "/h5"))
			symbol.SymbolAnimalDescription = htmlquery.InnerText(htmlquery.FindOne(i, "/p"))
		}
		if index == 6 {
			contentList := htmlquery.Find(i, "/div")
			for _, j := range contentList {
				var blue models.Blue
				blue = models.Blue{
					Image: assistance.FullURL(htmlquery.SelectAttr(htmlquery.FindOne(j, `/div[@class="img"]/img`), "src")),
				}
				text := htmlquery.FindOne(j, `/div[@class="text"]`)
				blue.Short = htmlquery.InnerText(htmlquery.FindOne(text, `/div/span[@class="point"]`))
				blue.EnName = htmlquery.InnerText(htmlquery.FindOne(text, `/div/span[@class="title"]/em`))
				blue.ChName = assistance.GetWordsHan(htmlquery.InnerText(htmlquery.FindOne(text, "/div/span[2]")))
				blue.Description = strings.Trim(assistance.GetWordsHan(htmlquery.InnerText(text)), blue.ChName)
				blues = append(blues, blue)
			}
		}

	}

	return symbol, blues, nil
}
