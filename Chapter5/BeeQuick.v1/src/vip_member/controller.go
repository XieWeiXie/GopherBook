package vip_member

import (
	"net/http"

	"GopherBook/Chapter5/BeeQuick.v1/model/v1"
	"GopherBook/Chapter5/BeeQuick.v1/pkg/database.v1"
	"GopherBook/Chapter5/BeeQuick.v1/pkg/error.v1"
	"github.com/kataras/iris"
)

func getVipMemberProscessor() ([]model_v1.VipMember, error) {
	var (
		vipMembers []model_v1.VipMember
		err        error
	)
	if dbErr := database_v1.BeeQuickDatabase.Find(&vipMembers); dbErr != nil {
		err = error_v1.ErrorV1{
			Code:    http.StatusNotFound,
			Message: dbErr.Error(),
			Detail:  "记录未找到",
		}
		return vipMembers, err
	}

	return vipMembers, nil

}

func getVipMemberHandle(ctx iris.Context) {
	results, err := getVipMemberProscessor()
	if err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	var resultSerializer []model_v1.VipMemberSerializer
	for _, i := range results {
		resultSerializer = append(resultSerializer, i.Serializer())
	}
	ctx.JSON(makeResponse(http.StatusOK, resultSerializer, false))
}

func getVipMemberOneHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	var vip model_v1.VipMember
	if _, dbError := database_v1.BeeQuickDatabase.Where("id = ?", id).Get(&vip); dbError != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, dbError, true))
		return
	}
	ctx.JSON(makeResponse(http.StatusOK, vip.Serializer(), false))
}

func patchVipMemberProcessor(id uint, param PatchVipMemberParam) (model_v1.VipMember, error) {
	var (
		result model_v1.VipMember
		err    error
	)
	if _, dbError := database_v1.BeeQuickDatabase.Where("id = ?", id).Get(&result); dbError != nil {
		err = error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "记录未找到",
			Message: dbError.Error(),
		}
		return result, err
	}

	if err := param.Valid().Struct(param); err != nil {
		err = error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "参数校验失败",
			Message: err.Error(),
		}
		return result, err
	}

	result.LevelName = param.Level
	result.Start = param.Start
	result.End = param.End
	result.Points = param.Points
	result.Period = param.Period
	result.ToValue = param.ToValue

	if _, dbError := database_v1.BeeQuickDatabase.ID(result.ID).Update(result); dbError != nil {
		err = error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "更新数据失败",
			Message: dbError.Error(),
		}
	}
	return result, nil

}

func patchVipMemberHandle(ctx iris.Context) {
	var param PatchVipMemberParam
	if err := ctx.ReadJSON(&param); err != nil {
		return
	}

	id, _ := ctx.Params().GetUint("id")
	result, err := patchVipMemberProcessor(id, param)
	if err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, err, true))
		return
	}
	ctx.JSON(makeResponse(http.StatusOK, result.Serializer(), false))
}
