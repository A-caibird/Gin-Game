package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"testing"
	"time"
)

func TestInitRedis(t *testing.T) {
	rdb := InitRedis()
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil && err == redis.Nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
	fmt.Printf("%#v\n", val)
}
