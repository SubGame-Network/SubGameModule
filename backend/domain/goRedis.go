//go:generate ~/go/bin/mockgen -source ./goRedis.go -destination ../mock/redis_mock.go -package domain_mock

package domain

type GoRedis interface {
	Set(key string, value interface{}) error
	Del(key string) error
	Get(key string) (string, error)
	Keys(key string) ([]string, error)
	INCR(key string) error
	Expire(key string, ttl int) error
	SetEx(key string, value interface{}, ttl int) error
}
