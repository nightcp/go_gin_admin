package resp

// AdminUserListResp 获取用户列表
type AdminUserListResp struct {
	ID        uint              `json:"id"`         // 用户ID
	Username  string            `json:"username"`   // 用户名
	Nickname  string            `json:"nickname"`   // 昵称
	Avatar    string            `json:"avatar"`     // 头像
	CreatedAt string            `json:"created_at"` // 创建时间
	Role      AdminUserRoleResp `json:"role"`       // 角色
}

// AdminUserRoleResp 获取用户角色
type AdminUserRoleResp struct {
	ID   uint   `json:"id"`   // 角色ID
	Name string `json:"name"` // 角色名称
}
