package admin

import (
	"gin_xiaomi_shop/database/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = mysql.DB

type baseService struct {
}

func (baseService) Success(ctx *gin.Context, message string, redirectUrl string) {
	ctx.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}

func (baseService) Error(ctx *gin.Context, message string, redirectUrl string) {
	ctx.HTML(http.StatusOK, "admin/public/error.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}
