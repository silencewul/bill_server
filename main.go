package main

import (
	"bill/apis"
	"bill/middleware"
	"bill/models"
	"bill/modules/log"
	"bill/modules/setting"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func getAddr() string {

	host := setting.GetConfig().Server.Host
	port := setting.GetConfig().Server.Port

	return host + ":" + port
}

func main() {
	if models.MasterDB == nil {
		log.GetSugar().Error("数据库初始化错误")
		return
	}

	if setting.GetConfig().Server.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(middleware.Ginzap(log.Get(), time.RFC3339, true))

	//r.Use(ginzap.RecoveryWithZap(log.Get(), true))

	//r.Use(middleware.InjectLogger(logger)) //将logger实例注入context,方便调用

	// 允许跨域访问，如需配置更严格的规则，使用cors.New(cors.Config{//初始化配置参数}})
	// 参考文档: https://github.com/gin-contrib/cors
	//corsConfig := cors.DefaultConfig()
	//corsConfig.AllowAllOrigins = true
	//corsConfig.AllowCredentials = true
	//corsConfig.AddAllowHeaders("authorization ")
	//
	//r.Use(cors.New(corsConfig))

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		//AllowOrigins:  []string{"http://localhost:3000", "http://127.0.0.1:3000", "https://dev.gumola.com:2015", "http://dev.gumola.com"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	apis.RegisterRoutes(r.Group("/api", middleware.CtxData))

	// Serve静态文件
	//r.Static("/static", "./static")

	url := ginSwagger.URL(setting.GetAppDomain() + "/api/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//models.ReloadPermissionsForDefaultAdminRoles([]string{models.RoleSuperAdmin})

	_ = r.Run(getAddr())
}
