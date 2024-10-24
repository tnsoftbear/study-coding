package model

import (
	"mem-cache/cache"
	"time"
)

type Profile struct {
	UUID   string   `json:"uuid"`
	Name   string   `json:"name"`
	Orders []*Order `json:"orders"`
}

type Order struct {
	UUID      string    `json:"uuid"`
	Value     any       `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p Profile) Clone() Profile {
	ordersCopy := make([]*Order, len(p.Orders))
	for idx, order := range p.Orders {
		orderCopy := order.Clone()
		ordersCopy[idx] = &orderCopy
	}
	return Profile{
		UUID:   p.UUID,
		Name:   p.Name,
		Orders: ordersCopy,
	}
}

func (p Profile) CloneToCacheValue() cache.CacheValue {
	profileCV := p.Clone()
	return profileCV
}

func (o Order) Clone() Order {
	return Order{
		UUID:      o.UUID,
		Value:     o.Value,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}

func (o Order) CloneToCacheValue() cache.CacheValue {
	orderCV := o.Clone()
	return orderCV
}
