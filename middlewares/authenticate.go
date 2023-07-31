package middlewares

import (
	"admin/configs"
	"admin/core"
	"admin/core/auth"
	"admin/core/redis"
	"admin/core/response"
	"admin/models"
	"admin/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
	"strings"
)

// AuthenticateHandler 处理认证
func AuthenticateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if r := util.ArrUtil.InArray(c.FullPath(), configs.AdminConfig.AuthenticateWhiteList); !r {
			header := c.Request.Header.Get("Authorization")
			if header == "" {
				response.Fail(c, "Unauthorized", nil, response.Options{HttpCode: 401})
			}
			// 按空格分割
			parts := strings.SplitN(header, " ", 2)
			if (len(parts) == 2 && parts[0] == "Bearer") == false {
				response.Fail(c, "Unauthorized", nil, response.Options{HttpCode: 401})
			}
			// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
			claims, err := auth.Jwt.ParseToken(parts[1])
			if err != nil {
				response.Fail(c, "Unauthorized", nil, response.Options{HttpCode: 401})
			}
			if carbon.Now().Timestamp() > claims.ExpireAt {
				response.Fail(c, "Unauthorized", nil, response.Options{HttpCode: 401})
			}
			switch claims.Identity {
			case "admin":
				adminUser := models.AdminUser{}
				core.DB.Select([]string{
					"id",
					"username",
					"nickname",
					"avatar",
					"password",
					"salt",
				}).Where("id = ?", claims.UserID).First(&adminUser)
				if adminUser.ID == 0 {
					response.Fail(c, "UserNotExists", nil)
				}
				adminUserTokenSetKey := configs.AdminConfig.AdminUserTokenSetKey + adminUser.Username
				score := redis.RDBHelper.SSGetScore(adminUserTokenSetKey, parts[1])
				// 如果token不在集合中, 则未登录
				if score == -1 {
					response.Fail(c, "Unauthorized", nil, response.Options{HttpCode: 401})
				}
				// 如果token在集合中的分数超过0分, 则已重新登录, 使用旧token访问提示异地登录, 并从集合中移除旧token
				if score > 0 {
					if redis.RDBHelper.SSDel(adminUserTokenSetKey, []string{parts[1]}) == false {
						response.Fail(c, "SystemError", response.Options{HttpCode: 500})
					}
					response.Fail(c, "RemoteLogin", nil)
				}
				c.Set(configs.AdminConfig.AdminIDKey, claims.UserID)
				c.Set(configs.AdminConfig.AdminUserKey, adminUser)
				if carbon.Now().DiffInMinutes(carbon.CreateFromTimestamp(claims.ExpireAt)) <= int64(core.Config.JwtRenewTTl) {
					newToken, tokenErr := auth.Jwt.MakeToken(auth.CustomClaims{
						UserID:   adminUser.ID,
						Username: adminUser.Username,
						Identity: "admin",
					}, core.Config.JwtTTl)
					if tokenErr != nil {
						core.Logger.Error("Remake token", zap.Any("error", tokenErr))
						response.Fail(c, "SystemError", nil, response.Options{HttpCode: 500})
					}
					// 集合中移除旧token
					if redis.RDBHelper.SSDel(adminUserTokenSetKey, []string{parts[1]}) == false {
						response.Fail(c, "SystemError", response.Options{HttpCode: 500})
					}
					// 集合中加入新token
					var ssMembers []redis.SSetMember
					ssMembers = append(ssMembers, redis.SSetMember{
						Score:  0,
						Member: newToken,
					})
					if redis.RDBHelper.SSAdd(adminUserTokenSetKey, ssMembers) == false {
						response.Fail(c, "SystemError", response.Options{HttpCode: 500})
					}
					c.Header("Authorization", "Bearer "+newToken)
				}
			default:
				response.Fail(c, "Unauthorized", nil)
			}
		}
		c.Next()
	}
}
