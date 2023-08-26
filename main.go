package main

import (
	"gin_xiaomi_shop/common"
	"gin_xiaomi_shop/controllers/admin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.SetFuncMap(template.FuncMap{ // 必需要在加载模板之前注入
		"UnixToTime": common.UnixToTime,
	})
	r.LoadHTMLGlob("templates/**/**/*")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession123", store))

	admin.Controller(r)

	err := r.Run()
	if err != nil {
		return
	}
}
