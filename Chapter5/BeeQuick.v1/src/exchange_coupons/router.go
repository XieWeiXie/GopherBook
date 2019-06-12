package exchange_coupons

import "github.com/kataras/iris"

type ControllerExchangeCoupon struct {
}

var Default = ControllerExchangeCoupon{}

func (controller ControllerExchangeCoupon) Register(application *iris.Application, path string) {
	middleware := func(ctx iris.Context) {
		ctx.Next()
	}

	exchangeCoupon := application.Party(path, middleware)
	exchangeCoupon.Get("/coupons", getCouponsHandler)
	exchangeCoupon.Get("/coupons/{account_id:int}", getCouponWithAccountHandler)
	exchangeCoupon.Post("/coupon", postCouponHandler)
	exchangeCoupon.Patch("/coupon/{coupon_id:int}", patchCouponHandler)
	exchangeCoupon.Post("/coupon_to_account/{account_id:int}", postCouponToAccountHandler)
}
