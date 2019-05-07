package cache

import (
	"github.com/go-redis/redis"
	"github.com/pantazheng/bci/config"
	"github.com/pelletier/go-toml"
)

/**
*@Author: PantaZheng
*@CreateAt: 2019/5/6 23:00
*@Title: redis.go
*@Package: cache
*@Description: 测试redis存储
@Software: GoLand
*/

func New() *redis.Client {
	_ = config.Conf.Get("redis").(*toml.Tree)
	configTree := config.Conf.Get("redis").(*toml.Tree)
	return redis.NewClient(&redis.Options{
		Addr:     configTree.Get("Addr").(string),
		Password: configTree.Get("Password").(string), // no password set
		DB:       int(configTree.Get("DB").(int64)),   // 因为系统是64位的，所以默认的 int 型是 int64
	})
}
