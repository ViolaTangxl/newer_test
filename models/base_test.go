package models

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
)

func initClient() *qmgo.Client {
	cfg := qmgo.Config{
		Uri: "mongodb://localhost:27017",
	}
	var cTimeout int64 = 0
	var sTimeout int64 = 500000
	var maxPoolSize uint64 = 30000
	var minPoolSize uint64 = 0
	cfg.ConnectTimeoutMS = &cTimeout
	cfg.SocketTimeoutMS = &sTimeout
	cfg.MaxPoolSize = &maxPoolSize
	cfg.MinPoolSize = &minPoolSize
	client, err := qmgo.NewClient(context.Background(), &cfg)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return client
}
