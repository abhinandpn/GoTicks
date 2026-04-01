package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Increment the counter
	val, err := rdb.Incr(ctx, "pageviews").Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pageviews: %d\n", val) // Pageviews: 1 (then 2, 3, ...)
}
