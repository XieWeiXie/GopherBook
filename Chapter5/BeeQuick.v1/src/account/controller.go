package account

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"GopherBook/Chapter5/BeeQuick.v1/pkg/database.v1"

	"GopherBook/Chapter5/BeeQuick.v1/model/v1"
	"GopherBook/Chapter5/BeeQuick.v1/pkg/error.v1"

	"github.com/kataras/iris"
)

func registerProcessor(param RegisterParam) (model_v1.AccountGroupVip, error) {
	var (
		account model_v1.AccountGroupVip
		errV1   error_v1.ErrorV1
	)
	if err := param.Valid().Struct(param); err != nil {
		return account, error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Message: "param not valid",
			Detail:  "请求参数校验不通过，请检查参数",
		}
	}

	var vipMember model_v1.VipMember

	if _, dbErr := database_v1.BeeQuickDatabase.Where("level_name = ?", strings.ToUpper("v0")).Get(&vipMember); dbErr != nil {
		return account, error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Message: dbErr.Error(),
			Detail:  "会员等级未存在",
		}
	}
	account.VipMember = vipMember

	hashPassword, _ := generateFromPassword(param.Password, 8)
	hashToken := generateToken(20)

	account.Account = model_v1.Account{
		Phone:       param.Phone,
		Password:    string(hashPassword),
		Token:       hashToken,
		Points:      0,
		VipMemberID: vipMember.ID,
		VipTime:     time.Now(),
	}

	if _, err := database_v1.BeeQuickDatabase.InsertOne(&account.Account); err != nil {
		errV1 = error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Detail:  "用户注册发生错误",
		}
		return account, errV1
	}
	return account, nil
}

func registerHandle(ctx iris.Context) {
	var param RegisterParam
	err := ctx.ReadJSON(&param)
	if err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, err, true))
		return
	}
	account, err := registerProcessor(param)
	if err != nil {
		ctx.JSON(makeResponse(http.StatusBadRequest, err, true))
		return
	}
	ctx.JSON(iris.Map(makeResponse(http.StatusOK, account.SerializerForGroup(), false)))

}

func signProcessor(param RegisterParam) (model_v1.Account, error) {

	var (
		account model_v1.Account
		err     error
	)

	if err := param.Valid().Struct(param); err != nil {
		err = error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "登录参数校验失败",
			Message: err.Error(),
		}
		return account, err
	}

	if _, err := database_v1.BeeQuickDatabase.Where("phone = ?", param.Phone).Limit(1).Get(&account); err != nil {
		err = error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "账号未注册",
			Message: err.Error(),
		}
		return account, err
	}
	if !compareHashAndPassword([]byte(account.Password), []byte(param.Password)) {
		err = error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "密码错误",
			Message: "password not correct",
		}
		return account, err
	}

	return account, nil
}

func signHandle(ctx iris.Context) {
	var param RegisterParam
	err := ctx.ReadJSON(&param)
	if err != nil {
		ctx.JSON(iris.Map(makeResponse(http.StatusBadRequest, err, true)))
		return
	}
	account, err := signProcessor(param)
	if err != nil {
		ctx.JSON(iris.Map(makeResponse(http.StatusBadRequest, err, true)))
		return
	}

	ctx.JSON(iris.Map(makeResponse(http.StatusOK, account.Serializer(), false)))

}

func logoutHandle(ctx iris.Context) {
	account := ctx.Values().Get("current_admin")
	ctx.JSON(iris.Map(makeResponse(http.StatusOK, account.(model_v1.Account).Serializer(), false)))
}

func getAccountProcessor(id uint) (model_v1.Account, error) {
	var (
		account model_v1.Account
		err     error
	)
	if _, dbError := database_v1.BeeQuickDatabase.ID(id).Get(&account); dbError != nil {
		err = error_v1.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "记录未存在",
			Message: dbError.Error(),
		}
		return account, err
	}
	return account, nil
}

func getAccountHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	account, err := getAccountProcessor(id)
	if err != nil {
		ctx.JSON(iris.Map(makeResponse(http.StatusBadRequest, err, true)))
		return
	}
	fmt.Println(account)
	ctx.JSON(iris.Map{
		"data": account,
	})
	ctx.JSON(iris.Map(makeResponse(http.StatusOK, account.Serializer(), false)))
}
