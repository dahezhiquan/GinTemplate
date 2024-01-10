package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key, value string, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	HSet(ctx context.Context, key string, field string, value string) error
	HGet(ctx context.Context, key string, field string) (string, error)
	HMGet(ctx context.Context, key string, field ...string) ([]interface{}, error)
	HDel(ctx context.Context, key string, field string) error
	HIncrBy(ctx context.Context, key string, field string, incr int64) error
	HDecrBy(ctx context.Context, key string, field string, decr int64) error
	HKeys(ctx context.Context, key string) ([]string, error)
	ZAdd(ctx context.Context, key string, value string, score float64) error
	ZIncrBy(ctx context.Context, key string, value string, incr float64) error
	ZRem(ctx context.Context, key string, member ...string) error
	ZScore(ctx context.Context, key string, member string) (float64, error)
	ZCard(ctx context.Context, key string) (int64, error)
	ZRange(ctx context.Context, key string, start, end int64) ([]string, error)
	ZRevRange(ctx context.Context, key string, start, end int64) ([]string, error)
	ZRevRank(ctx context.Context, key string, member string) (int64, error)
	ZRevRangeWithScores(ctx context.Context, key string, start int64, stop int64) ([]redis.Z, error)
	ZRevRangeTopN(ctx context.Context, key string, count int64) ([]redis.Z, error)
	SetBit(ctx context.Context, key string, offset int64, value int) error
	GetBit(ctx context.Context, key string, memId int64) (int64, error)
	BitCount(ctx context.Context, key string) (int64, error)
	PFAdd(ctx context.Context, key string, value interface{}) error
	PFCount(ctx context.Context, key string) (int64, error)
	Del(ctx context.Context, key string) error
	Eval(ctx context.Context, script string, key []string, value ...interface{}) (interface{}, error)
	Pipeline() redis.Pipeliner
	ErrNil() error
}
