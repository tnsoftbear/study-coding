package storage

import (
	"task_manager/internal/model"

	"github.com/go-redis/redis"
)

func (rm *RedisManager) SaveCustomerShipping(shipping model.CustomerShipping) error {
	// TODO: add transaction "multi", consider rollback
	hsetShippingId := makeCustomerShippingId(shipping.Id)
	hset := rm.rdb.HSet(hsetShippingId, "Id", shipping.Id)
	if hset.Err() != nil {
		return hset.Err()
	}

	hset = rm.rdb.HSet(hsetShippingId, "Name", shipping.Name)
	if hset.Err() != nil {
		return hset.Err()
	}

	hset = rm.rdb.HSet(hsetShippingId, "CreatedAt", shipping.CreatedAt)
	if hset.Err() != nil {
		return hset.Err()
	}

	if shipping.Address != nil {
		address := shipping.Address
		hsetAddressId := makeCustomerShippingAddressId(address.Id)
		if hset := rm.rdb.HSet(hsetAddressId, "Id", address.Id); hset.Err() != nil {
			return hset.Err()
		}
		if hset := rm.rdb.HSet(hsetAddressId, "City", address.City); hset.Err() != nil {
			return hset.Err()
		}
		if hset := rm.rdb.HSet(hsetAddressId, "Street", address.Street); hset.Err() != nil {
			return hset.Err()
		}
		if hset := rm.rdb.HSet(hsetAddressId, "Longitude", address.Longitude); hset.Err() != nil {
			return hset.Err()
		}
		if hset := rm.rdb.HSet(hsetAddressId, "Latitude", address.Latitude); hset.Err() != nil {
			return hset.Err()
		}
		// Set FK reference for `customer_shipping`
		if hset := rm.rdb.HSet(hsetShippingId, "AddressId", address.Id); hset.Err() != nil {
			return hset.Err()
		}
	}

	z := redis.Z{Score: float64(shipping.CreatedAt), Member: shipping.Id}
	zadd := rm.rdb.ZAdd("customer_shippings", z)
	if zadd.Err() != nil {
		return zadd.Err()
	}

	return nil
}
