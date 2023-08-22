package admin

import (
	"gin_xiaomi_shop/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginService struct {
}

func (LoginService) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/login/login.html", nil)
}

func (LoginService) GetCaptcha(ctx *gin.Context) {
	id, b64s, err := common.GenerateCaptcha()
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

func (LoginService) VerifyCaptcha(ctx *gin.Context) {

}
