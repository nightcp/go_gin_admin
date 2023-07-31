package migration

import (
	"admin/models"
	"gorm.io/gorm"
)

func DBMigrate(DB *gorm.DB) {
	_ = DB.AutoMigrate(&models.AdminUser{})
	_ = DB.AutoMigrate(&models.AdminRole{})
	DB.Exec("comment on table admin_users is '管理员用户表'")
	DB.Exec("comment on table admin_roles is '管理员角色表'")
}
