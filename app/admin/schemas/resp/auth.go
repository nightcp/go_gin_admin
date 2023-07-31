package resp

// LoginResp 登录
type LoginResp struct {
	Token string `json:"token"` // 访问令牌
}

// ProfileResp 获取用户详情
type ProfileResp struct {
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
}
