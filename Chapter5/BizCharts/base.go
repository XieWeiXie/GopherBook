package BizCharts

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type BaseTheme struct {
	Theme string `json:"theme"`
}

const (
	DEFAULT_THEME = "default"
	DARK_THEME    = "dark"
)

func (B *BaseTheme) SetTheme(theme string) {
	B.Theme = theme
}

type BaseData struct {
	Data interface{} `json:"data"`
}

func NewBaseData(data interface{}) *BaseData {

	return &BaseData{
		Data: data,
	}
}

func (B BaseData) Format() []interface{} {
	v := reflect.ValueOf(B.Data)
	var result []interface{}
	for i := 0; i < v.NumField(); i++ {
		//var temp = make(map[string]interface{})
		fmt.Println(reflect.TypeOf(i).Name(), reflect.ValueOf(i))
	}
	return result
}

type BaseLegend struct {
	Location string `json:"location"`
}

func (L *BaseLegend) SetLocation(position string) {
	L.Location = position
}

func ToMap(obj interface{}) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		tag := objT.Field(i).Tag
		tagJson := regexpHandle(fmt.Sprintf("%s", tag))
		data[tagJson] = objV.Field(i).Interface()
	}
	err = nil
	return
}

func regexpHandle(tag string) string {
	reg := `json:"(.*?)"`
	regx := regexp.MustCompile(reg)
	if strings.Contains(tag, "json") {
		result := regx.FindStringSubmatch(tag)
		if len(result) != 2 {
			return "-1"
		}
		return strings.TrimSpace(result[1])

	}
	return "-1"
}
