package req

import "admin/core/request"

// AddAdminRoleReq 新增角色参数
type AddAdminRoleReq struct {
	RoleName string `json:"role_name" binding:"required,max=32"` // 角色名称
}

// GetMessage 新增角色验证消息
func (addRoleReq AddAdminRoleReq) GetMessage() request.ValidatorMessages {
	return request.ValidatorMessages{
		"RoleName.required": "RoleNameIsRequired",
		"RoleName.max": request.ValidateMsg{
			Text: "RoleNameLongerThanMax",
			Args: []interface{}{32},
		},
	}
}

// EditAdminRoleReq 编辑角色参数
type EditAdminRoleReq struct {
	RoleName string `json:"role_name" binding:"required,max=32"` // 角色名称
}

// GetMessage 编辑角色验证消息
func (editRoleReq EditAdminRoleReq) GetMessage() request.ValidatorMessages {
	return request.ValidatorMessages{
		"RoleName.required": "RoleNameIsRequired",
		"RoleName.max": request.ValidateMsg{
			Text: "RoleNameLongerThanMax",
			Args: []interface{}{32},
		},
	}
}
