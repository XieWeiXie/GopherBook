package shop

import (
	"net/http"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/error.v1"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/database.v1"

	"github.com/kataras/iris"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/model/v1"
)

func createOneShopHandler(ctx iris.Context) {
	var param PostShopParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(makeResponse(http.StatusNotFound, err.Error(), true))
		return
	}

	var province model_v1.Province
	if ok, _ := database_v1.BeeQuickDatabase.ID(param.ProvinceId).Get(&province); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	var shop model_v1.Shop
	shop = model_v1.Shop{
		Name:       param.Name,
		ProvinceId: int64(province.ID),
		Location:   param.Location,
		Province:   province,
	}
	database_v1.BeeQuickDatabase.InsertOne(&shop)
	ctx.JSON(makeResponse(http.StatusOK, shop.Serializer(), false))

}

func patchOneShopHandler(ctx iris.Context) {
	var param PostShopParam
	id, _ := ctx.Params().GetInt("shop_id")
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	var shop model_v1.Shop

	tx := database_v1.BeeQuickDatabase.NewSession()
	if ok, _ := tx.ID(id).Get(&shop); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	var province model_v1.Province
	if ok, _ := tx.ID(param.ProvinceId).Get(&province); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	shop.ProvinceId = int64(province.ID)
	shop.Province = province
	shop.Location = param.Location
	shop.Name = param.Name

	if _, dbError := tx.Update(&shop); dbError != nil {
		tx.Rollback()
		ctx.JSON(makeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(makeResponse(http.StatusOK, shop.Serializer(), false))

}

func getOneShopHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("shop_id")
	var shop model_v1.Shop

	if ok, _ := database_v1.BeeQuickDatabase.ID(id).Get(&shop); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	var province model_v1.Province
	if ok, _ := database_v1.BeeQuickDatabase.ID(shop.ProvinceId).Get(&province); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	shop.Province = province
	ctx.JSON(makeResponse(http.StatusOK, shop.Serializer(), false))
}

func getAllShopHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("province_id")
	var province model_v1.Province

	if ok, _ := database_v1.BeeQuickDatabase.ID(id).Get(&province); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	var newProvince []model_v1.Province
	if dbError := database_v1.BeeQuickDatabase.Where("ad_code like ?", province.AdCode[:2]+"%").Find(&newProvince); dbError != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	var shops []model_v1.Shop
	var provinceIds []int
	if len(newProvince) != 0 {
		for _, i := range newProvince {
			provinceIds = append(provinceIds, int(i.ID))
		}
	}
	if dbError := database_v1.BeeQuickDatabase.In("province_id", provinceIds).Find(&shops); dbError != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	var results []model_v1.ShopSerializer
	for _, i := range shops {
		var p model_v1.Province
		database_v1.BeeQuickDatabase.ID(i.ProvinceId).Get(&p)
		i.Province = p
		results = append(results, i.Serializer())
	}
	ctx.JSON(makeResponse(http.StatusOK, results, false))
}
