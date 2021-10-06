package redis

import (
	"fmt"

	"context"
	"time"

	"github.com/SubGame-Network/SubGameModuleService/config"
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/go-redis/redis/v8"
)

type GoRedis struct {
	config *config.Config
	conn   *redis.Client
}

func NewGoRedis(config *config.Config) domain.GoRedis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Auth,
		DB:       config.Redis.Database,
	})

	return &GoRedis{
		config: config,
		conn:   rdb,
	}
}

func (r *GoRedis) Set(key string, value interface{}) error {
	return r.conn.Set(context.Background(), key, value, 0).Err()
}

func (r *GoRedis) Del(key string) error {
	return r.conn.Del(context.Background(), key).Err()
}

func (r *GoRedis) Get(key string) (string, error) {
	return r.conn.Get(context.Background(), key).Result()
}

func (r *GoRedis) Keys(key string) ([]string, error) {
	return r.conn.Keys(context.Background(), key).Result()
}

func (r *GoRedis) INCR(key string) error {
	return r.conn.Incr(context.Background(), key).Err()
}

func (r *GoRedis) Expire(key string, ttl int) error {
	return r.conn.Expire(context.Background(), key, time.Duration(ttl)*time.Second).Err()
}

func (r *GoRedis) SetEx(key string, value interface{}, ttl int) error {
	return r.conn.SetEX(context.Background(), key, value, time.Duration(ttl)*time.Second).Err()
}
