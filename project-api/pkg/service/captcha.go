package service

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

type GetCap struct {
	Id     string
	Base64 string
}

func (g *GetCap) GetCaptcha() (res interface{}) {
	// 生成默认数字
	//driver := base64Captcha.DefaultDriverDigit
	driver := base64Captcha.NewDriverDigit(180, 300, 5, 0.7, 80)

	// 生成base64图片
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := c.Generate()
	result := GetCap{Id: id, Base64: b64s}
	if err != nil {
		fmt.Println("Register GetCaptchaPhoto get base64Captcha has err:", err)
		return
	}
	return result

}

func (g *GetCap) VerifyCaptcha(captchaId, reCode string) bool {
	return store.Verify(captchaId, reCode, true)
}
