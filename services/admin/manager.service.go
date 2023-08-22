package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ManagerService struct {
}

func (ManagerService) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/index.html", nil)
}

func (ManagerService) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/welcome.html", nil)
}

func (ManagerService) Edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/edit.html", nil)
}

func (ManagerService) Delete(ctx *gin.Context) {
	ctx.String(http.StatusOK, "删除管理员")
}
