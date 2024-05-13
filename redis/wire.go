//go:build wireinject
// +build wireinject

package redis2

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

// InitRedis init redis client connection
func InitRedis() *redis.Client {
	wire.Build(NewRedisClient)
	return nil
}
