package common

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 创建store  默认保存在服务器内存中
var store = base64Captcha.DefaultMemStore

// GenerateCaptcha 创建验证码
func GenerateCaptcha() (id string, b64s string, err error) {
	var driver base64Captcha.Driver
	DriverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4, // 干扰的线 2 ｜ 4
		Length:          4,
		Source:          "123456789qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driver = DriverString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = c.Generate()
	return
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id string, verifyValue string) bool {
	if store.Verify(id, verifyValue, true) {
		return true
	} else {
		return false
	}
}
