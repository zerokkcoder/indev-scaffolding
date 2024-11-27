package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
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

const (
	// SexUnknown 未知
	SexUnknown = 0
	// SexMale 男
	SexMale = 1
	// SexFemale 女
	SexFemale = 2
)

const (
	// UserStatusEnabled 用户正常
	UserStatusEnabled = 1
	// UserStatusDisabled 用户禁用
	UserStatusDisabled = 0
)

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// BeforeCreate 创建前的钩子
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password == "" {
		return errors.New("密码不能为空")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// 设置默认值
	if u.Status == 0 {
		u.Status = UserStatusEnabled
	}
	if u.Nickname == "" {
		u.Nickname = u.Username
	}
	return nil
}

// BeforeUpdate 更新前的钩子
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// 如果更新密码，需要加密
	if tx.Statement.Changed("Password") {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// CheckPassword 检查密码是否正确
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// UpdateLastLogin 更新最后登录时间
func (u *User) UpdateLastLogin(db *gorm.DB) error {
	now := time.Now()
	u.LastLogin = now
	return db.Model(u).UpdateColumn("last_login", now).Error
}

// IsEnabled 检查用户是否启用
func (u *User) IsEnabled() bool {
	return u.Status == UserStatusEnabled
}

// Enable 启用用户
func (u *User) Enable(db *gorm.DB) error {
	u.Status = UserStatusEnabled
	return db.Model(u).UpdateColumn("status", UserStatusEnabled).Error
}

// Disable 禁用用户
func (u *User) Disable(db *gorm.DB) error {
	u.Status = UserStatusDisabled
	return db.Model(u).UpdateColumn("status", UserStatusDisabled).Error
}

// GetByUsername 根据用户名获取用户
func (u *User) GetByUsername(db *gorm.DB) error {
	return db.Model(u).Where("username = ?", u.Username).First(u).Error
}
