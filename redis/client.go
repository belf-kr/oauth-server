package redis

import (
	"github.com/go-redis/redis/v8"
)

var (
	Client *redis.Client
)

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 접근 url 및 port
		Password: "",               // password "" 값은 없다는 뜻
		DB:       0,                // 기본 DB 사용
	})
}
