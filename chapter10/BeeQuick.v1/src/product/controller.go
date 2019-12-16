package product

import (
	"net/http"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/error.v1"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/database.v1"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/model/v1"

	"github.com/kataras/iris"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/make_response"
)

func createProductHandle(ctx iris.Context) {
	var param CreateProductParam

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
	var product model_v1.Product
	product = model_v1.Product{
		ShopId:        param.ShopId,
		Name:          param.Name,
		Avatar:        param.Avatar,
		Price:         param.Price,
		Discount:      param.Discount,
		Specification: param.Specification,
		BrandId:       param.BrandId,
		UnitsId:       param.UnitId,
		TagsId:        param.TagId,
		Period:        param.Period,
	}
	var Brand model_v1.Brands
	if ok, dbError := tx.ID(param.BrandId).Get(&Brand); dbError != nil || !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	var Units model_v1.Units
	if ok, dbError := tx.ID(param.UnitId).Get(&Units); dbError != nil || !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	product.Units = Units

	var Tag model_v1.Tags
	if ok, dbError := tx.ID(param.TagId).Get(&Tag); dbError != nil || !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	var Shop model_v1.Shop
	if ok, dbError := tx.ID(param.ShopId).Get(&Shop); dbError != nil || !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	if _, dbError := tx.InsertOne(&product); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	var p2t model_v1.Product2Tags

	p2t = model_v1.Product2Tags{
		TagsId:    int64(Tag.ID),
		ProductId: int64(product.ID),
	}

	if ok, _ := tx.Where("tags_id = ? AND product_id = ?", p2t.TagsId, p2t.ProductId).Get(&model_v1.Product2Tags{}); !ok {
		if _, dbError := tx.InsertOne(&p2t); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}

	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, product.Serializer(), false))

}

func getOneProductHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("product_id")

	var product model_v1.Product
	if ok, dbError := database_v1.BeeQuickDatabase.ID(id).Get(&product); dbError != nil || !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	var shop model_v1.Shop
	database_v1.BeeQuickDatabase.ID(product.ShopId).Get(&shop)
	var tag model_v1.Tags
	database_v1.BeeQuickDatabase.ID(product.TagsId).Get(&tag)
	var brands model_v1.Brands
	database_v1.BeeQuickDatabase.ID(product.BrandId).Get(&brands)
	var units model_v1.Units
	database_v1.BeeQuickDatabase.ID(product.UnitsId).Get(&units)
	product.Units = units
	product.Shop = shop
	product.Brands = brands
	product.Tags = tag
	ctx.JSON(make_response.MakeResponse(http.StatusOK, product.Serializer(), false))
}

func getAllProductHandle(ctx iris.Context) {
	returnAll := ctx.URLParamDefault("return", "all_list")

	var (
		products []model_v1.Product
		count    int64
		dbError  error
	)

	if count, dbError = database_v1.BeeQuickDatabase.FindAndCount(&products); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	if returnAll == "all_count" {
		var results = make(map[string]interface{})
		results["count"] = count
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}

	var results []model_v1.ProductSerializer
	for _, i := range products {
		results = append(results, i.Serializer())
	}

	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))

}

func patchOneProductHandle(ctx iris.Context) {
	var param PatchOneParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	id, _ := ctx.Params().GetInt("product_id")

	tx := database_v1.BeeQuickDatabase.NewSession()
	defer tx.Close()
	tx.Begin()

	var product model_v1.Product
	if _, dbError := tx.ID(id).Get(&product); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	product.Name = param.Name
	product.Price = param.Price
	product.Discount = param.Discount

	if _, dbError := tx.ID(product.ID).Update(&product); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, product.Serializer(), false))

}

func postMultiplyProductHandle(ctx iris.Context) {

	var param CreateMultiplyParam
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

	var productIds []int64
	for _, i := range param.Data {
		var tempProduct model_v1.Product
		tempProduct = model_v1.Product{
			ShopId:        i.ShopId,
			Name:          i.Name,
			Avatar:        i.Avatar,
			Price:         i.Price,
			Discount:      i.Discount,
			Specification: i.Specification,
			BrandId:       i.BrandId,
			TagsId:        i.TagId,
			UnitsId:       i.UnitId,
			Period:        i.Period,
		}
		if ok, _ := tx.ID(i.ShopId).Get(&model_v1.Shop{}); !ok {
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
			return
		}
		if ok, _ := tx.ID(i.BrandId).Get(&model_v1.Brands{}); !ok {
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
			return
		}
		var tag model_v1.Tags
		if ok, _ := tx.ID(i.TagId).Get(&tag); !ok {
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
			return
		}
		var units model_v1.Units
		if ok, _ := tx.ID(i.UnitId).Get(&units); !ok {
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
			return
		}
		tempProduct.Units = units

		if _, dbError := tx.InsertOne(&tempProduct); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
			return
		}

		var p2t model_v1.Product2Tags
		p2t = model_v1.Product2Tags{
			TagsId:    int64(tag.ID),
			ProductId: int64(tempProduct.ID),
		}

		if ok, _ := tx.Where("tags_id = ? AND product_id = ?", p2t.TagsId, p2t.ProductId).Get(&model_v1.Product2Tags{}); !ok {
			if _, dbError := tx.InsertOne(&p2t); dbError != nil {
				tx.Rollback()
				ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
				return
			}
		}

		productIds = append(productIds, int64(tempProduct.ID))
	}

	var products []model_v1.Product
	if dbError := tx.In("id", productIds).Find(&products); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	tx.Commit()

	var results []model_v1.ProductSerializer
	for _, i := range products {
		results = append(results, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
}

func deleteProductHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("product_id")

	var product model_v1.Product
	tx := database_v1.BeeQuickDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	if ok, _ := tx.ID(id).Get(&product); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	if _, dbError := tx.ID(id).Delete(&product); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	var shop model_v1.Shop
	database_v1.BeeQuickDatabase.ID(product.ShopId).Get(&shop)
	var tag model_v1.Tags
	database_v1.BeeQuickDatabase.ID(product.TagsId).Get(&tag)
	var brands model_v1.Brands
	database_v1.BeeQuickDatabase.ID(product.BrandId).Get(&brands)
	var units model_v1.Units
	database_v1.BeeQuickDatabase.ID(product.UnitsId).Get(&units)
	product.Units = units
	product.Shop = shop
	product.Brands = brands
	product.Tags = tag
	var p2ts []model_v1.Product2Tags
	tx.Where("product_id = ?", product.ID).Find(&p2ts)
	for _, i := range p2ts {
		tx.Where("product_id = ?", i.ProductId).Delete(&model_v1.Product2Tags{})
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, product.Serializer(), false))
}
