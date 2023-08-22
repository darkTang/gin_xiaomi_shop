package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainService struct {
}

func (MainService) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/main/index.html", nil)
}

func (MainService) Welcome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/main/welcome.html", nil)
}
