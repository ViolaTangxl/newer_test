package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Articles struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`                        // ObjectID
	Title         string             `bson:"title" json:"title"`                   // 文章标题
	ArticleBanner string             `bson:"article_banner" json:"article_banner"` // 头图
	Context       string             `bson:"context" json:"context"`               // 文章正文内容
}
