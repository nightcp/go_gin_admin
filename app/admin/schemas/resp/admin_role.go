package resp

// AdminRoleListResp 获取角色列表
type AdminRoleListResp struct {
	ID        uint   `json:"id"`         // 角色ID
	Name      string `json:"name"`       // 角色名称
	CreatedAt string `json:"created_at"` // 创建时间
}
