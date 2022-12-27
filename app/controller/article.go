package controller

import (
	"context"
	"github.com/ViolaTangxl/newer_test/models"
	"github.com/gin-gonic/gin"
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
	// TODO
}

// UpdateArticles 文章编辑
func (h *ArticleHandler) UpdateArticles(ctx *gin.Context) {
	// TODO
}
