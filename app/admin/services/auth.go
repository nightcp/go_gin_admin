package services

import (
	"admin/app/admin/schemas/req"
	"admin/app/admin/schemas/resp"
	"admin/configs"
	"admin/core"
	"admin/core/auth"
	"admin/core/redis"
	"admin/core/response"
	"admin/models"
	"admin/util"
	"github.com/gin-gonic/gin"
)

type authService struct {
}

var AuthService = &authService{}

// Login 用户登录
func (service *authService) Login(c *gin.Context, loginReq req.LoginReq) resp.LoginResp {
	adminUser := models.AdminUser{}
	core.DB.Select("id", "username", "password", "salt").Where("username = ?", loginReq.Username).
		Find(&adminUser).Limit(1)
	if adminUser.ID == 0 {
		response.Fail(c, "UsernameOrPasswordIsIncorrect", nil)
	}
	if pass := util.StrUtil.CheckPassword(loginReq.Password, adminUser.Salt, adminUser.Password); !pass {
		response.Fail(c, "UsernameOrPasswordIsIncorrect", nil)
	}
	ttl := core.Config.JwtTTl
	token, err := auth.Jwt.MakeToken(auth.CustomClaims{
		UserID:   adminUser.ID,
		Username: adminUser.Username,
		Identity: "admin",
	}, ttl)
	if err != nil {
		response.Fail(c, "LoginFailed", nil)
	}
	adminUserTokenSetKey := configs.AdminConfig.AdminUserTokenSetKey + adminUser.Username
	var ssMembers []redis.SSetMember
	members := redis.RDBHelper.SSGetMembersByScore(adminUserTokenSetKey, "0", "1", 0, 0)
	for _, member := range members {
		ssMembers = append(ssMembers, redis.SSetMember{
			Score:  1,
			Member: member,
		})
	}
	ssMembers = append(ssMembers, redis.SSetMember{
		Score:  0,
		Member: token,
	})
	if redis.RDBHelper.SSAdd(adminUserTokenSetKey, ssMembers) == false {
		response.Fail(c, "LoginFailed", nil)
	}
	return resp.LoginResp{
		Token: "Bearer " + token,
	}
}

// GetProfile 获取用户信息
func (service *authService) GetProfile(c *gin.Context) resp.ProfileResp {
	adminUser := configs.AdminConfig.GetAdminUser(c)
	return resp.ProfileResp{
		Username: adminUser.Username,
		Nickname: adminUser.Nickname,
		Avatar:   adminUser.Avatar,
	}
}

// UpdateProfile 更新用户信息
func (service *authService) UpdateProfile(c *gin.Context, updateReq req.EditProfileReq) {
	adminUser := configs.AdminConfig.GetAdminUser(c)
	err := core.DB.Model(&adminUser).Updates(models.AdminUser{
		Nickname: updateReq.Nickname,
		Avatar:   updateReq.Avatar,
	}).Error
	if err != nil {
		response.Fail(c, "EditFailed", nil)
	}
}

// UpdatePassword 更新用户密码
func (service *authService) UpdatePassword(c *gin.Context, updateReq req.EditPasswordReq) {
	adminUser := configs.AdminConfig.GetAdminUser(c)
	if util.StrUtil.CheckPassword(updateReq.OldPassword, adminUser.Salt, adminUser.Password) == false {
		response.Fail(c, "OldPasswordInvalid", nil)
	}
	salt := util.StrUtil.MakeRandomStr(8)
	password := util.StrUtil.MakePassword(updateReq.NewPassword, salt)
	err := core.DB.Model(&adminUser).Updates(models.AdminUser{
		Password: password,
		Salt:     salt,
	}).Error
	if err != nil {
		response.Fail(c, "EditFailed", nil)
	}
}
