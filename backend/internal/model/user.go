package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户
type User struct {
	gorm.Model

	Username  string    `gorm:"size:32;comment:'用户名';uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"size:128;not null" json:"-"` // 密码哈希
	Nickname  string    `gorm:"size:32;comment:'昵称'" json:"nickname"`
	Phone     string    `gorm:"size:32;uniqueIndex;comment:'手机号'" json:"phone"`
	Email     string    `gorm:"size:128;uniqueIndex;comment:'邮箱'" json:"email"`
	Avatar    string    `gorm:"size:256;comment:'头像'" json:"avatar"`
	Sex       int       `gorm:"default:0;comment:'性别 0: 未知, 1: 男, 2: 女'" json:"sex"`
	Status    int       `gorm:"default:1;comment:'状态 1: 正常, 0: 禁用'" json:"status"`
	LoginIP   string    `gorm:"size:64;comment:'最后登录IP" json:"login_ip"`
	LastLogin time.Time `gorm:"comment:'最后登录时间" json:"last_login"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
