package brand

import (
	"net/http"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/error.v1"

	"github.com/kataras/iris"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/model/v1"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/database.v1"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/make_response"
)

func createBrandHandle(ctx iris.Context) {
	var param CreateBrandParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	var brand model_v1.Brands
	brand = model_v1.Brands{
		ChName: param.Name,
		EnName: param.EnName,
	}
	tx := database_v1.BeeQuickDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	if _, dbError := tx.InsertOne(&brand); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, brand.Serializer(), false))
}

func patchBrandHandle(ctx iris.Context) {
	var param PatchBrandParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	id, _ := ctx.Params().GetInt("brand_id")
	var brand model_v1.Brands
	tx := database_v1.BeeQuickDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	if _, dbError := tx.ID(id).Get(&brand); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	if param.Name != "" {
		brand.ChName = param.Name
		if _, dbError := tx.ID(brand.ID).Cols("ch_name").Update(&brand); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}
	if param.EnName != "" {
		brand.EnName = param.EnName
		if _, dbError := tx.ID(brand.ID).Cols("en_name").Update(&brand); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, brand.Serializer(), false))

}

func getBrandHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("brand_id")
	var brand model_v1.Brands

	if ok, dbError := database_v1.BeeQuickDatabase.ID(id).Get(&brand); dbError != nil || !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, brand.Serializer(), false))
}

func getBrandsHandle(ctx iris.Context) {
	returnAll := ctx.URLParamDefault("return", "all_list")

	var (
		brands  []model_v1.Brands
		count   int64
		dbError error
	)

	if count, dbError = database_v1.BeeQuickDatabase.FindAndCount(&brands); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	if returnAll == "all_count" {
		var results = make(map[string]interface{})
		results["count"] = count
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}
	var results []model_v1.BrandsSerializer
	for _, i := range brands {
		results = append(results, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))

}

func createBrandsHandle(ctx iris.Context) {
	var param CreateBrandsParam
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
	var brandIds []uint
	for _, i := range param.Data {
		var temp model_v1.Brands
		temp = model_v1.Brands{
			ChName: i.Name,
			EnName: i.EnName,
		}
		if _, dbError := tx.InsertOne(&temp); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
		brandIds = append(brandIds, temp.ID)
	}

	var brands []model_v1.Brands
	if dbError := tx.In("id", brandIds).Find(&brands); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	var results []model_v1.BrandsSerializer
	for _, i := range brands {
		results = append(results, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))

}
