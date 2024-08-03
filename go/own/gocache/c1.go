package main

import (
	"context"
	"fmt"

	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/marshaler"
	redis_store "github.com/eko/gocache/store/redis/v4"
	redis "github.com/redis/go-redis/v9"
)

type Book struct {
	Id int
	Name string
}

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	redisStore := redis_store.NewRedis(rdb)

	cacheManager := cache.New[any](redisStore)

	marshal := marshaler.New(cacheManager)
	key := "my-key"
	value := Book{Id: 1111, Name: "Book name"}

	//err := cacheManager.Set(ctx, "my-key", "my-value", store.WithExpiration(15*time.Second))
	err := marshal.Set(ctx, key, value)
	if err != nil {
		panic(err)
	}

	// value, err := cacheManager.Get(ctx, "my-key")
	book, err := marshal.Get(ctx, key, new(Book))
	switch err {
	case nil:
		fmt.Printf("Get the key '%s' from the redis cache. Result: %v", "my-key", book)
	case redis.Nil:
		fmt.Printf("Failed to find the key '%s' from the redis cache.", "my-key")
	default:
		fmt.Printf("Failed to get the value from the redis cache with key '%s': %v", "my-key", err)
	}
}