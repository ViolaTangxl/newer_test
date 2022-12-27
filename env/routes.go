package env

import (
	"context"
	"github.com/ViolaTangxl/newer_test/app/controller"
	"github.com/gin-gonic/gin"
)

func InitRouters(app *gin.Engine) {
	ctx := context.Background()
	articleHandler := controller.NewArticleHandler(ctx, Global.ArticleMgr)

	article := app.Group("/articles")
	{
		article.GET("/list", articleHandler.GetArticlesList)
		article.POST("", articleHandler.CreateArticles)
		article.PUT("", articleHandler.UpdateArticles)
	}
}
