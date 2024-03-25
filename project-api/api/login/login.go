package login

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"irms-api.com/project-api/pkg/dao"
	"irms-api.com/project-api/pkg/repo"
	project_common "irms-api.com/project-common"
	"log"
	"time"
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
// @title			获取title
// @version		1.0
// @Tags			Login
// @description	获取验证码接口
// @Produce		json
// @Router			/project/login/getCaptcha [get]
func (hl *HandleLogin) GetCaptcha(ctx *gin.Context) {
	result := &project_common.Result{}
	code := "123456"
	go func() {
		zap.L().Info("验证码生成成功")
		c, concel := context.WithTimeout(context.Background(), 2*time.Second)
		defer concel()
		err := hl.cache.Put(c, "验证码", code, 15*time.Second)
		if err != nil {
			log.Println("验证码存入redis错误,cause by:", err)
		}
	}()
	ctx.JSON(200, result.Success(code))
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
	fmt.Println(username, password)
	ctx.JSON(200, result.Success(username))
}
