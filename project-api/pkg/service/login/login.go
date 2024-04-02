package login

import (
	"irms-api.com/project-api/pkg/service"
)

type Result struct {
	Status bool
	Msg    string
}

func LoginService(username, password, captchaId, reCode string) bool {
	//r := Result{}
	Verify := &service.GetCap{}
	result := Verify.VerifyCaptcha(captchaId, reCode)
	if result != true {
		//r.Status = false
		//r.Msg = "验证码错误"
		return result
	}
	return true
}
