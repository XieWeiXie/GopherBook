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
			ctx.JSON(makeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}
	if param.Question != "" {
		rule.Question = param.Question
		//if _, dbError := session.ID(rule.ID).Cols("question")
	}

}

func ruleGetAllHandler(ctx iris.Context) {}

func ruleGetOneHandler(ctx iris.Context) {

}
