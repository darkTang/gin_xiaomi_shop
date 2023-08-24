package admin

import (
	"gin_xiaomi_shop/middlewares"
	"gin_xiaomi_shop/services/admin"
	"github.com/gin-gonic/gin"
)

func Controller(r *gin.Engine) {
	adminRouter := r.Group("/admin", middlewares.InitAdminAuthMiddleware)

	var loginService admin.LoginService
	var managerService admin.ManagerService
	var focusService admin.FocusService
	var mainService admin.MainService
	{
		adminRouter.GET("", mainService.Index)
		adminRouter.GET("/welcome", mainService.Welcome)

		adminRouter.GET("/login", loginService.Index)
		adminRouter.POST("/doLogin", loginService.DoLogin)
		adminRouter.GET("/logout", loginService.Logout)

		adminRouter.GET("/getCaptcha", loginService.GetCaptcha)

		adminRouter.GET("/manager", managerService.Index)
		adminRouter.GET("/manager/add", managerService.Add)
		adminRouter.GET("/manager/edit", managerService.Edit)
		adminRouter.GET("/manager/delete", managerService.Delete)

		adminRouter.GET("/focus", focusService.Index)
		adminRouter.GET("/focus/add", focusService.Add)
		adminRouter.GET("/focus/edit", focusService.Edit)
		adminRouter.GET("/focus/delete", focusService.Delete)

	}
}
