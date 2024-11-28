package infra

import (
	"fmt"
	"time"

	"github.com/zerokkcoder/indevsca/pkg/config"
	"github.com/zerokkcoder/indevsca/pkg/log"
	"github.com/zerokkcoder/indevsca/pkg/sloggorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewDB 创建数据库
func NewDB(conf *config.Config, l *log.Logger) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	logger := sloggorm.New(l.Logger)
	driver := conf.Database.Driver
	dsn := ""
	switch driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			conf.Database.Username, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.Database, conf.Database.Charset)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger,
		})
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			conf.Database.Host, conf.Database.Port, conf.Database.Username, conf.Database.Password, conf.Database.Database)
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
	case "sqlite":
		dsn = fmt.Sprintf("%s?cache=shared&charset=%s&parseTime=True&loc=Local",
			conf.Database.Database, conf.Database.Charset)
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		panic("unknown db driver")
	}
	if err != nil {
		panic(err)
	}
	if conf.App.Mode == "debug" {
		db = db.Debug()
	}

	// Connection Pool config
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
