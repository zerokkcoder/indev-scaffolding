package seed

import (
	"github.com/zerokkcoder/indevsca/internal/domain/entity"
	"github.com/zerokkcoder/indevsca/internal/pkg/utils"

	"gorm.io/gorm"
)

func CreateSuperAdmin(db *gorm.DB) error {
	admin := &entity.Admin{
		Username: "admin",
	}

	// 检查是否已存在超级管理员
	var count int64
	db.Model(&entity.Admin{}).Where("username = ?", admin.Username).Count(&count)
	if count > 0 {
		return nil
	}

	// 设置并加密密码
	hashedPassword, err := utils.HashPassword("123456")
	if err != nil {
		return err
	}
	admin.Password = hashedPassword

	// 创建超级管理员
	return db.Create(admin).Error
}
