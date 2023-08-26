package admin

import (
	"gin_xiaomi_shop/common"
	"gin_xiaomi_shop/models/admin"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RightsService struct {
	baseService
}

func (rs RightsService) Index(ctx *gin.Context) {
	var rightsList []admin.Rights
	db.Where("module_id = ?", 0).Preload("RightsItem").Find(&rightsList)
	ctx.HTML(http.StatusOK, "admin/rights/index.html", gin.H{
		"rightsList": rightsList,
	})
}

func (rs RightsService) Add(ctx *gin.Context) {
	var rightsList []admin.Rights
	db.Where("module_id = ?", 0).Find(&rightsList)
	ctx.HTML(http.StatusOK, "admin/rights/add.html", gin.H{
		"rightsList": rightsList,
	})
}

func (rs RightsService) DoAdd(ctx *gin.Context) {
	moduleName := strings.TrimSpace(ctx.PostForm("moduleName"))
	rightsType := uint8(common.Atoi(ctx.PostForm("type")))
	actionName := ctx.PostForm("actionName")
	description := ctx.PostForm("description")
	url := ctx.PostForm("url")
	moduleId := ctx.PostForm("moduleId")
	sort := common.Atoi(ctx.PostForm("sort"))
	status := uint8(common.Atoi(ctx.PostForm("status")))
	if moduleName == "" {
		rs.Error(ctx, "模块名称不能为空", "/admin/rights/add")
		return
	}
	rightsInfo := admin.Rights{
		ModuleName:  moduleName,
		Type:        rightsType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	if err := db.Create(&rightsInfo).Error; err == nil {
		rs.Success(ctx, "增加权限成功", "/admin/rights")
	} else {
		rs.Error(ctx, "增加权限失败", "/admin/rights/add")
	}
}
