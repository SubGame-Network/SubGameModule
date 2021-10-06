package redis

import (
	"errors"
	"fmt"
	"time"

	"sync"

	"github.com/SubGame-Network/SubGameModuleService/config"
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

var pool *redis.Pool
var redisOnce sync.Once

type redisUtil struct {
	config *config.Config
	conn   redis.Conn
}

// TODO 有斷線問題，看新套件狀況之後這個可能棄用
func NewRedis(config *config.Config) domain.RedisUtil {
	if pool == nil {
		redisOnce.Do(func() {
			pool = connectionPool(config)
			fmt.Printf("redis connect %v\n", pool)
		})
	}
	return &redisUtil{
		config: config,
		conn:   pool.Get(),
	}
}

func connectionPool(config *config.Config) *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool
		MaxIdle: config.Redis.Max_idle,

		// Maximum number of connections allocated by the pool at a given timeutil.
		MaxActive: config.Redis.Max_active,

		// Close connections after remaining idle for this duration.
		IdleTimeout: time.Duration(config.Redis.Idle_timeout) * time.Millisecond,

		// Dial is an application supplied function for creating and configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
				redis.DialDatabase(config.Redis.Database),
				redis.DialPassword(config.Redis.Auth),
			)
			if err != nil {
				zap.S().Warn(err)
				return nil, err
			}
			return c, nil
		},

		// PING PONG test
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

// Close redis connection
func (r *redisUtil) Close() {
	r.conn.Close()
}

// Exists returns if key exists
func (r *redisUtil) Exists(key string) (bool, error) {
	return redis.Bool(r.conn.Do("EXISTS", key))
}

/**
 * Set expire time
 * key = your key
 * ttl = expire time seconds
 */
func (r *redisUtil) Expire(key string, ttl int) error {
	_, err := r.conn.Do("EXPIRE", key, ttl)

	return err
}

// SetEx sets a key-value pair with expire timeutil
func (r *redisUtil) SetEx(key string, value interface{}, ttl int) error {
	_, err := r.conn.Do("SETEX", key, ttl, value)
	return err
}

func (r *redisUtil) Set(key string, value interface{}) error {
	_, err := r.conn.Do("SET", key, value)
	return err
}

// Del removes the specified keys. A key is ignored if it does not exist.
// Delete multiple keys should use []string for parameter {key}
func (r *redisUtil) Del(key interface{}) (err error) {
	switch key.(type) {
	case []string:
		_, err = r.conn.Do("DEL", redis.Args{}.AddFlat(key)...)
	default:
		_, err = r.conn.Do("DEL", key)
	}
	return
}

// Get gets value of given key
func (r *redisUtil) Get(key string) (string, error) {
	return redis.String(r.conn.Do("GET", key))
}

// Keys gets value of given key
func (r *redisUtil) Keys(key string) ([]string, error) {
	return redis.Strings(r.conn.Do("KEYS", key))
}

// MGet Returns the values of all specified keys. For every key that
// does not hold a string value or does not exist, the special value nil is returned.
// Because of this, the operation never fails.
func (r *redisUtil) MGet(key []string) ([]string, error) {
	var result []string

	value, err := r.conn.Do("MGET", redis.Args{}.AddFlat(key)...)

	if err != nil {
		return nil, err
	}

	switch value := value.(type) {
	case []interface{}:
		result = make([]string, len(value))
		for i := 0; i < len(value); i++ {
			switch v := value[i].(type) {
			case []uint8:
				result[i] = string(v)
			}
		}
	default:
		return nil, errors.New("unsupported type from redis MGET command")
	}

	return result, nil
}

// list
func (r *redisUtil) LPush(key string, value interface{}) error {
	_, err := r.conn.Do("LPUSH", key, value)
	return err
}

// HSet sets field in the hash stored at key to value
func (r *redisUtil) HSet(key string, field string, value interface{}) error {
	_, err := r.conn.Do("HSET", key, field, value)
	return err
}

// HGet gets value of a specific field of key
func (r *redisUtil) HGet(key string, field string) (string, error) {
	return redis.String(r.conn.Do("HGET", key, field))
}

// HDel Removes the specified fields from the hash stored at key
func (r *redisUtil) HDel(key string, fieldValue interface{}) (err error) {
	switch fieldValue.(type) {
	case []string, []int, []int32, []int64, []float32, []float64, []interface{}:
		_, err = r.conn.Do("HDEL", redis.Args{}.Add(key).AddFlat(fieldValue)...)
	default:
		_, err = r.conn.Do("HDEL", key, fieldValue)
	}
	return
}

// HMSet sets multiple fields for a specific key
func (r *redisUtil) HMSet(key string, fieldValue interface{}) error {
	_, err := r.conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(fieldValue)...)
	return err
}

//HMGet returns the values associated with the specified fields in the hash stored at key.
func (r *redisUtil) HMGet(key string, fields interface{}) ([]string, error) {
	var result []string

	value, err := r.conn.Do("HMGET", redis.Args{}.Add(key).AddFlat(fields)...)

	if err != nil {
		return nil, err
	}

	switch value := value.(type) {
	case []interface{}:
		result = make([]string, len(value))
		for i := 0; i < len(value); i++ {
			result[i] = string(value[i].([]uint8))
		}
	default:
		return nil, err
	}

	return result, nil
}

// HGetAll get all fields for a given key
func (r *redisUtil) HGetAll(key string) (map[string]string, error) {
	var result map[string]string

	value, err := r.conn.Do("HGETALL", key)

	if err != nil {
		return nil, err
	}

	switch value := value.(type) {
	case []interface{}:
		result = make(map[string]string, len(value))
		for i := 0; i < len(value); i += 2 {
			result[string(value[i].([]uint8))] = string(value[i+1].([]uint8))
		}
	}

	return result, nil
}

// SAdd add member to a set called key.  If key does not exist, a new set is created.
// param 'member' support array of string, int, int32, int64, float32, float64, interface{}
func (r *redisUtil) SAdd(key string, member interface{}) (err error) {
	switch member.(type) {
	case []string, []int, []int32, []int64, []float32, []float64, []interface{}:
		_, err = r.conn.Do("SADD", redis.Args{}.Add(key).AddFlat(member)...)
	default:
		_, err = r.conn.Do("SADD", key, member)
	}

	return
}

// SRem removes the specified members from the set stored at key.
// param 'member' support array of string, int, int32, int64, float32, float64, interface{}
func (r *redisUtil) SRem(key string, member interface{}) (err error) {
	switch member.(type) {
	case []string, []int, []int32, []int64, []float32, []float64, []interface{}:
		_, err = r.conn.Do("SREM", redis.Args{}.Add(key).AddFlat(member)...)
	default:
		_, err = r.conn.Do("SREM", key, member)
	}

	return
}

// SIsMember Returns if member is a member of the set stored at key.
func (r *redisUtil) SIsMember(key string, member interface{}) (bool, error) {
	exist, err := r.conn.Do("SISMEMBER", key, member)

	if exist.(int64) == 1 {
		return true, nil
	}
	return false, err
}

// SMembers get all members from the set called key
func (r *redisUtil) SMembers(key string) ([]string, error) {
	var result []string

	value, err := r.conn.Do("SMEMBERS", key)

	if err != nil {
		return nil, err
	}

	switch value := value.(type) {
	case []interface{}:
		result = make([]string, len(value))
		for i := 0; i < len(value); i++ {
			result[i] = string(value[i].([]uint8))
		}
	}

	return result, nil
}

func (r *redisUtil) INCR(key string) error {
	_, err := r.conn.Do("INCR", key)
	if err != nil {
		return err
	}
	return nil
}

func (r *redisUtil) ZADD(key string, score int64, val string) error {
	_, err := r.conn.Do("zadd", key, score, val)
	if err != nil {
		return err
	}
	return nil
}

func (r *redisUtil) ZRANGEBYSCORE(key string, scoreEnd int64) ([]string, error) {
	return redis.Strings(r.conn.Do("zrangebyscore", key, 0, scoreEnd))
}

func (r *redisUtil) ZREM(key string, val string) error {
	_, err := r.conn.Do("zrem", key, val)
	if err != nil {
		return err
	}
	return nil
}
