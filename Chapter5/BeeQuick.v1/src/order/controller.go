package order

import (
	"fmt"
	"net/http"
	"strconv"

	"GopherBook/Chapter5/BeeQuick.v1/model/v1"
	"GopherBook/Chapter5/BeeQuick.v1/pkg/database.v1"
	"GopherBook/Chapter5/BeeQuick.v1/pkg/error.v1"
	"GopherBook/Chapter5/BeeQuick.v1/src/make_response"
	"github.com/kataras/iris"
)

func getOneOrderHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("order_id")

	var order model_v1.Order
	if ok, _ := database_v1.BeeQuickDatabase.ID(id).Get(&order); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	var account model_v1.Account
	if ok, _ := database_v1.BeeQuickDatabase.ID(order.AccountId).Get(&account); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	order.Account = account
	ctx.JSON(make_response.MakeResponse(http.StatusOK, order.Serializer(), false))
}

func getAllOrderHandle(ctx iris.Context) {
	returnAll := ctx.URLParamDefault("return", "all_list")

	var (
		count   int64
		orders  []model_v1.Order
		dbError error
	)

	if count, dbError = database_v1.BeeQuickDatabase.FindAndCount(&orders); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	if returnAll == "all_count" {
		var results = make(map[string]interface{})
		results["count"] = count
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}

	var resultsSerializer []model_v1.OrderSerializer
	for _, i := range orders {
		var account model_v1.Account
		database_v1.BeeQuickDatabase.ID(i.AccountId).Get(&account)
		i.Account = account
		resultsSerializer = append(resultsSerializer, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, resultsSerializer, false))
}

func patchOrderHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("order_id")
	var param PatchOrderParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	var order model_v1.Order
	tx := database_v1.BeeQuickDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	if ok, _ := tx.ID(id).Get(&order); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	var account model_v1.Account
	tx.ID(order.AccountId).Get(&account)

	value := func(val string) int {
		for k, v := range model_v1.STATUS_MAP_EN {
			if val == v {
				return k
			}
		}
		return -1
	}
	order.Status = value(param.Status)
	order.Account = account
	if _, dbError := tx.ID(id).Cols("status").Update(&order); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, order.Serializer(), false))

}

func postOrderHandle(ctx iris.Context) {
	var param PostOrderParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	var account model_v1.Account
	tx := database_v1.BeeQuickDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	if ok, _ := tx.ID(param.AccountId).Get(&account); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	var products []model_v1.Product
	if dbError := tx.In("id", param.ProductIds).Find(&products); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	var total float64
	for _, i := range products {
		total += i.Price * i.Discount
	}
	t, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", total), 2)
	var order model_v1.Order
	order = model_v1.Order{
		ProductIds: param.ProductIds,
		Status:     0,
		AccountId:  int64(param.AccountId),
		Account:    account,
		Total:      t,
	}

	if _, dbError := tx.InsertOne(&order); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, order.Serializer(), false))

}
