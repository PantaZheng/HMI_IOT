package config

import (
	"github.com/pelletier/go-toml"
	"log"
)

var (
	Conf = New()
)

type Config struct {
	MySql  MySqlConfig
	Wechat WechatConfig
	APP    AppConfig
}

type MySqlConfig struct {
	DBDriver  string
	DBName    string
	UserName  string
	Password  string
	Charset   string
	Collation string
}

type WechatConfig struct {
	AppID         string
	AppSecret     string
	OriID         string
	Token         string
	EncodedAESKEY string
}

type AppConfig struct {
	Address string
	Port    int
}

/**
 * 返回单例实例
 * @method New
 */
func New() *Config {
	config := &Config{}
	tomlTree, err := toml.LoadFile("./config/config.toml")
	if err == nil {
		err = tomlTree.Unmarshal(config)
	}
	if err != nil {
		log.Println("TomlError ", err.Error())
	}
	return config
}
