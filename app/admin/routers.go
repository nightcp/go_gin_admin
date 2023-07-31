package admin

import (
	"admin/app/admin/handlers"
	"admin/middlewares"
	"github.com/gin-gonic/gin"
)

// @title           Admin API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @basePath /api/admin
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func Register(gin *gin.Engine) {
	group := gin.Group("api/admin")
	group.Use(middlewares.AuthenticateHandler())
	group.POST("auth/login", handlers.AuthHandler.Login)
	group.GET("auth/profile", handlers.AuthHandler.Profile)
	group.PUT("auth/profile", handlers.AuthHandler.EditProfile)
	group.PUT("auth/password", handlers.AuthHandler.EditPassword)
	group.GET("admin-roles", handlers.AdminRoleHandle.List)
	group.POST("admin-roles", handlers.AdminRoleHandle.Add)
	group.PUT("admin-roles/:id", handlers.AdminRoleHandle.Edit)
	group.DELETE("admin-roles/:id", handlers.AdminRoleHandle.Delete)
	group.GET("admin-users", handlers.AdminUserHandler.List)
	group.POST("admin-users", handlers.AdminUserHandler.Add)
	group.PUT("admin-users/:id", handlers.AdminUserHandler.Edit)
	group.DELETE("admin-users/:id", handlers.AdminUserHandler.Delete)
	group.POST("publish/upload", handlers.PublishHandler.UploadFile)
}
