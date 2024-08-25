package model

import (
	"fmt"
	"generic-cache/cache"
	"log"
)

type Profile struct {
	ID     string   `json:"uuid"`
	Orders []*Order `json:"orders"`
}

type Order struct {
	ID    string `json:"uuid"`
	Value any    `json:"value"`
}

func (p Profile) CloneToCacheValue() cache.CacheValue {
	ordersCopy := make([]*Order, len(p.Orders))
	for idx, order := range p.Orders {
		orderClone, ok := order.CloneToCacheValue().(Order)
		if !ok {
			log.Fatal("Error on type assignment")
		}
		ordersCopy[idx] = &orderClone
	}
	return Profile{
		ID:     p.ID,
		Orders: ordersCopy,
	}
}

func (p Profile) String() string {
	orderStrings := make([]string, len(p.Orders))
	for _, order := range p.Orders {
		orderStrings = append(orderStrings, order.String())
	}
	return fmt.Sprintf("ID: %s, Orders: %s", p.ID, orderStrings)
}

func (o Order) CloneToCacheValue() cache.CacheValue {
	return Order{
		ID:    o.ID,
		Value: o.Value,
	}
}

func (o Order) String() string {
	return fmt.Sprintf("ID: %s, Value: %v", o.ID, o.Value)
}
