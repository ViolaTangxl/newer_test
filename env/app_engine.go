package env

import (
	"context"
	"github.com/ViolaTangxl/newer_test/config"
	"github.com/ViolaTangxl/newer_test/models"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"github.com/sirupsen/logrus"
)

var Global GlobalEnv

// GlobalEnv 定义全局变量方便使用
type GlobalEnv struct {
	Cfg        *config.Config
	Logger     *logrus.Logger
	ArticleMgr *models.ArticleMgr
}

// InitMongo 初始化 mongo
func InitMongo(ctx context.Context, log *logrus.Logger, cfg *config.Config) *models.ArticleMgr {
	client, err := qmgo.NewClient(ctx, &cfg.MgoConfig)
	if err != nil {
		log.WithError(err).Fatal("init mongo client failed")
	}
	dataBase := client.Database(cfg.MgoConfig.Database)
	articlesMgr := models.NewArticleMgr(ctx, client, dataBase, "articles")
	return articlesMgr
}

// InitApp 初始化gin 相关
func InitApp(log *logrus.Logger, cfg *config.Config) *gin.Engine {
	// TODO 抽取配置变量
	gin.SetMode("test")

	app := gin.New()

	// TODO 其他前置操作

	return app
}
