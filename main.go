package main

import (
	"gin_xiaomi_shop/controllers/admin"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/**/*")

	admin.Controller(r)

	err := r.Run()
	if err != nil {
		return
	}
}
