package redis

import (
	"admin/core"
	"context"
	"go.uber.org/zap"
	"time"
)

// KSet 设置Key/value
func (helper *rdbHelper) KSet(key string, value interface{}, seconds int) bool {
	err := RDB.Set(context.Background(), core.Config.RedisPrefix+key, value, time.Duration(seconds)*time.Second).Err()
	if err != nil {
		core.Logger.Error("RDB KSet", zap.Any("error", err))
		return false
	}
	return true
}

// KGet 根据Key获取value
func (helper *rdbHelper) KGet(key string) string {
	res, err := RDB.Get(context.Background(), core.Config.RedisPrefix+key).Result()
	if err != nil {
		core.Logger.Error("RDB KGet", zap.Any("error", err))
		return ""
	}
	return res
}

// KDel 删除Keys
func (helper *rdbHelper) KDel(keys []string) bool {
	fullKeys := helper.fullKeys(keys)
	err := RDB.Del(context.Background(), fullKeys...).Err()
	if err != nil {
		core.Logger.Error("RDB KDel", zap.Any("error", err))
		return false
	}
	return true
}

// KExists Keys是否存在
func (helper *rdbHelper) KExists(keys []string) int64 {
	fullKeys := helper.fullKeys(keys)
	res, err := RDB.Exists(context.Background(), fullKeys...).Result()
	if err != nil {
		core.Logger.Error("RDB KExists", zap.Any("error", err))
		return -1
	}
	return res
}

// fullKeys 构建keys
func (helper *rdbHelper) fullKeys(keys []string) []string {
	var fullKeys []string
	for _, key := range keys {
		fullKeys = append(fullKeys, core.Config.RedisPrefix+key)
	}
	return fullKeys
}
