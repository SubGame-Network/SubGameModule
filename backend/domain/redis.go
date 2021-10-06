package domain

type RedisUtil interface {
	Close()
	Expire(key string, ttl int) error
	SetEx(key string, value interface{}, ttl int) error
	Set(key string, value interface{}) error
	Del(key interface{}) (err error)
	Get(key string) (string, error)
	Keys(key string) ([]string, error)
	LPush(key string, value interface{}) error
	MGet(key []string) ([]string, error)
	HSet(key string, field string, value interface{}) error
	HGet(key string, field string) (string, error)
	HDel(key string, fieldValue interface{}) (err error)
	HMSet(key string, fieldValue interface{}) error
	HMGet(key string, fields interface{}) ([]string, error)
	HGetAll(key string) (map[string]string, error)
	SAdd(key string, member interface{}) (err error)
	SRem(key string, member interface{}) (err error)
	SIsMember(key string, member interface{}) (bool, error)
	SMembers(key string) ([]string, error)
	INCR(key string) error
	ZADD(key string, score int64, val string) error
	ZRANGEBYSCORE(key string, scoreEnd int64) ([]string, error)
	ZREM(key string, val string) error
}
