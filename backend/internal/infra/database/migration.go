package database

import (
	"fmt"
	"time"

	"github.com/zerokkcoder/indevsca/internal/domain/entity"
	"github.com/zerokkcoder/indevsca/internal/infra/logger"
	"gorm.io/gorm"
)

// Migration 迁移记录
type Migration struct {
	ID        uint   `gorm:"primaryKey"`
	Version   string `gorm:"size:32;uniqueIndex;not null"`
	Name      string `gorm:"size:128;not null"`
	CreatedAt time.Time
}

// TableName 指定表名
func (Migration) TableName() string {
	return "migrations"
}

// Migrator 迁移管理器
type Migrator struct {
	db *gorm.DB
}

// NewMigrator 创建迁移管理器
func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{db: db}
}

// Initialize 初始化迁移表
func (m *Migrator) Initialize() error {
	// 创建迁移记录表
	if err := m.db.AutoMigrate(&Migration{}); err != nil {
		return fmt.Errorf("创建迁移表失败: %w", err)
	}
	return nil
}

// HasMigrated 检查是否已经迁移
func (m *Migrator) HasMigrated(version string) bool {
	var count int64
	m.db.Model(&Migration{}).Where("version = ?", version).Count(&count)
	return count > 0
}

// recordMigration 记录迁移
func (m *Migrator) recordMigration(version, name string) error {
	migration := &Migration{
		Version: version,
		Name:    name,
	}
	return m.db.Create(migration).Error
}

// RunMigrations 运行所有迁移
func (m *Migrator) RunMigrations() error {
	logger.Info("开始数据库迁移")

	// 初始化迁移表
	if err := m.Initialize(); err != nil {
		return err
	}

	// 定义迁移列表
	migrations := []struct {
		Version string
		Name    string
		Fn      func(*gorm.DB) error
	}{
		{
			Version: "202401010001",
			Name:    "create_user_table_202401010001",
			Fn: func(db *gorm.DB) error {
				return db.AutoMigrate(&entity.User{})
			},
		},
		// 在这里添加更多迁移...
	}

	// 执行迁移
	for _, migration := range migrations {
		// 检查是否已经迁移
		if m.HasMigrated(migration.Version) {
			logger.Info("跳过已执行的迁移",
				"version", migration.Version,
				"name", migration.Name,
			)
			continue
		}

		logger.Info("执行迁移",
			"version", migration.Version,
			"name", migration.Name,
		)

		// 开始事务
		err := m.db.Transaction(func(tx *gorm.DB) error {
			// 执行迁移
			if err := migration.Fn(tx); err != nil {
				return fmt.Errorf("迁移失败: %w", err)
			}

			// 记录迁移
			if err := m.recordMigration(migration.Version, migration.Name); err != nil {
				return fmt.Errorf("记录迁移失败: %w", err)
			}

			return nil
		})

		if err != nil {
			return fmt.Errorf("迁移 %s 失败: %w", migration.Version, err)
		}

		logger.Info("迁移完成",
			"version", migration.Version,
			"name", migration.Name,
		)
	}

	logger.Info("数据库迁移完成")
	return nil
}

// GetMigrationHistory 获取迁移历史
func (m *Migrator) GetMigrationHistory() ([]Migration, error) {
	var migrations []Migration
	err := m.db.Order("id desc").Find(&migrations).Error
	return migrations, err
}
