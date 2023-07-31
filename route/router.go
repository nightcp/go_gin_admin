package route

import (
	"admin/app/admin"
	"admin/core"
	"admin/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	gin.SetMode(core.Config.GinMode)
	r := gin.New()
	r.Static(core.Config.FileRequestPath, "storage")
	r.Use(
		middlewares.CorsHandler(),
		middlewares.LocaleHandler(),
		gin.Logger(),
		middlewares.ErrorHandler,
	)
	admin.Register(r)
	r.GET("/swagger/admin/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("admin")))
	return r
}
