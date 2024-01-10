package redis

import (
	"GinTemplate/config"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(config.Conf.ReadRedisOptions())
	Rc = &RedisCache{rdb: rdb}
}

func (rc *RedisCache) Set(ctx context.Context, key, value string, expire time.Duration) error {
	err := rc.rdb.Set(ctx, key, value, expire).Err()
	return err
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.rdb.Get(ctx, key).Result()
	return result, err
}

func (rc *RedisCache) HSet(ctx context.Context, key string, field string, value string) error {
	err := rc.rdb.HSet(ctx, key, field, value).Err()
	return err
}

func (rc *RedisCache) HGet(ctx context.Context, key string, field string) (string, error) {
	res, err := rc.rdb.HGet(ctx, key, field).Result()
	return res, err
}

func (rc *RedisCache) HMGet(ctx context.Context, key string, field ...string) ([]interface{}, error) {
	res, err := rc.rdb.HMGet(ctx, key, field...).Result()
	return res, err
}

func (rc *RedisCache) HDel(ctx context.Context, key string, field string) error {
	err := rc.rdb.HDel(ctx, key, field).Err()
	return err
}

func (rc *RedisCache) HIncrBy(ctx context.Context, key string, field string, incr int64) error {
	err := rc.rdb.HIncrBy(ctx, key, field, incr).Err()
	return err
}

func (rc *RedisCache) HDecrBy(ctx context.Context, key string, field string, decr int64) error {
	err := rc.rdb.HIncrBy(ctx, key, field, -decr).Err()
	return err
}

func (rc *RedisCache) HKeys(ctx context.Context, key string) ([]string, error) {
	result, err := rc.rdb.HKeys(ctx, key).Result()
	return result, err
}

func (rc *RedisCache) ZAdd(ctx context.Context, key string, value string, score float64) error {
	err := rc.rdb.ZAdd(ctx, key, &redis.Z{Member: value, Score: score}).Err()
	return err
}

func (rc *RedisCache) ZIncrBy(ctx context.Context, key string, value string, incr float64) error {
	err := rc.rdb.ZIncrBy(ctx, key, incr, value).Err()
	return err
}

func (rc *RedisCache) ZRem(ctx context.Context, key string, member ...string) error {
	err := rc.rdb.ZRem(ctx, key, member).Err()
	return err
}

func (rc *RedisCache) ZScore(ctx context.Context, key string, member string) (float64, error) {
	result, err := rc.rdb.ZScore(ctx, key, member).Result()
	return result, err
}

func (rc *RedisCache) ZCard(ctx context.Context, key string) (int64, error) {
	return rc.rdb.ZCard(ctx, key).Result()
}

func (rc *RedisCache) ZRange(ctx context.Context, key string, start, end int64) ([]string, error) {
	result, err := rc.rdb.ZRange(ctx, key, start, end).Result()
	return result, err
}
func (rc *RedisCache) ZRevRange(ctx context.Context, key string, start, end int64) ([]string, error) {
	result, err := rc.rdb.ZRevRange(ctx, key, start, end).Result()
	return result, err
}

func (rc *RedisCache) ZRevRank(ctx context.Context, key string, member string) (int64, error) {
	result, err := rc.rdb.ZRevRank(ctx, key, member).Result()
	return result, err
}

func (rc *RedisCache) ZRevRangeWithScores(ctx context.Context, key string, start int64, stop int64) ([]redis.Z, error) {
	result, err := rc.rdb.ZRevRangeWithScores(ctx, key, start, stop).Result()
	return result, err
}

func (rc *RedisCache) ZRevRangeTopN(ctx context.Context, key string, count int64) ([]redis.Z, error) {
	result, err := rc.rdb.ZRevRangeWithScores(ctx, key, 0, count-1).Result()
	return result, err
}

func (rc *RedisCache) PFAdd(ctx context.Context, key string, value interface{}) error {
	err := rc.rdb.PFAdd(ctx, key, value).Err()
	return err
}

func (rc *RedisCache) PFCount(ctx context.Context, key string) (int64, error) {
	result, err := rc.rdb.PFCount(ctx, key).Result()
	return result, err
}

func (rc *RedisCache) SetBit(ctx context.Context, key string, offset int64, value int) error {
	_, err := rc.rdb.SetBit(ctx, key, offset, value).Result()
	return err
}

func (rc *RedisCache) GetBit(ctx context.Context, key string, memId int64) (int64, error) {
	result, err := rc.rdb.GetBit(ctx, key, memId).Result()
	return result, err
}

func (rc *RedisCache) BitCount(ctx context.Context, key string) (int64, error) {
	result, err := rc.rdb.BitCount(ctx, key, &redis.BitCount{Start: 0, End: -1}).Result()
	return result, err
}

func (rc *RedisCache) Del(ctx context.Context, key string) error {
	err := rc.rdb.Del(ctx, key).Err()
	return err
}

func (rc *RedisCache) Eval(ctx context.Context, script string, keys []string, values ...interface{}) (interface{}, error) {
	result, err := rc.rdb.Eval(ctx, script, keys, values).Result()
	return result, err
}

func (rc *RedisCache) Pipeline() redis.Pipeliner {
	return rc.rdb.Pipeline()
}

func (rc *RedisCache) ErrNil() error {
	return redis.Nil
}
