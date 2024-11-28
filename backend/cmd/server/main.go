package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/zerokkcoder/indevsca/cmd/server/wire"
	"github.com/zerokkcoder/indevsca/pkg/config"
	"github.com/zerokkcoder/indevsca/pkg/log"
)

func main() {
	// 初始化配置
	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf, err := config.NewConfig(*envConf)
	if err != nil {
		panic(err)
	}
	// 初始化日志
	logger := log.NewLogger(conf)
	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	logger.Info("server start", "host", fmt.Sprintf("http://%v:%d", conf.App.Host, conf.App.Port))
	// 启动服务器
	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}

}
