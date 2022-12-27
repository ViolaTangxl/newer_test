package config

import (
	"github.com/qiniu/qmgo"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

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

// LoadConfig 加载配置文件
func LoadConfig(yamlPath string) (*Config, error) {
	bytes, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = yaml.UnmarshalStrict(bytes, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
