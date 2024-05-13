package redis2

import (
	"Game/tools"
	"github.com/redis/go-redis/v9"
)

var (
	config = redis.Options{
		Addr:           tools.Conf.RedisServer.Host + ":" + tools.Conf.RedisServer.Port,
		Password:       tools.Conf.RedisServer.Password,
		DB:             tools.Conf.RedisServer.Database, // Selecting the database to be used
		MaxActiveConns: tools.Conf.RedisServer.MaxActiveConns,
		MaxIdleConns:   tools.Conf.RedisServer.MaxIdleConns,
	}
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&config)
}
