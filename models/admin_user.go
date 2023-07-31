package models

import "github.com/golang-module/carbon/v2"

type AdminUser struct {
	ID        uint            `gorm:"type:bigint;not null;primaryKey;autoIncrement;comment:主键ID" json:"id"`
	CreatedAt carbon.DateTime `gorm:"timestamptz;comment:创建时间" json:"created_at"`
	UpdatedAt carbon.DateTime `gorm:"timestamptz;comment:更新时间" json:"updated_at"`
	Username  string          `gorm:"type:varchar(32);not null;default:'';uniqueIndex;comment:用户名" json:"username"`
	Password  string          `gorm:"type:varchar(32);not null;default:'';comment:登录密码" json:"password"`
	Salt      string          `gorm:"type:varchar(8);not null;default:'';comment:加密盐" json:"salt"`
	Nickname  string          `gorm:"type:varchar(32);not null;default:'';comment:昵称" json:"nickname"`
	Avatar    string          `gorm:"type:varchar(255);not null;default:'';comment:头像" json:"avatar"`
	RoleID    uint            `gorm:"type:bigint;not null;default:0;index;comment:角色ID" json:"role_id"`
	Role      AdminRole       `gorm:"foreignKey:RoleID" json:"role"`
}
