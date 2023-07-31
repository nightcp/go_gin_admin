package redis

import (
	"admin/core"
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// SSAdd 向有序集合添加一个或多个成员，或者更新已存在成员的分数
func (helper *rdbHelper) SSAdd(key string, members []SSetMember) bool {
	ssMembers := helper.makeSSMembers(members)
	err := RDB.ZAdd(context.Background(), core.Config.RedisPrefix+key, ssMembers...).Err()
	if err != nil {
		core.Logger.Error("RDB SSAdd", zap.Any("error", err))
		return false
	}
	return true
}

// SSGetMembersByScore 返回有序集合中指定分数区间的成员列表
func (helper *rdbHelper) SSGetMembersByScore(key string, min string, max string, offset int64, count int64) []string {
	zRangeBy := redis.ZRangeBy{
		Min: min,
		Max: max,
	}
	if offset > 0 {
		zRangeBy.Offset = offset
	}
	if count > 0 {
		zRangeBy.Count = count
	}
	res, err := RDB.ZRangeByScore(context.Background(), core.Config.RedisPrefix+key, &zRangeBy).Result()
	if err != nil {
		core.Logger.Error("RDB SSGetByScore", zap.Any("error", err))
		return []string{}
	}
	return res
}

// SSGetScore 返回有序集中，成员的分数值
func (helper *rdbHelper) SSGetScore(key string, member string) float64 {
	res, err := RDB.ZScore(context.Background(), core.Config.RedisPrefix+key, member).Result()
	if err != nil {
		return -1
	}
	return res
}

// SSDel 移除有序集中的一个或多个成员，不存在的成员将被忽略
func (helper *rdbHelper) SSDel(key string, members []string) bool {
	err := RDB.ZRem(context.Background(), core.Config.RedisPrefix+key, members).Err()
	if err != nil {
		core.Logger.Error("RDB SSDel", zap.Any("error", err))
		return false
	}
	return true
}

// makeSSMembers 构建有序集合Member
func (helper *rdbHelper) makeSSMembers(members []SSetMember) []redis.Z {
	var ssMembers []redis.Z
	for _, member := range members {
		ssMembers = append(ssMembers, redis.Z{
			Score:  member.Score,
			Member: member.Member,
		})
	}
	return ssMembers
}
