package main

import (
	"context"
	"fmt"
	"time"

	"mem-cache/cache"
	"mem-cache/model"
)

func main() {
	ctx := context.TODO()
	ttl := 2 * time.Second
	cache := cache.NewCache(ctx, ttl)
	profile := model.Profile{
		UUID: "id1",
		Name: "name1",
		Orders: []*model.Order{
			{
				UUID:      "id1",
				Value:     "value1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UUID:      "id1",
				Value:     "value2",
				CreatedAt: time.Now().Add(-5 * time.Minute),
				UpdatedAt: time.Now().Add(-3 * time.Minute),
			},
		},
	}
	cache.Set("key1", profile)
	profile_read, err := cache.Get("key1")
	if err != nil {
		fmt.Printf("Error on read occured: %v", err)
		return
	}
	fmt.Printf("Read profile is %v", profile_read)
}
