package main

import (
	"context"
	"flag"
	"github.com/ViolaTangxl/newer_test/config"
	"github.com/ViolaTangxl/newer_test/env"
	"github.com/sirupsen/logrus"
)

func main() {
	// 加载配置文件
	var configPath string

	flag.StringVar(&configPath, "conf", "config.yml", "config file path")
	flag.Parse()

	conf, err := config.LoadConfig(configPath)
	if err != nil {
		logrus.WithError(err).Fatal("failed to load config")
	}

	ctx := context.Background()
	logger := logrus.StandardLogger()
	env.Global.Cfg = conf
	env.Global.Logger = logger
	env.Global.ArticleMgr = env.InitMongo(ctx, logger, conf)

	// 初始化项目配置
	app := env.InitApp(logger, conf)

	// 绑定路由
	env.InitRouters(app)

	// TODO 抽取配置文件
	err = app.Run(":11009")
	if err != nil {
		logger.WithError(err).Fatal("service run failed")
	}
}
