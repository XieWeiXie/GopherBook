package exchange_coupons

import (
	"log"
	"net/http"

	"GopherBook/Chapter5/BeeQuick.v1/model/v1"
	"GopherBook/Chapter5/BeeQuick.v1/pkg/database.v1"
	"GopherBook/Chapter5/BeeQuick.v1/pkg/error.v1"
	"github.com/kataras/iris"
)

func getCouponsProcessor(param CouponsParam) ([]model_v1.ExchangeCoupon, error) {
	var (
		result []model_v1.ExchangeCoupon
	)
	if ok, err := param.Suitable(); !ok || err != nil {
		return result, error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "请求参数不合法",
			Message: err.Error(),
		}
	}
	if dbError := database_v1.BeeQuickDatabase.Find(&result); dbError != nil {
		return result, error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "无记录",
			Message: dbError.Error(),
		}
	}
	return result, nil

}

func getCouponsHandler(ctx iris.Context) {
	var param CouponsParam
	log.Println("param", param, ctx.URLParam("return"))
	param.Return = ctx.URLParam("return")
	results, err := getCouponsProcessor(param)
	if err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, err, true))
		return
	}
	var data []model_v1.ExchangeCouponSerializer
	if len(results) == 0 {
		ctx.JSON(makeResponse(http.StatusOK, make([]model_v1.ExchangeCouponSerializer, 0), false))
		return
	}
	for _, i := range results {
		data = append(data, i.Serializer(""))
	}
	ctx.JSON(makeResponse(http.StatusOK, data, false))
}

func getCouponWithAccountHandler(ctx iris.Context) {}

func postCouponHandler(ctx iris.Context) {}

func patchCouponHandler(ctx iris.Context) {}

func postCouponToAccountHandler(ctx iris.Context) {}
