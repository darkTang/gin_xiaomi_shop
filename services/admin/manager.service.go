package admin

import (
	"gin_xiaomi_shop/common"
	"gin_xiaomi_shop/models/admin"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ManagerService struct {
	baseService
}

func (ManagerService) Index(ctx *gin.Context) {
	var managerList []admin.Manager
	db.Preload("Role").Find(&managerList)
	ctx.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}

func (ms ManagerService) Add(ctx *gin.Context) {
	// 获取所有角色
	var roleList []admin.Role
	db.Find(&roleList)
	ctx.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

func (ms ManagerService) DoAdd(ctx *gin.Context) {
	roleId := ctx.PostForm("roleId")
	username := strings.TrimSpace(ctx.PostForm("username"))
	password := strings.TrimSpace(ctx.PostForm("password"))
	mobile := strings.TrimSpace(ctx.PostForm("mobile"))
	email := strings.TrimSpace(ctx.PostForm("email"))
	if len(username) < 2 || len(password) < 6 {
		ms.Error(ctx, "用户名或者密码的长度不合法", "/admin/manager/add")
		return
	}
	var manager admin.Manager
	if err := db.Where("username = ?", username).First(&manager).Error; err == nil {
		ms.Error(ctx, "管理员已存在", "/admin/manager/add")
	} else {
		manager = admin.Manager{
			Username: username,
			Password: common.Md5(password),
			Mobile:   mobile,
			Email:    email,
			RoleId:   roleId,
		}
		if err := db.Create(&manager).Error; err == nil {
			ms.Success(ctx, "增加管理员成功", "/admin/manager")
		} else {
			ms.Error(ctx, "增加管理员失败，请重试", "/admin/manager/add")
		}
	}
}

func (ms ManagerService) Edit(ctx *gin.Context) {
	// 获取管理员和角色
	id := ctx.Query("id")
	var roleList []admin.Role
	var managerInfo admin.Manager
	db.Find(&roleList)
	if err := db.Preload("Role").Where("id = ?", id).First(&managerInfo).Error; err != nil {
		ms.Error(ctx, "查询角色失败", "/admin/manager")
	}
	ctx.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"managerInfo": managerInfo,
		"roleList":    roleList,
	})
}

func (ms ManagerService) DoEdit(ctx *gin.Context) {
	id := ctx.PostForm("id")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	mobile := ctx.PostForm("mobile")
	roleId := ctx.PostForm("roleId")
	var manager admin.Manager
	db.Where("id = ?", id).Find(&manager)
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = roleId
	if strings.TrimSpace(password) != "" {
		if len(password) < 6 {
			ms.Error(ctx, "密码的长度不合法，密码长度不能小于6位", "/admin/manager/edit?id="+id)
		} else {
			manager.Password = common.Md5(password)
		}
	}
	if err := db.Save(&manager).Error; err == nil {
		ms.Success(ctx, "修改成功", "/admin/manager")
	}
}

func (ms ManagerService) Delete(ctx *gin.Context) {
	id := ctx.Query("id")
	var manager admin.Manager
	if err := db.Where("id = ?", id).Delete(&manager).Error; err == nil {
		ms.Success(ctx, "删除成功", "/admin/manager")
	} else {
		ms.Error(ctx, "删除失败", "/admin/manager")
	}
}
