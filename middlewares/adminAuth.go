package middlewares

import (
	"encoding/json"
	"gin_xiaomi_shop/models/admin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// InitAdminAuthMiddleware 进行权限判断，没有登录的用户无法进入后台管理系统
func InitAdminAuthMiddleware(ctx *gin.Context) {
	// 1. 获取访问url地址，包含query参数
	pathname := getPathname(ctx) //  /admin/getCaptcha?t=1692851505664
	// 2. 获取session保存的信息
	// 获取userInfo对应的session
	session := sessions.Default(ctx)
	u := session.Get("userInfo")
	uStr, ok := u.(string)
	if ok {
		var userInfo admin.Manager
		err := json.Unmarshal([]byte(uStr), &userInfo)
		if err != nil && !isPathsWhitelist(pathname) {
			ctx.Redirect(http.StatusFound, "/admin/login")
		}
	} else {
		// 用户没有登录但不再白名单内
		if !isPathsWhitelist(pathname) {
			ctx.Redirect(http.StatusFound, "/admin/login")
		}
	}
}

func getPathname(ctx *gin.Context) (pathname string) {
	pathname = ctx.Request.URL.String()
	if strings.Contains(pathname, "?") {
		pathname = strings.Split(pathname, "?")[0]
	}
	return
}

func isPathsWhitelist(path string) bool {
	var pathsWhitelist = []string{"/admin/getCaptcha", "/admin/doLogin", "/admin/login"}
	for _, s := range pathsWhitelist {
		if s == path {
			return true
		}
	}
	return false
}
