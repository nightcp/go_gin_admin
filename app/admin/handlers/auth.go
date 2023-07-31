package handlers

import (
	"admin/app/admin/schemas/req"
	"admin/app/admin/services"
	"admin/core/request"
	"admin/core/response"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
}

var AuthHandler = &authHandler{}

// Login 认证-登录
// @Summary 用户登录
// @Description 获取访问令牌
// @Tags 认证
// @Accept json
// @Produce json
// @Param data body req.LoginReq true "登录参数"
// @Success 200 {object} response.Resp{data=resp.LoginResp}
// @Router /auth/login [post]
func (handler *authHandler) Login(c *gin.Context) {
	var loginReq req.LoginReq
	request.ValidateForm(c, &loginReq)
	data := services.AuthService.Login(c, loginReq)
	response.Success(c, "LoginSucceeded", data)
}

// Profile 认证-用户信息
// @Summary 用户信息
// @Description 获取当前用户数据
// @Tags 认证
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} response.Resp{data=resp.ProfileResp}
// @Router /auth/profile [get]
func (handler *authHandler) Profile(c *gin.Context) {
	data := services.AuthService.GetProfile(c)
	response.Success(c, "Succeeded", data)
}

// EditProfile 认证-修改用户信息
// @Summary 修改用户信息
// @Description 修改用户昵称, 头像
// @Tags 认证
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data body req.EditProfileReq true "修改用户信息参数"
// @Success 200 {object} response.Resp
// @Router /auth/profile [put]
func (handler *authHandler) EditProfile(c *gin.Context) {
	updateReq := req.EditProfileReq{}
	request.ValidateForm(c, &updateReq)
	services.AuthService.UpdateProfile(c, updateReq)
	response.Success(c, "EditSucceeded", nil)
}

// EditPassword 认证-修改用户密码
// @Summary 修改用户密码
// @Description 设置新密码, 之后使用新密码登录
// @Tags 认证
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param date body req.EditPasswordReq true "修改用户密码参数"
// @Success 200 {object} response.Resp
// @Router /auth/password [put]
func (handler *authHandler) EditPassword(c *gin.Context) {
	updateReq := req.EditPasswordReq{}
	request.ValidateForm(c, &updateReq)
	services.AuthService.UpdatePassword(c, updateReq)
	response.Success(c, "EditSucceeded", nil)
}
