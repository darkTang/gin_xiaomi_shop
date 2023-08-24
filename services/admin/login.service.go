package admin

import (
	"encoding/json"
	"gin_xiaomi_shop/common"
	"gin_xiaomi_shop/database/mysql"
	"gin_xiaomi_shop/models/admin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = mysql.DB

type LoginService struct {
	baseService
}

func (LoginService) Index(ctx *gin.Context) {
	//fmt.Println(common.Md5("123456"))
	ctx.HTML(http.StatusOK, "admin/login/login.html", nil)
}

func (ls LoginService) DoLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	captchaId := ctx.PostForm("captchaId")
	verifyValue := ctx.PostForm("verifyValue")
	if flag := common.VerifyCaptcha(captchaId, verifyValue); flag {
		// 查看数据库的username、password是否存在\
		userInfo := admin.Manager{}
		if err := db.Where("username = ? and password = ?", username, common.Md5(password)).First(&userInfo).Error; err == nil {
			// 保存用户信息，cookie或session
			session := sessions.Default(ctx)
			session.Options(sessions.Options{
				MaxAge:   2 * 3600,
				Domain:   "localhost",
				Path:     "/",
				Secure:   false,
				HttpOnly: false,
				SameSite: 0,
			})
			marshal, err := json.Marshal(userInfo)
			if err != nil {
				return
			}
			session.Set("userInfo", string(marshal))
			err = session.Save()
			if err != nil {
				return
			}
			ls.success(ctx, "登录成功", "/admin")
		} else {
			ls.error(ctx, "用户名或密码错误", "/admin/login")
		}
	} else {
		ls.error(ctx, "验证码验证失败", "/admin/login")
	}
}

func (ls LoginService) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Options(sessions.Options{Path: "/", MaxAge: -1}) // 清除浏览器储存sessionId的cookie
	session.Delete("userInfo")                               // 删除储存在服务器内存的session，此时浏览器的cookie还存在，但是获取到的是null，因此session的值被删除了
	// 坑：session的Set和Delete都会产生cookie，要想使Delete生效，必须将Delete中的sessions.Options的path配置和Set的sessions.Options的path配置一样，这样将会覆盖Set的Cookie，达到删除效果
	// 如果Delete的sessions.Options的path配置和Set的sessions.Options的path配置不一样，就会产生新的cookie
	err := session.Save()
	if err != nil {
		return
	}
	ls.success(ctx, "退出登录成功", "/admin/login")
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
