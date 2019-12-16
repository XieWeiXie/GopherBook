package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/model/v1"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/database.v1"

	"github.com/kataras/iris"
)

func LoggerForProject(c iris.Context) {
	c.Application().Logger().Debugf("Path: %s | IP: %s | Time: %s",
		c.Path(), c.RemoteAddr(), time.Now().Format("2006-01-02 15:04:05"))
	c.Next()
}

func TokenForProject(c iris.Context) {
	token := c.GetHeader("Authorization")
	tokenList := strings.Split(token, " ")
	if len(tokenList) != 2 || tokenList[0] != "Bearer" {
		c.JSON(iris.Map{
			"code": http.StatusNotFound,
			"err":  "Header Add Authorization: Bearer xxx",
		})
		return
	}
	realToken := tokenList[1]
	var account model_v1.Account
	if _, err := database_v1.BeeQuickDatabase.Where("token = ?", realToken).Get(&account); err != nil {
		c.JSON(iris.Map{
			"code": http.StatusNotFound,
			"err":  err.Error(),
		})
		return
	}
	c.Values().Set("current_admin", account)
	c.Values().Set("current_admin_id", account.ID)
	c.Next()

}
