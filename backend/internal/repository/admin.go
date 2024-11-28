package repository

import (
	"context"

	"github.com/zerokkcoder/indevsca/internal/model"
)

type AdminRepository interface {
	GetByUserName(ctx context.Context, username string) (*model.Admin, error)
	// Create(ctx context.Context, user *model.User) error
	// Update(ctx context.Context, user *model.User) error
	// GetByID(ctx context.Context, id string) (*model.User, error)
	// GetByEmail(ctx context.Context, email string) (*model.User, error)
}

func NewAdminRepository(
	r *Repository,
) AdminRepository {
	return &adminRepository{
		Repository: r,
	}
}

type adminRepository struct {
	*Repository
}

// GetByUserName 通过用户名获取管理员
func (r *adminRepository) GetByUserName(ctx context.Context, username string) (*model.Admin, error) {
	var admin model.Admin
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
