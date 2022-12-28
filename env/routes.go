package env

import (
	"context"
	"time"

	"github.com/ViolaTangxl/newer_test/app/controller"
	cache "github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
)

func InitRouters(app *gin.Engine) {
	ctx := context.Background()
	articleHandler := controller.NewArticleHandler(ctx, Global.ArticleMgr)

	article := app.Group("/articles")
	{
		article.GET("/list",
			cache.CacheByRequestURI(Global.RedisStore, 5*time.Minute),
			articleHandler.GetArticlesList,
		)
		article.POST("", articleHandler.CreateArticles)
		article.PUT("", articleHandler.UpdateArticles)
	}
}
