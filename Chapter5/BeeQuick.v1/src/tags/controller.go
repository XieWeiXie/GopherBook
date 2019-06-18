package tags

import (
	"net/http"

	"GopherBook/Chapter5/BeeQuick.v1/pkg/error.v1"

	"GopherBook/Chapter5/BeeQuick.v1/model/v1"
	"GopherBook/Chapter5/BeeQuick.v1/pkg/database.v1"

	"GopherBook/Chapter5/BeeQuick.v1/src/make_response"
	"github.com/kataras/iris"
)

func postTagHandle(ctx iris.Context) {
	var param CreateTagParam
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

	var tag model_v1.Tags
	tag.Name = param.Name
	if _, dbError := tx.InsertOne(&tag); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, tag.Serializer(), false))
}

func postTagMultiplyHandle(ctx iris.Context) {
	var param CreateTagsParam
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

	var tagIds []uint
	for _, i := range param.Data {
		var tempTag model_v1.Tags
		tempTag = model_v1.Tags{
			Name: i.Name,
		}
		if _, dbError := tx.InsertOne(&tempTag); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
		tagIds = append(tagIds, tempTag.ID)
	}

	var tags []model_v1.Tags
	if dbError := tx.In("id", tagIds).Find(&tags); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	var results []model_v1.TagSerializer
	for _, i := range tags {
		results = append(results, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
}

func patchTagHandle(ctx iris.Context) {
	var param CreateTagParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	id, _ := ctx.Params().GetInt("tag_id")

	tx := database_v1.BeeQuickDatabase.NewSession()
	defer tx.Close()
	tx.Begin()

	var tag model_v1.Tags
	if ok, _ := tx.ID(id).Get(&tag); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	tag.Name = param.Name
	if _, dbError := tx.ID(tag.ID).Cols("name").Update(&tag); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, tag.Serializer(), false))
}

func getTagHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("tag_id")

	var tag model_v1.Tags
	if ok, dbError := database_v1.BeeQuickDatabase.ID(id).Get(&tag); dbError != nil || !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	ctx.JSON(make_response.MakeResponse(http.StatusOK, tag.Serializer(), false))
}

func getTagsHandle(ctx iris.Context) {

	returnAll := ctx.URLParamDefault("return", "all_list")

	var (
		tags    []model_v1.Tags
		count   int64
		dbError error
	)

	if count, dbError = database_v1.BeeQuickDatabase.FindAndCount(&tags); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	if returnAll == "all_count" {
		var results = make(map[string]interface{})
		results["count"] = count
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}
	var results []model_v1.TagSerializer
	for _, i := range tags {
		results = append(results, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
}
