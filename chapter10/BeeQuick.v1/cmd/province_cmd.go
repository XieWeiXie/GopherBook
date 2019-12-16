package cmd

import (
	"io/ioutil"
	"net/http"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/model/v1"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/database.v1"

	"github.com/tidwall/gjson"

	"github.com/spf13/cobra"
)

var api = "https://restapi.amap.com/v3/config/district?subdistrict=3&key=bb4198a1f146184af53322d424732f6b"

var provinceCMD = &cobra.Command{
	Use:        "province",
	ArgAliases: []string{"p", "-p", "P", "-P"},
	Run:        importData,
}

func importData(cmd *cobra.Command, args []string) {
	response, err := http.Get(api)
	if err != nil {
		return
	}
	database_v1.DataBaseInit()
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	jsonByte := gjson.ParseBytes(content)
	for _, i := range jsonByte.Get("districts").Array()[0].Get("districts").Array() {
		if i.Get("level").String() == "province" {

			var province model_v1.Province
			province = model_v1.Province{
				Name:   i.Get("name").String(),
				AdCode: i.Get("adcode").String(),
				Center: i.Get("center").String(),
				Level:  i.Get("level").String(),
			}
			//fmt.Println("Province", province)
			database_v1.BeeQuickDatabase.InsertOne(&province)
			if len(i.Get("districts").Array()) != 0 {
				for _, j := range i.Get("districts").Array() {
					if j.Get("level").String() == "city" {
						var city model_v1.Province
						city = model_v1.Province{
							Name:     j.Get("name").String(),
							AdCode:   j.Get("adcode").String(),
							Center:   j.Get("center").String(),
							Level:    j.Get("level").String(),
							CityCode: j.Get("citycode").String(),
						}
						//fmt.Println("city", city)
						database_v1.BeeQuickDatabase.InsertOne(&city)
						if len(j.Get("districts").Array()) != 0 {
							for _, k := range j.Get("districts").Array() {
								if k.Get("level").String() == "district" {
									var district model_v1.Province
									district = model_v1.Province{
										Name:     k.Get("name").String(),
										AdCode:   k.Get("adcode").String(),
										Center:   k.Get("center").String(),
										Level:    k.Get("level").String(),
										CityCode: k.Get("citycode").String(),
									}
									//fmt.Println("district", district)
									database_v1.BeeQuickDatabase.InsertOne(&district)
								}
							}
						}
					}

				}

			}
		}
	}

}
