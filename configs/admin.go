package configs

import (
	"admin/models"
	"github.com/gin-gonic/gin"
)

var AdminConfig = adminConfig{
	AuthenticateWhiteList: []string{
		"/api/admin/auth/login",
	},
	AdminIDKey:           "admin_id",
	AdminUserKey:         "admin_user",
	AdminUserTokenSetKey: "admin_user_token_set:",
}

type adminConfig struct {
	AuthenticateWhiteList []string
	AdminIDKey            string
	AdminUserKey          string
	AdminUserTokenSetKey  string
}

// GetAdminId 获取用户ID
func (config adminConfig) GetAdminId(c *gin.Context) uint {
	adminId, _ := c.Get(config.AdminIDKey)
	return adminId.(uint)
}

// GetAdminUser 获取用户详情
func (config adminConfig) GetAdminUser(c *gin.Context) models.AdminUser {
	adminUser, _ := c.Get(config.AdminUserKey)
	return adminUser.(models.AdminUser)
}
