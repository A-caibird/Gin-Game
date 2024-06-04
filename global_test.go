package main

import (
	"Game/handler"
	"Game/redis"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	rdb := redis2.InitRedis()
	key := "a"
	val := "b"
	expri := 2 * time.Second
	if err := rdb.Set(context.Background(), key, val, expri); err != nil {
		fmt.Println("ok")
	}
	time.Sleep(3 * time.Second)
	if _, err := rdb.Get(context.Background(), key).Result(); errors.Is(err, redis.Nil) {
		fmt.Println("no non o")
	}
}

func Test2(t *testing.T) {
	rdb := redis2.InitRedis()
	key := "2"
	val := "202020"
	if err := rdb.Set(context.Background(), key, val, 5*60*time.Second); err != nil {
		fmt.Println("ok", err)
	}
	if val, err := rdb.Get(context.Background(), key).Result(); errors.Is(err, redis.Nil) {
		fmt.Println("no non o")
	} else {
		fmt.Printf("%#v", val)
	}
}

func Test3(t *testing.T) {
	res, err := handler.NotifyFriend("16608278954")
	fmt.Printf("%#v %#v", res, err)
}
