package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Articles struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`                        // ObjectID
	Title         string             `bson:"title" json:"title"`                   // 文章标题
	ArticleBanner string             `bson:"article_banner" json:"article_banner"` // 头图
	Context       string             `bson:"context" json:"context"`               // 文章正文内容
	CreateAt      time.Time          `bson:"create_at" json:"create_at"`           // 创建时间
	UpdateAt      time.Time          `bson:"update_at" json:"update_at"`           // 修改时间
}

// SaveArticle 保存文章 (考虑到新增测试数据方法，支持批量新增)
func (d *ArticleMgr) SaveArticle(ctx context.Context, models []Articles) error {
	_, err := d.mgr.InsertMany(ctx, models)
	if err != nil {
		return err
	}

	return nil
}

// UpdateArticle 修改文章
// TODO
// 考虑文章正文如果内容较大且假设没有修改的情况，现在的操作比较浪费资源，可以讲正文内容hash以后存储hash string，每次修改对比原有hash
// 和新hash string是否一致，如果不一致则再修改正文
func (d *ArticleMgr) UpdateArticle(ctx context.Context, id primitive.ObjectID, model Articles) error {
	err := d.mgr.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{
			"title":          model.Title,
			"article_banner": model.ArticleBanner,
			"context":        model.Context,
			"update_at":      model.UpdateAt,
		}})
	if err != nil {
		return err
	}
	return nil
}

// ArticlesListCron
type ArticlesListCron struct {
	Id    string
	Title string
}

// GetArticlesList 根据条件分页查询文章列表
func (d *ArticleMgr) GetArticlesList(
	ctx context.Context,
	param ArticlesListCron,
	page uint,
	pageSize uint,
) ([]Articles, uint64, error) {
	result := make([]Articles, 0)
	query := bson.M{}
	if param.Id != "" {
		query["_id"], _ = primitive.ObjectIDFromHex(param.Id)
	}
	if param.Title != "" {
		query["title"] = primitive.Regex{
			Pattern: ".*" + param.Title + ".*",
			Options: "i",
		}
	}

	err := d.mgr.Find(ctx, query).Skip(int64(page)).Limit(int64(pageSize)).All(&result)
	if err != nil {
		return nil, 0, err
	}
	count, err := d.mgr.Find(ctx, query).Count()
	if err != nil {
		return nil, 0, err
	}
	return result, uint64(count), nil
}
