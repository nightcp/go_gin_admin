package redis

import (
	"admin/core"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type rdbHelper struct {
}

var RDB = initRedis()
var RDBHelper = &rdbHelper{}

// initRedis 初始化Redis
func initRedis() *redis.Client {
	opt, err := redis.ParseURL(core.Config.RedisUrl)
	if err != nil {
		panic("InitRedis connect err: " + err.Error())
	}
	opt.PoolSize = core.Config.RedisPoolSize
	client := redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		core.Logger.Error("InitRedis client.Ping err: " + err.Error())
	}
	return client
}
