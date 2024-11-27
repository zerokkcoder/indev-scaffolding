package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/app/requests/admin"
	"github.com/zerokkcoder/indevsca/internal/domain/entity"
	"github.com/zerokkcoder/indevsca/internal/domain/service"
	"github.com/zerokkcoder/indevsca/internal/pkg/response"
	"gorm.io/gorm"
)

// UserHandler 用户管理处理器
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler 创建用户管理处理器
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// List 获取用户列表
func (h *UserHandler) List(c *gin.Context) {
	var req admin.ListUserRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, response.ErrInvalidParams)
		return
	}

	users, total, err := h.userService.List(req.Page, req.Size)
	if err != nil {
		response.Error(c, err)
		return
	}

	// 转换为响应结构
	var userResponses []*admin.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &admin.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			Email:     user.Email,
			Phone:     user.Phone,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	response.Success(c, gin.H{
		"total": total,
		"items": userResponses,
	})
}

// Update 更新用户
func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, response.ErrInvalidParams)
		return
	}

	var req admin.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ErrInvalidParams)
		return
	}

	user := &entity.User{
		Model: gorm.Model{
			ID: uint(id),
		},
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	if err := h.userService.Update(user); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}

// Delete 删除用户
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, response.ErrInvalidParams)
		return
	}

	if err := h.userService.Delete(uint(id)); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}

// Enable 启用用户
func (h *UserHandler) Enable(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, response.ErrInvalidParams)
		return
	}

	if err := h.userService.Enable(uint(id)); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}

// Disable 禁用用户
func (h *UserHandler) Disable(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, response.ErrInvalidParams)
		return
	}

	if err := h.userService.Disable(uint(id)); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}
