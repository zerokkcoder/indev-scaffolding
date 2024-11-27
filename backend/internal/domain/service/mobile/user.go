package mobile

import (
	"github.com/zerokkcoder/indevsca/internal/domain/entity"
	"gorm.io/gorm"
)

// UserService 移动端用户服务
type UserService struct {
	db *gorm.DB
}

// NewUserService 创建移动端用户服务
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// Profile 获取用户信息
func (s *UserService) Profile(id uint) (*entity.User, error) {
	var user entity.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateProfile 更新用户信息
func (s *UserService) UpdateProfile(user *entity.User) error {
	return s.db.Model(user).Updates(map[string]interface{}{
		"nickname": user.Nickname,
		"avatar":   user.Avatar,
	}).Error
}

// UpdatePassword 更新密码
func (s *UserService) UpdatePassword(id uint, newPassword string) error {
	user := &entity.User{
		Model: gorm.Model{ID: id},
		Password: newPassword,
	}
	return s.db.Model(user).Update("password", newPassword).Error
}

// BindPhone 绑定手机号
func (s *UserService) BindPhone(id uint, phone string) error {
	return s.db.Model(&entity.User{}).Where("id = ?", id).Update("phone", phone).Error
}

// BindEmail 绑定邮箱
func (s *UserService) BindEmail(id uint, email string) error {
	return s.db.Model(&entity.User{}).Where("id = ?", id).Update("email", email).Error
}
