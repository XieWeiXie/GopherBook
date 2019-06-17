package unit

import (
	"net/http"

	"GopherBook/Chapter5/BeeQuick.v1/pkg/error.v1"

	"GopherBook/Chapter5/BeeQuick.v1/model/v1"
	"GopherBook/Chapter5/BeeQuick.v1/pkg/database.v1"

	"GopherBook/Chapter5/BeeQuick.v1/src/make_response"
	"github.com/kataras/iris"
)

func createUintHandle(ctx iris.Context) {
	var (
		param         CreateUintParam
		readJsonError error
		validError    error
	)
	readJsonError = ctx.ReadJSON(&param)
	validError = param.Valid()

	if readJsonError != nil && validError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, validError.Error(), true))
		return
	}

	tx := database_v1.BeeQuickDatabase.NewSession()
	defer tx.Close()
	tx.Begin()

	var results []model_v1.UintsSerializer
	for _, i := range param.Data {
		var u model_v1.Uints
		u = model_v1.Uints{
			Name:      i.Name,
			EnName:    i.EnName,
			ShortCode: i.Code,
		}
		if _, dbError := tx.InsertOne(&u); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
		results = append(results, u.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))

}

func patchUintHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("uint_id")

	var param PatchUintParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	tx := database_v1.BeeQuickDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	var u model_v1.Uints
	if ok, _ := tx.ID(id).Get(&u); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	u.Name = param.Name
	u.EnName = param.EnName
	u.ShortCode = param.Code

	if _, dbError := tx.Update(&u); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, u.Serializer(), false))
}

func getUintHandle(ctx iris.Context) {
	returnAll := ctx.URLParamDefault("return", "all_list")

	var param GetUintParam
	param.ReturnAll.ReturnAll = returnAll

	if error := param.Valid(); error != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error.Error(), true))
		return
	}
	var (
		us      []model_v1.Uints
		count   int64
		dbError error
	)

	if count, dbError = database_v1.BeeQuickDatabase.FindAndCount(&us); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	if returnAll == "all_list" {
		var results []model_v1.UintsSerializer
		for _, i := range us {
			results = append(results, i.Serializer())
		}
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}
	if returnAll == "all_count" {
		var results = make(map[string]interface{})
		results["count"] = count
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}

}
