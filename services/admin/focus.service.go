package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FocusService struct {
}

func (FocusService) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/index.html", nil)
}

func (FocusService) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/add.html", nil)
}

func (FocusService) Edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/edit.html", nil)
}

func (FocusService) Delete(ctx *gin.Context) {
	ctx.String(http.StatusOK, "删除轮播图")
}
