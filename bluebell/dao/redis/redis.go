package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// 声明一个全局的rdb变量
var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.PassWord,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})
	_, err = client.Ping().Result()
	if err != nil {
		zap.L().Error("connect rdb failed...", zap.Error(err))
		return
	}
	return
}

func Close() {
	_ = client.Close()
}
