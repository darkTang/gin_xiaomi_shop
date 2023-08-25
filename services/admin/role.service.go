package admin

import (
	"gin_xiaomi_shop/models/admin"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RoleService struct {
	baseService
}

func (rs RoleService) Index(ctx *gin.Context) {
	var roleList []admin.Role
	if err := db.Find(&roleList).Error; err != nil {
		rs.Error(ctx, "获取角色失败，请重试", "/admin/role")
	}
	ctx.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

func (RoleService) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/role/add.html", nil)
}

func (rs RoleService) DoAdd(ctx *gin.Context) {
	title := strings.TrimSpace(ctx.PostForm("title"))
	description := strings.TrimSpace(ctx.PostForm("description"))
	var role admin.Role
	if title == "" {
		rs.Error(ctx, "角色标题不能为空", "/admin/role/add")
	} else {
		if err := db.Where("title = ?", title).First(&role).Error; err != nil {
			role.Title = title
			role.Description = description
			if err = db.Create(&role).Error; err == nil {
				rs.Success(ctx, "增加角色成功", "/admin/role")
			} else {
				rs.Error(ctx, "增加角色失败，请重试", "/admin/role/add")
			}
		} else {
			rs.Error(ctx, "角色标题已存在", "/admin/role/add")
		}
	}
}

func (rs RoleService) Edit(ctx *gin.Context) {
	id := ctx.Query("id")
	var roleInfo admin.Role
	if err := db.Where("id = ?", id).First(&roleInfo).Error; err != nil {
		rs.Error(ctx, "查询角色失败", "/admin/role")
	}
	ctx.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
		"roleInfo": roleInfo,
	})
}

func (rs RoleService) DoEdit(ctx *gin.Context) {
	var role admin.Role
	id := ctx.PostForm("id")
	title := strings.TrimSpace(ctx.PostForm("title"))
	description := strings.TrimSpace(ctx.PostForm("description"))
	if title == "" {
		rs.Error(ctx, "角色标题不能为空", "/admin/role/edit")
	} else {
		if err := db.Model(&role).Where("id = ?", id).Updates(admin.Role{
			Title:       title,
			Description: description,
		}).Error; err == nil {
			rs.Success(ctx, "修改角色成功", "/admin/role")
		} else {
			rs.Error(ctx, "修改角色失败", "/admin/role/edit")
		}
	}
}

func (rs RoleService) Delete(ctx *gin.Context) {
	id := ctx.Query("id")
	var role admin.Role
	if err := db.Where("id = ?", id).Delete(&role).Error; err == nil {
		rs.Success(ctx, "删除角色成功", "/admin/role")
	} else {
		rs.Error(ctx, "删除角色失败", "/admin/role")
	}
}
