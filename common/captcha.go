package common

import (
	"gin_xiaomi_shop/database/redis"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 创建store  默认保存在当前服务器内存中，若要部署分布式，需要将验证码信息保存在redis中
//var store = base64Captcha.DefaultMemStore

// 自定义store，实现redis存储
var store base64Captcha.Store = redis.Store{}

// GenerateCaptcha 创建验证码
func GenerateCaptcha() (id string, b64s string, err error) {
	var driver base64Captcha.Driver
	DriverString := base64Captcha.DriverString{ // 配置信息可视化见 https://captcha.mojotv.cn
		Height:          40,
		Width:           100,
		NoiseCount:      0,     // 干扰数量
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
	/*
		1. id 生成验证码的唯一id，用于后续检验验证码是否通过
		2. b64s base64格式的图片，返回给前端
		3. err 错误信息
		4. 通过store会将id保存在服务器内存中
	*/
}

/*
VerifyCaptcha 验证验证码
1. id 前端传递过来的id，用于检验验证码是否通过
2. verifyValue 前端传过来的值
3. 首先前端传递过来的id会在store寻找是否存在相同的id，然后通过比较传过来的值和在服务器存储的值是否一致，一致则检验通过
*/
func VerifyCaptcha(id string, verifyValue string) bool {
	if store.Verify(id, verifyValue, true) { // true表示验证成功就会清空store存储的数据
		return true
	} else {
		return false
	}
}
