package handlers

import (
	"admin/app/admin/schemas/req"
	"admin/app/admin/services"
	"admin/core/request"
	"admin/core/response"
	"github.com/gin-gonic/gin"
)

type adminRoleHandler struct {
}

var AdminRoleHandle = &adminRoleHandler{}

// List 角色-列表
// @Summary 角色列表
// @Description 分页获取角色列表
// @Tags 角色
// @Security ApiKeyAuth
// @Produce json
// @Param page_no query int true "页码, 最小为1"
// @Param page_size query int true "每页记录数, 最小为10, 最大为1000"
// @Success 200 {object} response.Resp{data=response.PageResp{list=[]resp.AdminRoleListResp}}
// @Router /admin-roles [get]
func (handler *adminRoleHandler) List(c *gin.Context) {
	pageReq := request.PageReq{}
	request.ValidateQuery(c, &pageReq)
	data := services.AdminRoleService.GetRoleList(pageReq)
	response.Success(c, "Succeeded", data)
}

// Add 角色-新增
// @Summary 新增角色
// @Description 新增角色
// @Tags 角色
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data body req.AddAdminRoleReq true "新增角色参数"
// @Success 200 {object} response.Resp
// @Router /admin-roles [post]
func (handler *adminRoleHandler) Add(c *gin.Context) {
	addReq := req.AddAdminRoleReq{}
	request.ValidateForm(c, &addReq)
	services.AdminRoleService.StoreRole(c, addReq)
	response.Success(c, "AddSucceeded", nil)
}

// Edit 角色-编辑
// @Summary 编辑角色
// @Description 编辑角色
// @Tags 角色
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param data body req.EditAdminRoleReq true "编辑角色参数"
// @Success 200 {object} response.Resp
// @Router /admin-roles/{id} [put]
func (handler *adminRoleHandler) Edit(c *gin.Context) {
	editReq := req.EditAdminRoleReq{}
	request.ValidateForm(c, &editReq)
	services.AdminRoleService.UpdateRole(c, editReq)
	response.Success(c, "EditSucceeded", nil)
}

// Delete 角色-删除
// @Summary 删除角色
// @Description 删除角色
// @Tags 角色
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} response.Resp
// @Router /admin-roles/{id} [delete]
func (handler *adminRoleHandler) Delete(c *gin.Context) {
	services.AdminRoleService.DestroyRole(c)
	response.Success(c, "DeleteSucceeded", nil)
}
