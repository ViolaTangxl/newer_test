package main

import (
	"flag"
	"github.com/ViolaTangxl/newer_test/config"
	"github.com/sirupsen/logrus"
)

func main() {
	// 加载配置文件
	var configPath string

	flag.StringVar(&configPath, "conf", "config.yml", "config file path")
	flag.Parse()

	_, err := config.LoadConfig(configPath)
	if err != nil {
		logrus.WithError(err).Fatal("failed to load config")
	}

}
