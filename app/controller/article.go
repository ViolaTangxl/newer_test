package controller

import (
	"context"
	"github.com/ViolaTangxl/newer_test/models"
	"github.com/gin-gonic/gin"
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
	// TODO
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
	// TODO
}
