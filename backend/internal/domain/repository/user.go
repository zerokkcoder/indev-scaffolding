package repository

import (
	"github.com/zerokkcoder/indevsca/internal/domain/entity"
)

// UserRepository 用户仓储接口
type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id uint) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
	ExistsByUsername(username string) (bool, error)
}
