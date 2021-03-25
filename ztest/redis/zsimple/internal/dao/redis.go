package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
)

func NewRedis() (r *redis.Redis, cf func(), err error) {
	var cfg struct {
		Client *redis.Config
	}
	if err = paladin.Get("redis.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	r = redis.NewRedis(cfg.Client)
	return
}

func (d *dao) RedisDel(ctx context.Context, key string) (err error) {
	_, err = d.redis.Do(ctx, "DEL", key)
	return
}

func (d *dao) RedisIsExist(ctx context.Context, key string) (isExist bool, err error) {
	return redis.Bool(d.redis.Do(ctx, "EXISTS", key))
}

func (d *dao) RedisLRem(ctx context.Context, key string, start int32, end string) (err error) {
	_, err = d.redis.Do(ctx, "LREM", key, start, end)
	return
}

func (d *dao) RedisLRange(ctx context.Context, key string, start, end int32) (ret []string, err error) {
	return redis.Strings(d.redis.Do(ctx, "LRange", key, start, end))
}

func (d *dao) RedisLTrim(ctx context.Context, key string, start, end int32) (err error) {
	_, err = d.redis.Do(ctx, "LTRIM", key, start, end)
	return
}

func (d *dao) RedisLPush(ctx context.Context, key, value string) (err error) {
	_, err = d.redis.Do(ctx, "LPUSH", key, value)
	return
}

func (d *dao) RedisHKeys(ctx context.Context, key string) (ret []string, err error) {
	return redis.Strings(d.redis.Do(ctx, "HKEYS", key))
}

func (d *dao) RedisHExist(ctx context.Context, key string, filed int64) (isExist bool, err error) {
	return redis.Bool(d.redis.Do(ctx, "HEXISTS", key, filed))
}

func (d *dao) RedisHSet(ctx context.Context, key, file, value string) (err error) {
	_, err = d.redis.Do(ctx, "HSET", key, file, value)
	if err != nil {
		return err
	}
	return
}

func (d *dao) RedisHGet(ctx context.Context, key, file string) (ret string, err error) {
	return redis.String(d.redis.Do(ctx, "HGET", key, file))
}

func (d *dao) RedisHMSet(ctx context.Context, key string, m map[string]interface{}) (err error) {
	value := []interface{}{key}
	for k, v := range m {
		value = append(value, k)
		value = append(value, v)
	}
	_, err = d.redis.Do(ctx, "HMSET", value...)
	return
}

func (d *dao) RedisHGetAll(ctx context.Context, key string) (ret []string, err error) {
	return redis.Strings(d.redis.Do(ctx, "HGETALL", key))
}

func (d *dao) RedisHIncr(ctx context.Context, key, file, value string) (err error) {
	_, err = d.redis.Do(ctx, "HINCRBY", key, file, value)
	return
}

func (d *dao) RedisHDel(ctx context.Context, key, file string) (err error) {
	_, err = d.redis.Do(ctx, "HDEL", key, file)
	return
}

func (d *dao) RedisHLen(ctx context.Context, key string) (cnt int64, err error) {
	cnt, err = redis.Int64(d.redis.Do(ctx, "HLEN", key))
	return
}

func (d *dao) RedisHGetAllStringMap(ctx context.Context, key string) (ret map[string]string, err error) {
	ret, err = redis.StringMap(d.redis.Do(ctx, "HGETALL", key))
	return
}

func (d *dao) RedisPublishServerMsg(ctx context.Context, serverChan string, msg []byte) (err error) {
	_, err = d.redis.Do(ctx, "PUBLISH", serverChan, msg)
	if err != nil {
		log.Error("redis publish information failed,err:%v", err)
	}
	return
}

func (d *dao) RedisGetLock(ctx context.Context, key, value string, ExpireTime int64) (GetLock bool, err error) {
	var sign string
	if sign, err = redis.String(d.redis.Do(ctx, "SET", key, value, "NX", "PX", ExpireTime)); err != nil {
		return
	}
	if sign != "OK" {
		return
	}
	GetLock = true
	return
}

func (d *dao) RedisPipeline() (p redis.Pipeliner) {
	p = d.redis.Pipeline()
	return
}