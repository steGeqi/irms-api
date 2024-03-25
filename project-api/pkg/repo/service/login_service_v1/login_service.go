package login_service_v1

import (
	"context"
	"go.uber.org/zap"
	"irms-api.com/project-api/pkg/dao"
	"irms-api.com/project-api/pkg/repo"
	"log"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (hl *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	//result := &project_common.Result{}
	mobile := msg.Mobile
	// 手机号校验
	code := "123456"
	go func() {
		zap.L().Info("验证码生成成功")
		c, concel := context.WithTimeout(context.Background(), 2*time.Second)
		defer concel()
		err := hl.cache.Put(c, "验证码"+mobile, code, 15*time.Second)
		if err != nil {
			log.Println("验证码存入redis错误,cause by:", err)
		}
	}()
	return &CaptchaResponse{}, nil
}
