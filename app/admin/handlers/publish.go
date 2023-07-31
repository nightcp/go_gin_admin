package handlers

import (
	"admin/app/admin/services"
	"admin/core/request"
	"admin/core/response"
	"github.com/gin-gonic/gin"
)

type publishHandler struct {
}

var PublishHandler = &publishHandler{}

// UploadFile 公共-上传文件
// @Summary 上传文件
// @Description 上传文件并获取文件路径
// @Tags 公共
// @Security ApiKeyAuth
// @Produce json
// @Param file formData file true "文件参数"
// @Success 200 {object} response.Resp{data=resp.UploadFileResp}
// @Router /publish/upload [post]
func (handler *publishHandler) UploadFile(c *gin.Context) {
	file := request.ValidateFile(c)
	data := services.PublishService.UploadFile(c, file)
	response.Success(c, "UploadFileSuccessfully", data)
}
