package mobile

import (
	"time"

	"github.com/zerokkcoder/indevsca/internal/domain/entity"
	"github.com/zerokkcoder/indevsca/internal/pkg/utils"
	"gorm.io/gorm"
)

// AuthService 移动端认证服务
type AuthService struct {
	db *gorm.DB
}

// NewAuthService 创建移动端认证服务
func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

// Login 用户登录
func (s *AuthService) Login(username, password string) (string, *entity.User, error) {
	var user entity.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", nil, nil
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", nil, err
	}

	// 更新最后登录时间
	user.LastLogin = time.Now()
	if err := s.db.Save(&user).Error; err != nil {
		return "", nil, err
	}

	return token, &user, nil
}

// Register 用户注册
func (s *AuthService) Register(user *entity.User) error {
	// 检查用户名是否已存在
	var count int64
	if err := s.db.Model(&entity.User{}).Where("username = ?", user.Username).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	// 设置默认状态
	user.Status = entity.UserStatusEnabled

	// 创建用户
	return s.db.Create(user).Error
}

// SendVerificationCode 发送验证码
func (s *AuthService) SendVerificationCode(phone string) error {
	// TODO: 实现验证码发送逻辑
	return nil
}
