package models

import (
	"context"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
)

// ArticleMgr 文章 client
type ArticleMgr struct {
	mgr *qmgo.QmgoClient
}

// NewArticleMgr 新建 ArticleMgr
func NewArticleMgr(
	ctx context.Context,
	client *qmgo.Client,
	database *qmgo.Database,
	name string,
) *ArticleMgr {
	qClient := &qmgo.QmgoClient{
		Client:     client,
		Database:   database,
		Collection: database.Collection(name),
	}

	// 创建索引 此处仅根据文章标题创建额外索引
	_ = qClient.CreateOneIndex(
		ctx,
		options.IndexModel{
			Key: []string{"title"},
		})

	return &ArticleMgr{
		mgr: qClient,
	}
}
