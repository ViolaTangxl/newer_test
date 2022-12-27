package main

import "flag"

func main() {
	// 加载配置文件
	var configPath string

	flag.StringVar(&configPath, "conf", "config.yml", "config file path")
	flag.Parse()

}
