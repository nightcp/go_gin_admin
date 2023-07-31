package handlers

import (
	"admin/app/admin/schemas/req"
	"admin/app/admin/services"
	"admin/core/request"
	"admin/core/response"
	"github.com/gin-gonic/gin"
)

type adminUserHandler struct {
}

var AdminUserHandler = &adminUserHandler{}

// List 用户-列表
// @Summary 用户列表
// @Description 分页获取用户列表
// @Tags 用户
// @Security ApiKeyAuth
// @Produce json
// @Param page_no query int true "页码, 最小为1"
// @Param page_size query int true "每页记录数, 最小为10, 最大为1000"
// @Success 200 {object} response.Resp{data=response.PageResp{list=resp.AdminUserListResp}}
// @Router /admin-users [get]
func (handler *adminUserHandler) List(c *gin.Context) {
	pageReq := request.PageReq{}
	request.ValidateQuery(c, &pageReq)
	data := services.AdminUserService.GetUserList(pageReq)
	response.Success(c, "Succeeded", data)
}

// Add 用户-新增
// @Summary 新增用户
// @Description 新增用户
// @Tags 用户
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data body req.AddAdminUserReq true "新增用户参数"
// @Success 200 {object} response.Resp
// @Router /admin-users [post]
func (handler *adminUserHandler) Add(c *gin.Context) {
	addReq := req.AddAdminUserReq{}
	request.ValidateForm(c, &addReq)
	services.AdminUserService.StoreUser(c, addReq)
	response.Success(c, "AddSucceeded", nil)
}

// Edit 用户-编辑
// @Summary 编辑用户
// @Description 编辑用户
// @Tags 用户
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param data body req.EditAdminUserReq true "编辑用户参数"
// @Success 200 {object} response.Resp
// @Router /admin-users/{id} [put]
func (handler *adminUserHandler) Edit(c *gin.Context) {
	editReq := req.EditAdminUserReq{}
	request.ValidateForm(c, &editReq)
	services.AdminUserService.UpdateUser(c, editReq)
	response.Success(c, "EditSucceeded", nil)
}

// Delete 用户-删除
// @Summary 删除用户
// @Description admin用户不可删除
// @Tags 用户
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} response.Resp
// @Router /admin-users/{id} [delete]
func (handler *adminUserHandler) Delete(c *gin.Context) {
	services.AdminUserService.DestroyUser(c)
	response.Success(c, "DeleteSucceeded", nil)
}
