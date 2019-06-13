package rule

import (
	"net/http"

	"GopherBook/Chapter5/BeeQuick.v1/pkg/database.v1"

	"GopherBook/Chapter5/BeeQuick.v1/model/v1"

	"GopherBook/Chapter5/BeeQuick.v1/pkg/error.v1"
	"github.com/kataras/iris"
)

func ruleCreateOneHandler(ctx iris.Context) {
	var param PostRuleParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.ReadJSON(makeResponse(http.StatusNotFound, error_v1.ErrorBodyJson, true))
		return
	}
	if !param.notNull() {
		ctx.ReadJSON(makeResponse(http.StatusBadRequest, error_v1.ErrorBodyIsNull, true))
		return
	}
	if err := param.Valid(); err != nil {
		ctx.ReadJSON(makeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	var rule model_v1.RuleForExchangeOrCoupon
	rule = model_v1.RuleForExchangeOrCoupon{
		Question: param.Question,
		Answer:   param.Answer,
		Type:     param.Type,
	}
	database_v1.BeeQuickDatabase.InsertOne(&rule)
	ctx.JSON(makeResponse(http.StatusOK, rule.Serializer(), false))
}

func rulePatchOneHandler(ctx iris.Context) {

	var param PostRuleParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	id, _ := ctx.Params().GetInt("rule_id")
	var rule model_v1.RuleForExchangeOrCoupon
	if ok, _ := database_v1.BeeQuickDatabase.Where("id = ?", id).Get(&rule); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}

	session := database_v1.BeeQuickDatabase.NewSession()
	session.Begin()
	if param.Answer != "" {
		rule.Answer = param.Answer
		if _, dbError := session.ID(rule.ID).Cols("answer").Update(&rule); dbError != nil {
			session.Rollback()
			ctx.JSON(makeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}
	if param.Question != "" {
		rule.Question = param.Question
		if _, dbError := session.ID(rule.ID).Cols("question").Update(&rule); dbError != nil {
			session.Rollback()
			ctx.JSON(makeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}
	rule.Type = param.Type
	if _, dbError := session.ID(rule.ID).Cols("type").Update(&rule); dbError != nil {
		session.Rollback()
		ctx.JSON(makeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	session.Commit()
	ctx.JSON(makeResponse(http.StatusOK, rule.Serializer(), false))
}

func ruleGetAllHandler(ctx iris.Context) {
	p := ctx.URLParamDefault("return", "all_list")
	var param GetRuleParam
	param.Return = p
	if err := param.Valid(); err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	var rules []model_v1.RuleForExchangeOrCoupon
	if dbError := database_v1.BeeQuickDatabase.OrderBy("id").Desc("id").Find(&rules); dbError != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	if param.Return == "all_list" {
		var results []model_v1.RuleForExchangeOrCouponSerializer
		for _, i := range rules {
			results = append(results, i.Serializer())
		}
		ctx.JSON(makeResponse(http.StatusOK, results, false))
		return
	}
	if param.Return == "all_count" {
		var count = make(map[string]int)
		count["count"] = len(rules)
		ctx.JSON(makeResponse(http.StatusOK, count, false))
		return
	}

}

func ruleGetOneHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("rule_id")

	var rule model_v1.RuleForExchangeOrCoupon
	if ok, _ := database_v1.BeeQuickDatabase.ID(id).Get(&rule); !ok {
		ctx.JSON(makeResponse(http.StatusBadRequest, error_v1.ErrorRecordNotFound, true))
		return
	}
	ctx.JSON(makeResponse(http.StatusOK, rule.Serializer(), false))

}
