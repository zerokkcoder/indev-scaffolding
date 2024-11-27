package admin

import (
	"time"

	"github.com/zerokkcoder/indevsca/internal/domain/entity"
	"github.com/zerokkcoder/indevsca/internal/pkg/response"
	"github.com/zerokkcoder/indevsca/internal/pkg/utils"
	"gorm.io/gorm"
)

// AuthService 管理员认证服务
type AuthService struct {
	db *gorm.DB
}

// NewAuthService 创建管理员认证服务实例
func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

// Login 管理员登录
func (s *AuthService) Login(username, password string) (string, *entity.Admin, error) {
	var admin entity.Admin
	if err := s.db.Where("username = ?", username).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil, response.ErrInvalidCredentials
		}
		return "", nil, err
	}

	// 验证密码
	if !utils.CheckPasswordHash(password, admin.Password) {
		return "", nil, response.ErrInvalidCredentials
	}

	// 生成JWT token
	token, err := utils.GenerateToken(admin.ID)
	if err != nil {
		return "", nil, err
	}

	// 更新最后登录时间
	admin.LastLogin = time.Now()
	if err := s.db.Save(&admin).Error; err != nil {
		return "", nil, err
	}

	return token, &admin, nil
}
