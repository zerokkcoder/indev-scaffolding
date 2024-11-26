//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/zerokkcoder/indevsca/internal/app/handler"
	"github.com/zerokkcoder/indevsca/internal/app/handler/admin"
	"github.com/zerokkcoder/indevsca/internal/app/handler/mobile"
	"github.com/zerokkcoder/indevsca/internal/domain/repository"
	"github.com/zerokkcoder/indevsca/internal/domain/service"
	"github.com/zerokkcoder/indevsca/internal/infra/config"
	"github.com/zerokkcoder/indevsca/internal/infra/database"
	infraRepo "github.com/zerokkcoder/indevsca/internal/repository"
)

// InitializeApp 初始化应用
func InitializeApp() (*App, error) {
	wire.Build(
		// 基础设置
		config.New,
		database.New,

		// 仓储层
		infraRepo.NewUserRepository,
		wire.Bind(new(repository.UserRepository), new(*infraRepo.UserRepository)),

		// 服务层
		service.NewAuthService,
		service.NewUserService,

		// 处理器 - 管理端
		admin.NewAuthHandler,
		admin.NewUserHandler,

		// 处理器 - 移动端
		mobile.NewAuthHandler,
		mobile.NewUserHandler,

		// 处理器集合
		handler.NewHandlers,

		// 应用
		NewApp,
	)
	return nil, nil
}
