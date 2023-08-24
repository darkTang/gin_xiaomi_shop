package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}
}
