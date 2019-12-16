package activity

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/error.v1"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/database.v1"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/model/v1"

	"github.com/kataras/iris"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/make_response"
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
	tx.Begin()
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

	for _, i := range shops {
		var shop2Activity model_v1.Shop2Activity
		shop2Activity = model_v1.Shop2Activity{
			ShopId:     int64(i.ID),
			ActivityId: int64(activity.ID),
		}
		if _, dbError := tx.InsertOne(&shop2Activity); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}

	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, activity.Serializer(), false))

}

func patchOneActivityHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("activity_id")
	var param PatchActivityParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	tx := database_v1.BeeQuickDatabase.NewSession()
	defer tx.Commit()
	tx.Begin()
	var activity model_v1.Activity
	if ok, _ := tx.ID(id).Get(&activity); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	activity = model_v1.Activity{
		Name:    param.Name,
		Title:   param.Title,
		ShopIds: param.ShopIds,
		Start:   toTime(param.Start),
		End:     toTime(param.End),
		Status:  param.Status,
		Avatar:  param.Avatar,
	}

	if _, dbError := tx.Update(&activity); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, activity.Serializer(), false))
}

func getOneActivityHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("activity_id")
	var activity model_v1.Activity
	if ok, dbError := database_v1.BeeQuickDatabase.ID(id).Get(&activity); dbError != nil || !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	ctx.JSON(make_response.MakeResponse(http.StatusOK, activity.Serializer(), false))

}

func getAllActivityHandle(ctx iris.Context) {

	var activities []model_v1.Activity
	status := ctx.URLParam("status")
	returnAll := ctx.URLParamDefault("return", "all_list")
	var param GetActivityParam
	param.Status = status
	param.ReturnAll = returnAll
	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	query := database_v1.BeeQuickDatabase.NewSession()
	defer query.Close()
	query.Begin()
	if param.Status != "" {
		key := func(status string) int {
			for k, v := range model_v1.ActivityStatusEn {
				if strings.ToUpper(v) == strings.ToUpper(status) {
					return k
				}
			}
			return -1
		}
		query = query.Where("status = ?", key(param.Status))
	}
	var (
		total int64
		err   error
	)

	if total, err = query.FindAndCount(&activities); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	if param.ReturnAll == "all_list" {
		var results []model_v1.ActivitySerializer
		for _, i := range activities {
			results = append(results, i.Serializer())
		}
		query.Commit()
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}

	if param.ReturnAll == "all_count" {
		var results = make(map[string]int64)
		results["count"] = total
		query.Commit()
		ctx.JSON(make_response.MakeResponse(
			http.StatusOK, results, false,
		))
	}

}

func getAllByShopIdActivityHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("shop_id")
	fmt.Println(id, "id")
	var (
		dbError error
	)
	var shop2Activity []model_v1.Shop2Activity
	if _, dbError = database_v1.BeeQuickDatabase.Where("shop_id = ?", id).FindAndCount(&shop2Activity); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	var activityIds []int64
	for _, i := range shop2Activity {
		activityIds = append(activityIds, i.ActivityId)
	}
	var activities []model_v1.Activity
	if dbError := database_v1.BeeQuickDatabase.In("id", activityIds).Find(&activities); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	var results []model_v1.ActivitySerializer
	for _, i := range activities {
		results = append(results, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))

}
