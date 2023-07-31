package services

import (
	"admin/app/admin/schemas/req"
	"admin/app/admin/schemas/resp"
	"admin/core"
	"admin/core/request"
	"admin/core/response"
	"admin/models"
	"admin/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type adminUserService struct {
}

var AdminUserService = &adminUserService{}

// GetUserList 获取用户列表
func (service *adminUserService) GetUserList(pageReq request.PageReq) response.PageResp {
	var (
		count      int64
		adminUsers []models.AdminUser
		listResp   []resp.AdminUserListResp
	)
	limit, offset := core.DBHelper.Page(pageReq.PageNo, pageReq.PageSize)
	tx := core.DB.Model(&models.AdminUser{}).Session(&gorm.Session{})
	tx.Count(&count)
	tx.Select([]string{
		"admin_users.id",
		"admin_users.username",
		"admin_users.nickname",
		"admin_users.avatar",
		"admin_users.role_id",
		"admin_users.created_at",
	}).Joins("Role").Offset(offset).Limit(limit).Find(&adminUsers)
	for _, v := range adminUsers {
		listResp = append(listResp, resp.AdminUserListResp{
			ID:        v.ID,
			Username:  v.Username,
			Nickname:  v.Nickname,
			Avatar:    v.Avatar,
			CreatedAt: v.CreatedAt.ToDateTimeString(),
			Role: resp.AdminUserRoleResp{
				ID:   v.Role.ID,
				Name: v.Role.Name,
			},
		})
	}
	return response.PageResp{
		Count:    count,
		PageNo:   pageReq.PageNo,
		PageSize: pageReq.PageSize,
		List:     listResp,
	}
}

// StoreUser 保存用户
func (service *adminUserService) StoreUser(c *gin.Context, addReq req.AddAdminUserReq) {
	adminUser := models.AdminUser{}
	adminRole := models.AdminRole{}
	core.DB.Select("id").Where("username = ?", addReq.Username).Find(&adminUser).Limit(1)
	if adminUser.ID > 0 {
		response.Fail(c, "UserIsExists", nil)
	}
	core.DB.Select("id").Where("id = ?", addReq.RoleID).Find(&adminRole).Limit(1)
	if adminRole.ID == 0 {
		response.Fail(c, "RoleIsNotExists", nil)
	}
	salt := util.StrUtil.MakeRandomStr(8)
	password := util.StrUtil.MakePassword(addReq.Password, salt)
	adminUser.Username = addReq.Username
	adminUser.Nickname = addReq.Nickname
	adminUser.Avatar = addReq.Avatar
	adminUser.RoleID = addReq.RoleID
	adminUser.Password = password
	adminUser.Salt = salt
	err := core.DB.Create(&adminUser).Error
	if err != nil {
		response.Fail(c, "AddFailed", nil)
	}
}

// UpdateUser 更新用户
func (service *adminUserService) UpdateUser(c *gin.Context, editReq req.EditAdminUserReq) {
	id := c.Param("id")
	adminUser := models.AdminUser{}
	adminRole := models.AdminRole{}
	core.DB.Select([]string{"id"}).Where("id = ?", id).Find(&adminUser).Limit(1)
	if adminUser.ID == 0 {
		response.Fail(c, "UserNotExists", nil)
	}
	core.DB.Select("id").Where("id = ?", editReq.RoleID).Find(&adminRole).Limit(1)
	if adminRole.ID == 0 {
		response.Fail(c, "RoleNotExists", nil)
	}
	updates := map[string]interface{}{
		"nickname": editReq.Nickname,
		"avatar":   editReq.Avatar,
		"role_id":  editReq.RoleID,
	}
	if editReq.NewPassword != "" {
		salt := util.StrUtil.MakeRandomStr(8)
		password := util.StrUtil.MakePassword(editReq.NewPassword, salt)
		updates["salt"] = salt
		updates["password"] = password
	}
	err := core.DB.Model(&adminUser).Updates(updates).Error
	if err != nil {
		response.Fail(c, "EditFialed", nil)
	}
}

// DestroyUser 删除用户
func (service *adminUserService) DestroyUser(c *gin.Context) {
	id := c.Param("id")
	adminUser := models.AdminUser{}
	core.DB.Select([]string{"id", "username"}).Where("id = ?", id).Find(&adminUser).Limit(1)
	if adminUser.ID == 0 {
		response.Fail(c, "UserNotExists", nil)
	}
	if adminUser.Username == "admin" {
		response.Fail(c, "DeleteFailed", nil)
	}
	err := core.DB.Delete(&adminUser).Error
	if err != nil {
		response.Fail(c, "DeleteFailed", nil)
	}
}
