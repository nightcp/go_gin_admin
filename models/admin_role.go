package models

import "github.com/golang-module/carbon/v2"

type AdminRole struct {
	ID        uint            `gorm:"type:bigint;not null;primaryKey;autoIncrement;comment:主键ID" json:"id"`
	CreatedAt carbon.DateTime `gorm:"timestamptz;comment:创建时间" json:"created_at"`
	UpdatedAt carbon.DateTime `gorm:"timestamptz;comment:更新时间" json:"updated_at"`
	Name      string          `gorm:"type:varchar(32);not null;default:'';uniqueIndex;comment:角色名称" json:"name"`
}
