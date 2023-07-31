package req

import "admin/core/request"

// AddAdminUserReq 新增用户参数
type AddAdminUserReq struct {
	Username string `json:"username" binding:"required,max=32"`      // 用户名
	Password string `json:"password" binding:"required,password"`    // 密码
	Nickname string `json:"nickname" binding:"required,max=32"`      // 昵称
	RoleID   uint   `json:"role_id" binding:"required"`              // 角色ID
	Avatar   string `json:"avatar" binding:"max=150,file_is_exists"` // 头像
}

// GetMessage 新增用户验证消息
func (addAdminUserReq AddAdminUserReq) GetMessage() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Username.required": "UsernameIsRequired",
		"Username.max": request.ValidateMsg{
			Text: "UsernameLongerThanMax",
			Args: []interface{}{32},
		},
		"Password.required": "PasswordIsRequired",
		"Password.password": "PasswordInvalid",
		"Nickname.required": "NicknameIsRequired",
		"Nickname.max": request.ValidateMsg{
			Text: "NicknameLongerThanMax",
			Args: []interface{}{32},
		},
		"RoleID.required": "RoleIDIsRequired",
		"Avatar.max": request.ValidateMsg{
			Text: "AvatarLongerThanMax",
			Args: []interface{}{150},
		},
		"Avatar.file_is_exists": "FileDoesNotExist",
	}
}

// EditAdminUserReq 编辑用户参数
type EditAdminUserReq struct {
	Nickname        string `json:"nickname" binding:"required,max=32"`                    // 昵称
	Avatar          string `json:"avatar" binding:"required,max=150,file_is_exists"`      // 头像
	RoleID          uint   `json:"role_id" binding:"required"`                            // 角色ID
	NewPassword     string `json:"new_password" binding:"password"`                       // 新密码
	ConfirmPassword string `json:"confirm_password" binding:"max=32,eqfield=NewPassword"` // 确认密码
}

// GetMessage 编辑用户验证信息
func (editAdminUserReq EditAdminUserReq) GetMessage() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Nickname.required": "NicknameIsRequired",
		"Nickname.max": request.ValidateMsg{
			Text: "NicknameLongerThanMax",
			Args: []interface{}{32},
		},
		"Avatar.required": "AvatarIsRequired",
		"Avatar.max": request.ValidateMsg{
			Text: "AvatarLongerThanMax",
			Args: []interface{}{150},
		},
		"Avatar.file_is_exists": "FileDoesNotExist",
		"RoleID.required":       "RoleIDIsRequired",
		"NewPassword.password":  "NewPasswordInvalid",
		"ConfirmPassword.max": request.ValidateMsg{
			Text: "ConfirmPasswordLongerThanMax",
			Args: []interface{}{32},
		},
		"ConfirmPassword.eqfield": "ConfirmPasswordEqNewPassword",
	}
}
