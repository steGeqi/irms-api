package login

import (
	"irms-api.com/project-api/pkg/service"
)

type Result struct {
	Status bool
	Msg    string
}

func LoginService(username, password, captchaId, reCode string) bool {
	Verify := &service.GetCap{}
	result := Verify.VerifyCaptcha(captchaId, reCode)
	if result == false {
		return result
	} else {
		
		return true
	}
}
