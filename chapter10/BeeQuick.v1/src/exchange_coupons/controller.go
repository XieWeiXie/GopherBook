package exchange_coupons

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kataras/iris"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/model/v1"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/database.v1"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/error.v1"
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
	if dbError := database_v1.BeeQuickDatabase.OrderBy("id desc").Find(&result); dbError != nil {
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

func getCouponWithAccountHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("account_id")

	var account model_v1.Account
	if ok, _ := database_v1.BeeQuickDatabase.Where("id = ?", id).Get(&account); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	var account2ExchangeCoupons []model_v1.Account2ExchangeCoupon

	if err := database_v1.BeeQuickDatabase.Where("account_id = ?", account.ID).Find(&account2ExchangeCoupons); err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	var ids []int
	for _, i := range account2ExchangeCoupons {
		ids = append(ids, int(i.ExchangeCouponId))
	}

	var exchangeCoupons []model_v1.ExchangeCoupon
	if dbError := database_v1.BeeQuickDatabase.In("id", ids).Find(&exchangeCoupons); dbError != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	var results []model_v1.ExchangeCouponSerializer
	for _, i := range exchangeCoupons {
		results = append(results, i.Serializer(""))
	}
	ctx.JSON(makeResponse(http.StatusOK, results, false))
}

func postCouponHandler(ctx iris.Context) {
	var param PostCouponParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, err, true))
		return
	}

	valid := param.Valid()
	if err := valid.Struct(param); err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	var exchange model_v1.ExchangeCoupon
	var (
		start time.Time
		end   time.Time
	)
	if param.Start != "" {
		start, _ = toTime(param.Start)
	}
	if param.End != "" {
		end, _ = toTime(param.End)
	}
	exchange = model_v1.ExchangeCoupon{
		Name:  param.Name,
		Price: param.Price,
		Total: param.Total,
		Start: start,
		End:   end,
		Type:  param.Type,
		Token: param.Token,
	}
	session := database_v1.BeeQuickDatabase.NewSession()
	session.Begin()

	if _, dbErr := session.InsertOne(&exchange); dbErr != nil {
		session.Rollback()
		ctx.JSON(makeResponse(http.StatusBadRequest, dbErr, true))
		return
	}
	session.Commit()
	ctx.JSON(makeResponse(http.StatusOK, exchange.Serializer(""), false))
}

func patchCouponHandler(ctx iris.Context) {
	var param PatchCouponParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.ReadJSON(makeResponse(http.StatusBadRequest, err, true))
		return
	}
	id, _ := ctx.Params().GetInt("coupon_id")
	fmt.Println(id)
	var exchangeCoupon model_v1.ExchangeCoupon
	if ok, _ := database_v1.BeeQuickDatabase.Where("id = ?", id).Get(&exchangeCoupon); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, fmt.Errorf("record not found").Error(), true))
		return
	}
	if exchangeCoupon.ID == 0 {
		ctx.JSON(makeResponse(http.StatusBadRequest, fmt.Errorf("record not found"), true))
		return
	}
	if param.Name != "" {
		exchangeCoupon.Name = param.Name
	}
	if param.Price != 0 && param.Total != 0 {
		exchangeCoupon.Price = param.Price
		exchangeCoupon.Total = param.Total
	}
	if param.Start != "" && param.End != "" {
		exchangeCoupon.Start, _ = toTime(param.Start)
		exchangeCoupon.End, _ = toTime(param.End)
	}
	database_v1.BeeQuickDatabase.ID(exchangeCoupon.ID).Update(&exchangeCoupon)
	ctx.JSON(makeResponse(http.StatusOK, exchangeCoupon.Serializer(""), false))
}

func postCouponToAccountHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("account_id")
	var param PostCouponToAccount
	err := ctx.ReadJSON(&param)
	if err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	var exchangeCoupon model_v1.ExchangeCoupon
	var account model_v1.Account

	if ok, _ := database_v1.BeeQuickDatabase.Where("id = ?", param.ExchangeCouponId).Get(&exchangeCoupon); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, fmt.Errorf("record not found"), true))
		return
	}
	if ok, _ := database_v1.BeeQuickDatabase.Where("id = ?", id).Get(&account); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, fmt.Errorf("record not found"), true))
		return
	}

	var account2ExchangeCoupon model_v1.Account2ExchangeCoupon
	account2ExchangeCoupon = model_v1.Account2ExchangeCoupon{
		ExchangeCouponId: int64(exchangeCoupon.ID),
		AccountId:        int64(account.ID),
		Status:           model_v1.NEW,
	}
	database_v1.BeeQuickDatabase.InsertOne(&account2ExchangeCoupon)
	ctx.JSON(makeResponse(http.StatusOK, exchangeCoupon.Serializer(""), false))
}
