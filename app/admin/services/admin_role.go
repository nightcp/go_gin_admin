package services

import (
	"admin/app/admin/schemas/req"
	"admin/app/admin/schemas/resp"
	"admin/core"
	"admin/core/request"
	"admin/core/response"
	"admin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type adminRoleService struct {
}

var AdminRoleService = &adminRoleService{}

// GetRoleList 获取角色列表
func (service *adminRoleService) GetRoleList(pageReq request.PageReq) response.PageResp {
	var (
		count      int64
		adminRoles []models.AdminRole
		listResp   []resp.AdminRoleListResp
	)
	limit, offset := core.DBHelper.Page(pageReq.PageNo, pageReq.PageSize)
	tx := core.DB.Model(&models.AdminRole{}).Session(&gorm.Session{})
	tx.Count(&count)
	tx.Select([]string{"id", "name", "created_at"}).Offset(offset).Limit(limit).Order("id desc").Find(&adminRoles)
	for _, v := range adminRoles {
		listResp = append(listResp, resp.AdminRoleListResp{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: v.CreatedAt.ToDateTimeString(),
		})
	}
	return response.PageResp{
		Count:    count,
		PageNo:   pageReq.PageNo,
		PageSize: pageReq.PageSize,
		List:     listResp,
	}
}

// StoreRole 保存角色
func (service *adminRoleService) StoreRole(c *gin.Context, addReq req.AddAdminRoleReq) {
	var count int64
	core.DB.Model(&models.AdminRole{}).Where("name = ?", addReq.RoleName).Count(&count)
	if count > 0 {
		response.Fail(c, "RoleNameIsExists", nil)
	}
	err := core.DB.Create(&models.AdminRole{Name: addReq.RoleName}).Error
	if err != nil {
		response.Fail(c, "AddFailed", nil)
	}
}

// UpdateRole 更新角色
func (service *adminRoleService) UpdateRole(c *gin.Context, editReq req.EditAdminRoleReq) {
	var count int64
	id := c.Param("id")
	adminRole := models.AdminRole{}
	core.DB.Select("id").Where("id = ?", id).Find(&adminRole).Limit(1)
	if adminRole.ID == 0 {
		response.Fail(c, "RoleNotExists", nil)
	}
	core.DB.Model(&models.AdminRole{}).Where("id <> ? AND name = ?", adminRole.ID, editReq.RoleName).Count(&count)
	if count > 0 {
		response.Fail(c, "RoleNameIsExists", nil)
	}
	err := core.DB.Model(&adminRole).Update("name", editReq.RoleName).Error
	if err != nil {
		response.Fail(c, "EditFailed", nil)
	}
}

// DestroyRole 删除角色
func (service *adminRoleService) DestroyRole(c *gin.Context) {
	id := c.Param("id")
	adminRole := models.AdminRole{}
	core.DB.Select("id").Where("id = ?", id).Find(&adminRole).Limit(1)
	if adminRole.ID == 0 {
		response.Fail(c, "RoleNotExists", nil)
	}
	txErr := core.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.AdminUser{}).Where("role_id = ?", adminRole.ID).
			Updates(map[string]interface{}{"role_id": 0}).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&adminRole).Error
		if err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		response.Fail(c, "DeleteFailed", nil)
	}
}
