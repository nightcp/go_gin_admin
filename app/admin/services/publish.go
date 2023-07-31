package services

import (
	"admin/app/admin/schemas/resp"
	"admin/core"
	"admin/core/response"
	"admin/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"mime/multipart"
	"os"
	"path"
)

type publishService struct {
}

var PublishService = &publishService{}

// UploadFile 上传文件
func (service *publishService) UploadFile(c *gin.Context, file *multipart.FileHeader) resp.UploadFileResp {
	curDate := carbon.Now().ToDateString()
	dirPath := core.Config.UploadFilePath + curDate
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		core.Logger.Error("MkdirAll: " + err.Error())
		response.Fail(c, "FailedToUploadFile", nil)
	}
	fileName := util.StrUtil.MakeRandomStr(32) + path.Ext(file.Filename)
	filePath := path.Join(dirPath, fileName)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		core.Logger.Error("Save uploaded file: " + err.Error())
		response.Fail(c, "FailedToUploadFile", nil)
	}
	return resp.UploadFileResp{Path: filePath}
}
