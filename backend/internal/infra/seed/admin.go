package seed

import (
	"time"

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

	// 设置默认头像
	admin.Avatar = "https://ui-avatars.com/api/?name=Admin&background=random&color=fff&length=2&font-size=0.33&rounded=true"
	admin.LastLogin = time.Now()
	admin.LoginIP = "127.0.0.1"

	// 创建超级管理员
	return db.Create(admin).Error
}
