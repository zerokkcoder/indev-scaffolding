package service

import (
	"github.com/zerokkcoder/indevsca/internal/domain/entity"
	"github.com/zerokkcoder/indevsca/internal/domain/repository"
	"github.com/zerokkcoder/indevsca/internal/pkg/errors"
	"github.com/zerokkcoder/indevsca/internal/pkg/utils"
)

// AuthService 认证服务
type AuthService struct {
	userRepo repository.UserRepository
}

// NewAuthService 创建认证服务实例
func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Login 用户登录
func (s *AuthService) Login(username, password string) (string, error) {
	// 获取用户
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.NewAppError(errors.ErrUserNotFound, "用户不存在", err)
	}

	// 验证密码
	if !user.CheckPassword(password) {
		return "", errors.NewAppError(errors.ErrInvalidPassword, "用户名或密码错误", nil)
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", errors.NewAppError(errors.ErrInternal, "生成令牌失败", err)
	}

	return token, nil
}

// Register 用户注册
func (s *AuthService) Register(username, password string) error {
	// 检查用户名是否已存在
	exists, err := s.userRepo.ExistsByUsername(username)
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "检查用户是否存在失败", err)
	}
	if exists {
		return errors.NewAppError(errors.ErrUserExists, "用户已存在", nil)
	}
	// 创建用户
	user := &entity.User{
		Username: username,
		Password: password,
	}
	if err := s.userRepo.Create(user); err != nil {
		return errors.NewAppError(errors.ErrInternal, "创建用户失败", err)
	}

	return nil
}
