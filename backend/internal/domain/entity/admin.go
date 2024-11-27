package entity

import (
	"time"

	"gorm.io/gorm"
)

// Admin 管理员
type Admin struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"` // 使用 - 在JSON中隐藏密码
	Avatar    string         `json:"avatar" gorm:"size:256;comment:'头像'"`
	Status    int            `json:"status" gorm:"default:1;comment:'状态 1: 正常, 0: 禁用'"`
	LastLogin time.Time      `json:"last_login" gorm:"comment:'最后登录时间'"`
	LoginIP   string         `json:"login_ip" gorm:"size:64;comment:'最后登录IP'"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (a *Admin) TableName() string {
	return "admin"
}
