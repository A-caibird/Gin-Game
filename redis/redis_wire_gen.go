// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package redis2

import (
	"github.com/redis/go-redis/v9"
)

// Injectors from wire.go:

// InitRedis init redis client connection
func InitRedis() *redis.Client {
	client := NewRedisClient()
	return client
}
