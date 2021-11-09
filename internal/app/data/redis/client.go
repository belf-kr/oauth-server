package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	Client *redis.Client
)

func init() {
	redisHost := viper.GetString("REDIS_HOST")
	redisPort := viper.GetString("REDIS_PORT")
	redisPass := viper.GetString("REDIS_PASSWORD")
	redisDbname := viper.GetInt("REDIS_DATABASE")

	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,        // 접근 url 및 port
		Password: redisPass,   // password "" 값은 없다는 뜻
		DB:       redisDbname, // 0이면 기본 DB 사용
	})

	ctx := context.Background()
	log.Println("redis 연결 중...")
	if _, err := Client.Ping(ctx).Result(); err != nil {
		log.Panic("redis 연결에 실패하였습니다.\n\t" + err.Error())
	}

	log.Printf("redis starting at %s, DB: %d\n", addr, redisDbname)
}
