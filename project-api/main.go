package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"irms-api.com/project-api/config"
	_ "irms-api.com/project-api/docs"
	"irms-api.com/project-api/middlewares"
	"irms-api.com/project-api/router"
	srv "irms-api.com/project-common"
)

// @title			接口文档
// @version		1.0
// @description	信息资源管理系统后端
// @contact.name	戈亓
func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	config.AppConf.InitZaplog()
	// 初始化路由
	router.InitRouter(r)
	// 调用Grpc
	grpc := router.RegisterGrpc()
	stop := func() {
		grpc.Stop()
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	srv.Run(r, config.AppConf.SC.Name, config.AppConf.SC.Addr, stop)
}
