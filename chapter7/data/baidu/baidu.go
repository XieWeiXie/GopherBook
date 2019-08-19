package baidu

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/jinzhu/gorm"

	"github.com/PuerkitoBio/goquery"

	"github.com/chromedp/chromedp"
	_ "github.com/go-sql-driver/mysql"
)

func GetBaiDu(url string) {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Fatalf))
	defer cancel()
	var response string
	err := chromedp.Run(ctx, Tasks(url, &response))
	if err != nil {
		log.Println(err)
		return
	}
	urls := Parse(response)
	now := time.Now()
	var results []ResultBaiDu
	for _, i := range urls {
		var childResponse string
		err := chromedp.Run(ctx, AnotherTasks(i, &childResponse))
		if err != nil {
			log.Println(err)
			return
		}
		results = append(results, AnotherParse(childResponse)...)

	}

	fmt.Println(time.Since(now))
	//SaveTxt(results)
	//SaveJSON(results)
	//SaveCSV(results)
	//SaveDB(results)
	SaveRedis(results)

}

func Tasks(url string, response *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("#flist", chromedp.ByQuery),
		chromedp.OuterHTML("body", response),
	}
}

func Parse(response string) []string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Println(err)
		return nil
	}
	//fmt.Println(doc.Html())
	var urls []string
	doc.Find("#flist div ul li").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		if v, ok := selection.Find("a").Attr("href"); ok {
			urls = append(urls, strings.Replace(v, ".", ROOT, 1))
		}
	})
	return urls

}

func AnotherTasks(url string, response *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("tbody", chromedp.ByQuery),
		chromedp.OuterHTML("body", response),
	}
}

func AnotherParse(response string) []ResultBaiDu {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Println(err)
		return nil
	}
	var results []ResultBaiDu
	doc.Find("tbody tr").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		if v, ok := selection.Attr("class"); ok {
			if v == "item-tr" {
				return
			}
		}
		var r ResultBaiDu
		keyword := selection.Find(`td[class="keyword"] a`).Eq(0)
		r.Keyword = strings.TrimSpace(keyword.Text())
		if v, ok := keyword.Attr("href"); ok {
			r.Href = v
		}
		r.Number, _ = strconv.Atoi(selection.Find(`td[class="last"] span`).Text())
		//fmt.Println(r)
		results = append(results, r)

	})
	return results
}

func SaveTxt(results []ResultBaiDu) {

	f, err := os.Open(FILE_NAME_TEXT)
	if err != nil {
		f, err = os.Create(FILE_NAME_TEXT)
	}
	var w *bufio.Writer
	w = bufio.NewWriter(f)
	for _, i := range results {
		c, err := json.Marshal(i)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(c)
		w.WriteString("\n")
	}
	w.Flush()

}

func SaveJSON(results []ResultBaiDu) {
	content, err := json.MarshalIndent(results, " ", " ")
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile(FILE_NAME_JSON, content, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

func SaveCSV(results []ResultBaiDu) {
	f, err := os.Open(FILE_NAME_CSV)
	if err != nil {
		f, err = os.Create(FILE_NAME_CSV)
	}
	header := []string{"KEY", "URL", "NUMBER"}
	var values [][]string

	for _, i := range results {
		var line []string
		line = append(line, i.Keyword, i.Href, strconv.Itoa(i.Number))
		values = append(values, line)
	}
	w := csv.NewWriter(f)
	w.Write(header)
	for _, i := range values {
		w.Write(i)
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		log.Println(err)
		return
	}
}

var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:admin123@/baidu?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		panic(err)
		return
	}
	DB = db
	DB.LogMode(true)
}
func SaveDB(results []ResultBaiDu) {
	DB.AutoMigrate(&ResultBaiDu{})

	for _, i := range results {
		var one ResultBaiDu
		if dbError := DB.Where("keyword = ?", i.Keyword).
			First(&one).Error; dbError != nil {
			one = i
			if dbError := DB.Save(&one).Error; dbError != nil {
				log.Println(dbError)
				return
			}
		} else {
			if dbError := DB.Model(&one).Updates(map[string]interface{}{
				"number": i.Number}).Error; dbError != nil {
				log.Println(dbError)
				return
			}
		}
	}
}

var REDIS redis.Conn

func init() {
	connect, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Println(err)
		return
	}
	REDIS = connect
}

func SaveRedis(results []ResultBaiDu) {
	for index, i := range results {
		reply, err := REDIS.Do("HMSET", fmt.Sprintf(BAIDUKEY+":%d", index),
			"key", i.Keyword, "href", i.Href, "number", i.Number)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(reply)

	}
}
