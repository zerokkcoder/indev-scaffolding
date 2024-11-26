package app

import (
	"github.com/google/wire"
	"github.com/zerokkcoder/indevsca/internal/app/handler"
	"github.com/zerokkcoder/indevsca/internal/domain/repository"
	"github.com/zerokkcoder/indevsca/internal/domain/service"
	"github.com/zerokkcoder/indevsca/internal/infra/cache"
	"github.com/zerokkcoder/indevsca/internal/infra/config"
	"github.com/zerokkcoder/indevsca/internal/infra/database"
	infraRepo "github.com/zerokkcoder/indevsca/internal/repository"
)

var (
	// 基础设置
	infraSet = wire.NewSet(
		config.New,
		database.New,
		cache.New,
	)

	// 仓储层
	repoSet = wire.NewSet(
		infraRepo.NewUserRepository,
		wire.Bind(new(repository.UserRepository), new(*infraRepo.UserRepository)),
	)

	// 服务层
	serviceSet = wire.NewSet(
		service.NewAuthService,
	)

	// 处理器provider
	handlerSet = wire.NewSet(
		handler.NewAuthHandler,
		handler.NewHandlers,
	)

	// 完整应用集合
	appSet = wire.NewSet(
		infraSet,
		repoSet,
		serviceSet,
		handlerSet,
		NewApp,
	)
)

// InitializeApp 初始化应用
func InitializeApp() (*App, error) {
	wire.Build(appSet)
	return nil, nil
}
