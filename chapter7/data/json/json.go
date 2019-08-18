package jsonExplain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func ParseJSON() {
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Println(err)
		return
	}
	var result ResultForJSON
	err = json.Unmarshal(file, &result)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(result)
}

func MarshalJSON() {
	var object ResultForJSON
	object.Data.Directors = []string{"郑伟文", "陈家霖"}
	object.Data.Casts = []string{"肖战", "王一博", "孟子义", "宣璐", "于斌"}
	object.Data.Title = "陈情令"
	object.Data.Rate = "7.7"
	object.Data.Star = "40"
	object.Data.Cover = 3000
	object.Data.URL = "https://movie.douban.com/subject/27195020/"

	content, err := json.Marshal(object)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(content))

}
