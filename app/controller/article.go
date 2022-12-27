package controller

import (
	"context"
	"github.com/ViolaTangxl/newer_test/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type ArticleHandler struct {
	ctx context.Context
	mgr *models.ArticleMgr
}

// NewArticleHandler 新建文章handler
func NewArticleHandler(ctx context.Context, mgr *models.ArticleMgr) *ArticleHandler {
	return &ArticleHandler{
		ctx: ctx,
		mgr: mgr,
	}
}

// GetArticlesList 获取文章列表
func (h *ArticleHandler) GetArticlesList(ctx *gin.Context) {
	var param models.ArticlesListCron
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	if param.Page == 0 {
		param.Page = 1
	}
	if param.PageSize == 0 {
		// TODO 这里写死的默认10条
		param.PageSize = uint64(10)
	}

	param.Page = param.PageSize * (param.Page - 1)

	list, count, err := h.mgr.GetArticlesList(context.Background(), param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Artiles []models.Articles
		Count   uint64
	}{
		Artiles: list,
		Count:   count,
	})
}

// CreateArticles 新增文章
func (h *ArticleHandler) CreateArticles(ctx *gin.Context) {
	var articles = make([]models.Articles, 0)
	if err := ctx.ShouldBindJSON(&articles); err != nil {
		// TODO 抽取公用的返回体包装
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	var now = time.Now()
	for i := range articles {
		articles[i].Id = primitive.NewObjectID()
		articles[i].CreateAt = now
		articles[i].UpdateAt = now
	}

	err := h.mgr.SaveArticle(context.Background(), articles)
	if err != nil {
		// TODO log
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// UpdateArticles 文章编辑
func (h *ArticleHandler) UpdateArticles(ctx *gin.Context) {
	var article models.UpdateArticleParam
	if err := ctx.ShouldBindJSON(&article); err != nil {
		// TODO 抽取公用的返回体包装
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 如果id字段不存在，则报错
	if article.Id == "" {
		logrus.Error("invalid article id")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "invalid article id",
		})
	}

	id, err := primitive.ObjectIDFromHex(article.Id)
	if err != nil {
		logrus.Error("invalid article id")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "invalid article id",
		})
	}

	article.UpdateAt = time.Now()
	err = h.mgr.UpdateArticle(context.Background(), id, article)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
