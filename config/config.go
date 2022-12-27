package config

import "github.com/qiniu/qmgo"

// Config 配置定义
type Config struct {
	RedisConfig RedisConfig `yaml:"redis_config"`
	MgoConfig   qmgo.Config `yaml:"mongo_config"`
}

// RedisConfig redis 配置结构定义
type RedisConfig struct {
	Addrs      []string `yaml:"addrs"`
	MasterName string   `yaml:"master_name"`
	Failover   bool     `yaml:"failover"`
	Password   string   `yaml:"password"`
	DB         int      `yaml:"db"`
	Size       int      `yaml:"size"`
	Networt    string   `yaml:"network"`
	KeyPairs   string   `yaml:"key_pairs"`
}
