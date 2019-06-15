package activity

import (
	"net/http"
	"time"

	"GopherBook/Chapter5/BeeQuick.v1/pkg/error.v1"

	"GopherBook/Chapter5/BeeQuick.v1/pkg/database.v1"

	"GopherBook/Chapter5/BeeQuick.v1/model/v1"

	"GopherBook/Chapter5/BeeQuick.v1/src/make_response"
	"github.com/kataras/iris"
)

func createOneActivityHandle(ctx iris.Context) {
	var param CreateActivityParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(
			http.StatusBadRequest, err.Error(), true,
		))
		return
	}

	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(
			http.StatusBadRequest, err.Error(), true,
		))
		return
	}

	var (
		start, end time.Time
		err        error
	)
	start, end, err = param.timeHandle()
	if err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	tx := database_v1.BeeQuickDatabase.NewSession()
	defer tx.Close()
	var shops []model_v1.Shop
	if dbError := tx.In("id", param.ShopIds).Find(&shops); dbError != nil {
		ctx.JSON(make_response.MakeResponse(
			http.StatusBadRequest, error_v1.ErrorRecordNotFound, true,
		))
		return
	}
	var activity model_v1.Activity
	activity = model_v1.Activity{
		Name:    param.Name,
		Title:   param.Title,
		Start:   start,
		End:     end,
		ShopIds: param.ShopIds,
		Status:  model_v1.DOING,
	}
	if param.Avatar != "" {
		activity.Avatar = param.Avatar
	}

	if _, dbError := tx.InsertOne(&activity); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(
			http.StatusBadRequest, dbError.Error(), true,
		))
		return
	}

	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, activity.Serializer(), false))

}
