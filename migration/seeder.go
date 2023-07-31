package migration

import (
	"admin/models"
	"gorm.io/gorm"
)

func DBSeed(DB *gorm.DB) {
	adminUserSeeder(DB)
	adminRoleSeeder(DB)
}

// adminUserSeeder 填充管理员用户数据
func adminUserSeeder(DB *gorm.DB) {
	username := "admin"
	adminUser := models.AdminUser{}
	DB.Where("username = ?", username).Find(&adminUser).Limit(1)
	if adminUser.ID == 0 {
		adminUser.Username = username
		adminUser.Nickname = "管理员"
		adminUser.RoleID = 1
		DB.Create(&adminUser)
	}
}

// adminRoleSeeder 填充管理员角色数据
func adminRoleSeeder(DB *gorm.DB) {
	name := "超级管理员"
	adminRole := models.AdminRole{}
	DB.Where("name = ?", name).Find(&adminRole).Limit(1)
	if adminRole.ID == 0 {
		adminRole.Name = name
		DB.Create(&adminRole)
	}
}
