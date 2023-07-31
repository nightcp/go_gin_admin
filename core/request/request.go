package request

import (
	"admin/core"
	"admin/core/response"
	"admin/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	"path"
	"strings"
)

// 自定义验证规则
func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("file_is_exists", Validate.FileIsExists)
		_ = v.RegisterValidation("password", Validate.Password)
	}
}

// Validator 验证器接口
type Validator interface {
	// GetMessage 获取验证器自定义错误信息
	GetMessage() ValidatorMessages
}

// ValidatorMessages 验证器自定义错误信息
type ValidatorMessages map[string]interface{}

// ValidateForm 表单验证
func ValidateForm(c *gin.Context, data Validator) {
	if err := c.ShouldBind(data); err != nil {
		msg := GetErrorMsg(data, err)
		if _msg, ok := msg.(ValidateMsg); ok {
			response.Fail(c, _msg.Text, nil, response.Options{MsgArgs: _msg.Args})
		}
		response.Fail(c, msg.(string), nil)
	}
}

// ValidateFile 上传文件验证
func ValidateFile(c *gin.Context) *multipart.FileHeader {
	file, err := c.FormFile("file")
	if err != nil {
		response.Fail(c, "FailedToUploadFile", nil)
	}
	ext := path.Ext(file.Filename)
	if util.ArrUtil.InArray(ext, core.Config.UploadFileType) == false {
		fileType := strings.Join(core.Config.UploadFileType, "|")
		response.Fail(c, "InvalidFileType", nil, response.Options{MsgArgs: []interface{}{fileType}})
	}
	maxSize := core.Config.UploadFileMaxSize
	if file.Size > maxSize*1024*1024 {
		response.Fail(c, "InvalidFileSize", nil, response.Options{MsgArgs: []interface{}{maxSize}})
	}
	return file
}

// ValidateQuery Query验证
func ValidateQuery(c *gin.Context, data Validator) {
	if err := c.ShouldBindQuery(data); err != nil {
		msg := GetErrorMsg(data, err)
		if _msg, ok := msg.(ValidateMsg); ok {
			response.Fail(c, _msg.Text, nil, response.Options{MsgArgs: _msg.Args})
		}
		response.Fail(c, msg.(string), nil)
	}
}

// GetErrorMsg 获取错误信息
func GetErrorMsg(valid Validator, err error) interface{} {
	for _, v := range err.(validator.ValidationErrors) {
		if message, exist := valid.GetMessage()[v.Field()+"."+v.Tag()]; exist {
			return message
		}
		return v.Error()
	}
	return "ParameterError"
}
