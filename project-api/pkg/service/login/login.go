package login

import (
	dao "irms-api.com/project-api/pkg/dao/user"
	"irms-api.com/project-api/pkg/service"
)

type Result struct {
	Status bool
	Msg    string
}

func LoginService(username, password, captchaId, reCode string) (bool, string) {
	Verify := &service.GetCap{}
	result := Verify.VerifyCaptcha(captchaId, reCode)
	if result != true {
		return result, "验证码错误"
	} else {
		user := dao.UserDao{}
		result := user.VerifyUser(username, password)
		return result, "登录成功"
	}
}
