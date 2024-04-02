package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"irms-api.com/project-api/pkg/dao"
	"irms-api.com/project-api/pkg/repo"
	"irms-api.com/project-api/pkg/service"
	"irms-api.com/project-api/pkg/service/login"
	project_common "irms-api.com/project-common"
)

type HandleLogin struct {
	cache repo.Cache
}

func New() *HandleLogin {
	return &HandleLogin{
		cache: dao.Rc,
	}
}

// @Summary		获取验证码
// @title		获取title
// @version		1.0
// @Tags		Login
// @description	获取验证码接口
// @Produce		json
// @Router		/project/login/getCaptcha [get]
func (hl *HandleLogin) GetCaptcha(ctx *gin.Context) {
	result := &project_common.Result{}
	Captcha := &service.GetCap{}
	data := Captcha.GetCaptcha()
	//go func() {
	//	zap.L().Info("验证码生成成功")
	//	c, concel := context.WithTimeout(context.Background(), 2*time.Second)
	//	defer concel()
	//	err := hl.cache.Put(c, "验证码Id", "dd", 15*time.Second)
	//	if err != nil {
	//		log.Println("验证码存入redis错误,cause by:", err)
	//	}
	//}()
	ctx.JSON(200, result.Success(data))
}

// @Summary 登录
// @title
// @Tags Login
// @description 登录接口
// @Produce json
// @Router /project/login/login
func (hl *HandleLogin) Login(ctx *gin.Context) {
	result := project_common.Result{}
	username := ctx.Query("username")
	password := ctx.Query("password")
	captchaId := ctx.Query("captchaId")
	reCode := ctx.Query("reCode")
	fmt.Println(username, password, captchaId, reCode)
	resu := login.LoginService(username, password, captchaId, reCode)
	if resu != true {
		ctx.JSON(200, result.Fail(203, "验证码错误"))
	}
	//fmt.Println(result)
	//ctx.JSON(200, result.Success(resu))
}
