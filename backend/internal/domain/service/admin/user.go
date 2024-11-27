package admin

import (
	"github.com/zerokkcoder/indevsca/internal/domain/entity"
	"gorm.io/gorm"
)

// UserService 管理员用户服务
type UserService struct {
	db *gorm.DB
}

// NewUserService 创建管理员用户服务
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// List 获取用户列表
func (s *UserService) List(page, size int) ([]*entity.User, int64, error) {
	var users []*entity.User
	var total int64
	offset := (page - 1) * size

	if err := s.db.Model(&entity.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := s.db.Offset(offset).Limit(size).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Get 获取用户详情
func (s *UserService) Get(id uint) (*entity.User, error) {
	var user entity.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update 更新用户
func (s *UserService) Update(user *entity.User) error {
	return s.db.Model(user).Updates(map[string]interface{}{
		"nickname": user.Nickname,
		"phone":    user.Phone,
		"email":    user.Email,
		"status":   user.Status,
	}).Error
}

// Delete 删除用户
func (s *UserService) Delete(id uint) error {
	return s.db.Delete(&entity.User{}, id).Error
}

// Enable 启用用户
func (s *UserService) Enable(id uint) error {
	user := &entity.User{
		Model: gorm.Model{ID: id},
	}
	return user.Enable(s.db)
}

// Disable 禁用用户
func (s *UserService) Disable(id uint) error {
	user := &entity.User{
		Model: gorm.Model{ID: id},
	}
	return user.Disable(s.db)
}
