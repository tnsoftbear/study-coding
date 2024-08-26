package main

import (
	"fmt"
	"generic-cache/cache"
	"generic-cache/model"
	"log"
)

func main() {
	order1 := model.Order{
		ID:    "id1",
		Value: "value1",
	}
	order2 := model.Order{
		ID:    "id2",
		Value: "value2",
	}

	profile := model.Profile{
		ID:     "id1",
		Orders: []*model.Order{&order1, &order2},
	}

	cache := cache.NewCache()
	cache.Set("profile1", profile.Clone())
	cache.Set("order2", order2.Clone())
	profileCacheValue := cache.Get("profile1")
	orderCacheValue := cache.Get("order2")

	orderFromCache, ok := orderCacheValue.(model.Order)
	if !ok {
		log.Fatal("Type assertion error for Order")
	}

	profileFromCache, ok := profileCacheValue.(model.Profile)
	if !ok {
		log.Fatal("Type assertion error for Profile")
	}

	profileStr := profileToString(&profileFromCache)
	orderStr := orderToString(&orderFromCache)
	fmt.Printf("profileStr: %s, orderStr: %s", profileStr, orderStr)
}

func profileToString(profile *model.Profile) string {
	return fmt.Sprintf("profile: %+v", profile)
}

func orderToString(order *model.Order) string {
	return fmt.Sprintf("order: %+v", order)
}
