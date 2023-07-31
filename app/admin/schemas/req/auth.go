package req

import (
	"admin/core/request"
)

// LoginReq 登录参数
type LoginReq struct {
	Username string `json:"username" binding:"required,max=32"` // 用户名
	Password string `json:"password" binding:"required,max=32"` // 密码
}

// GetMessage 登录验证消息
func (adminLoginReq LoginReq) GetMessage() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Username.required": "UsernameIsRequired",
		"Username.max": request.ValidateMsg{
			Text: "UsernameLongerThanMax",
			Args: []interface{}{32},
		},
		"Password.required": "PasswordIsRequired",
		"password.max": request.ValidateMsg{
			Text: "PasswordLongerThanMax",
			Args: []interface{}{32},
		},
	}
}

// EditProfileReq 修改用户信息参数
type EditProfileReq struct {
	Nickname string `json:"nickname" binding:"required,max=15"`               // 昵称
	Avatar   string `json:"avatar" binding:"required,max=150,file_is_exists"` // 头像
}

// GetMessage 修改用户信息验证消息
func (editProfileReq EditProfileReq) GetMessage() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Nickname.required": "NicknameIsRequired",
		"Nickname.max": request.ValidateMsg{
			Text: "NicknameLongerThanMax",
			Args: []interface{}{15},
		},
		"Avatar.required": "AvatarIsRequired",
		"Avatar.max": request.ValidateMsg{
			Text: "AvatarLongerThanMax",
			Args: []interface{}{150},
		},
		"Avatar.file_is_exists": "FileDoesNotExist",
	}
}

// EditPasswordReq 修改密码参数
type EditPasswordReq struct {
	OldPassword     string `json:"old_password" binding:"required,max=32"`                         // 旧密码
	NewPassword     string `json:"new_password" binding:"required,password"`                       // 新密码
	ConfirmPassword string `json:"confirm_password" binding:"required,max=32,eqfield=NewPassword"` // 确认密码
}

// GetMessage 修改密码验证消息
func (editPasswordReq EditPasswordReq) GetMessage() request.ValidatorMessages {
	return request.ValidatorMessages{
		"OldPassword.required": "OldPasswordIsRequired",
		"OldPassword.max": request.ValidateMsg{
			Text: "OldPasswordLongerThanMax",
			Args: []interface{}{32},
		},
		"NewPassword.required":     "NewPasswordIsRequired",
		"NewPassword.password":     "NewPasswordInvalid",
		"ConfirmPassword.required": "ConfirmPasswordIsRequired",
		"ConfirmPassword.max": request.ValidateMsg{
			Text: "ConfirmPasswordLongerThanMax",
			Args: []interface{}{32},
		},
		"ConfirmPassword.eqfield": "ConfirmPasswordEqNewPassword",
	}
}
