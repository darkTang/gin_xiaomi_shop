package main

import (
	"gin_xiaomi_shop/controllers/admin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/**/*")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession123", store))

	admin.Controller(r)

	err := r.Run()
	if err != nil {
		return
	}
}
