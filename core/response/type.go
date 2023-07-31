package response

const (
	HttpStatusOk      = 200 // 成功
	HttpUnauthorized  = 401 // 未登录
	HttpAbnormalLogin = 402 // 异常登录
	HttpServerError   = 500 // 系统错误
)

// Resp 统一返回格式
type Resp struct {
	Code   int         `json:"code"`   // 状态码
	Status bool        `json:"status"` // 状态
	Msg    string      `json:"msg"`    // 消息
	Data   interface{} `json:"data"`   // 数据
}

// Options 返回选项
type Options struct {
	HttpCode int           // HTTP状态码
	MsgArgs  []interface{} // 消息参数
}

// PageResp 分页返回格式
type PageResp struct {
	Count    int64       `json:"count"`     // 总数
	PageNo   int         `json:"page_no"`   // 当前页
	PageSize int         `json:"page_size"` // 每页记录数
	List     interface{} `json:"list"`      // 列表
}
