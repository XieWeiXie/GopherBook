package province

import (
	"fmt"
	"net/http"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/error.v1"

	"github.com/kataras/iris"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/model/v1"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/database.v1"
)

func getProvinceHandler(ctx iris.Context) {
	var param GetProvinceParam
	returnAll := ctx.URLParamDefault("return", "all_list")
	param.Return = returnAll
	level := ctx.URLParam("level")
	if level != "" {
		param.Level = level
	}

	if err := param.Valid(); err != nil {
		ctx.JSON(makeResponse(http.StatusNotFound, err.Error(), true))
		return
	}

	query := database_v1.BeeQuickDatabase.NewSession()
	if level != "" {
		query = query.Where("level = ?", level)
	}

	var provinces []model_v1.Province
	var count int64
	var dbError error
	if count, dbError = query.FindAndCount(&provinces); dbError != nil {
		ctx.JSON(makeResponse(http.StatusNotFound, dbError.Error(), true))
		return
	}
	if returnAll == "all_count" {
		var countMap = make(map[string]int64)
		countMap["count"] = count
		ctx.JSON(makeResponse(http.StatusOK, countMap, false))
		return
	}
	var results []model_v1.ProvinceSerializer
	for _, i := range provinces {
		results = append(results, i.Serializer())
	}
	ctx.JSON(makeResponse(http.StatusOK, results, false))
}

func getOneHandler(ctx iris.Context) {
	name := ctx.URLParam("name")

	var province model_v1.Province

	fmt.Println(name)
	if name == "" {
		ctx.JSON(makeResponse(http.StatusBadRequest, fmt.Errorf("url param name should not be null"), true))
		return
	}
	if ok, _ := database_v1.BeeQuickDatabase.Where("name like ?", "%"+name+"%").Get(&province); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	var child []model_v1.Province
	var results = make(map[string]interface{})
	if province.Level == "province" {
		fmt.Println(province.AdCode[:1] + "%")
		database_v1.BeeQuickDatabase.Where("ad_code like ? AND level = ?", province.AdCode[:2]+"%", "city").Find(&child)
	} else if province.Level == "city" {
		database_v1.BeeQuickDatabase.Where("city_code = ? AND level = ?", province.CityCode, "district").Find(&child)
	}
	if len(child) == 0 {
		ctx.JSON(makeResponse(http.StatusOK, province.Serializer(), false))
		return
	}
	results[province.Level] = province.Serializer()
	var childResults []model_v1.ProvinceSerializer
	for _, i := range child {
		childResults = append(childResults, i.Serializer())
	}
	results["child"] = childResults
	ctx.JSON(makeResponse(http.StatusOK, results, false))
}
