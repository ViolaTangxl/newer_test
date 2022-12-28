package models

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestSaveArticle(t *testing.T) {
	client := initClient()
	ctx := context.Background()
	articleMgr := NewArticleMgr(ctx, client, client.Database("test"), "articles")

	defer func() {
		_ = articleMgr.mgr.Close(ctx)
		_ = client.Close(ctx)
	}()

	// 批量插入1000W数据，每1W条一组插入
	models := make([]Articles, 0)
	now := time.Now()
	start := now

	fmt.Printf("开始批量插入数据，当前时间为：%s \n", start)
	for i := 1; i < 10000001; i++ {
		if i%10000 == 0 {
			err := articleMgr.SaveArticle(ctx, models)
			assert.NoError(t, err)

			now = time.Now()
			models = models[0:0]
		}
		idStr := time.Now().Add(time.Minute * time.Duration(i))
		id := primitive.NewObjectIDFromTimestamp(idStr)
		article := Articles{
			Id:            id,
			Title:         fmt.Sprintf("文章%d", i),
			ArticleBanner: "http://www.huayuhua.com/uploads/images/20200506/1588757921502785.jpg",
			Context:       fmt.Sprintf("文章正文%d", i),
			CreateAt:      now,
			UpdateAt:      now,
		}
		models = append(models, article)
	}
	end := time.Now()
	fmt.Printf("批量插入数据完成，当前时间为：%s, 共耗时:%.2f 分钟 \n", end, end.Sub(start).Minutes())
}
