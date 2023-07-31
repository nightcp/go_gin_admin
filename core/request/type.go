package request

type ValidateMsg struct {
	Text string
	Args []interface{}
}

// PageReq 分页参数
type PageReq struct {
	PageNo   int `form:"page_no" binding:"required,number,min=1"`             // 页码
	PageSize int `form:"page_size" binding:"required,number,min=10,max=1000"` // 每页记录数
}

// GetMessage 分页验证消息
func (pageReq PageReq) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"PageNo.required": "PageNoIsRequired",
		"PageNo.number":   "PageNoIsNumber",
		"PageNo.min": ValidateMsg{
			Text: "PageNoMinimum",
			Args: []interface{}{1},
		},
		"PageSize.required": "PageSizeIsRequired",
		"PageSize.number":   "PageSizeIsNumber",
		"PageSize.min": ValidateMsg{
			Text: "PageSizeMinimum",
			Args: []interface{}{10, 1000},
		},
		"PageSize.max": ValidateMsg{
			Text: "PageSizeMaximum",
			Args: []interface{}{10, 1000},
		},
	}
}
