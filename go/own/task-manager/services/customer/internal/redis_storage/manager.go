package redis_storage

import (
	"fmt"
	"strconv"
	"task_manager/internal/config"
	"task_manager/internal/types"

	"github.com/go-redis/redis"
)

type RedisManager struct {
	rdb *redis.Client
}

func NewRedisManager() *RedisManager {
	addr := fmt.Sprintf("%s:%s",
		config.GetStrEnv("REDIS_HOST", "localhost"),
		config.GetStrEnv("REDIS_PORT", "6379"))
	return &RedisManager{
		rdb: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: config.GetStrEnv("REDIS_PASSWORD", ""),
			DB:       config.GetIntEnv("REDIS_DB", 0),
		}),
	}
}

func (rm *RedisManager) LoadCustomerShippingById(id string) (*types.CustomerShipping, error) {
	shippingHGetAll := rm.rdb.HGetAll(makeCustomerShippingId(id))
	if err := shippingHGetAll.Err(); err != nil {
		return nil, err
	}
	shippingRes, err := shippingHGetAll.Result()
	if err != nil {
		return nil, err
	}
	if len(shippingRes) == 0 {
		return nil, nil
	}

	createdAt, _ := strconv.ParseInt(shippingRes["CreatedAt"], 10, 64)
	shipping := types.CustomerShipping{
		Id:        shippingRes["Id"],
		Name:      shippingRes["Name"],
		Address:   nil,
		CreatedAt: createdAt,
	}

	if addressId, is := shippingRes["AddressId"]; is {
		addressHGetAll := rm.rdb.HGetAll(makeCustomerShippingAddressId(addressId))
		if err := addressHGetAll.Err(); err != nil {
			return nil, err
		}
		addressRes, err := addressHGetAll.Result()
		if err != nil {
			return nil, err
		}
		if len(addressRes) > 0 {
			longitude, err := strconv.ParseFloat(addressRes["Longitude"], 64)
			if err != nil {
				return nil, err
			}
			latitude, err := strconv.ParseFloat(addressRes["Latitude"], 64)
			if err != nil {
				return nil, err
			}
			address := types.CustomerAddress{
				Id:        addressRes["Id"],
				City:      addressRes["City"],
				Street:    addressRes["Street"],
				Longitude: longitude,
				Latitude:  latitude,
			}
			shipping.Address = &address
		}
	}

	return &shipping, nil
}

func (rm *RedisManager) LoadAllCustomerShippings() ([]*types.CustomerShipping, error) {
	var shippings []*types.CustomerShipping = make([]*types.CustomerShipping, 0)
	zRange := rm.rdb.ZRange("customer_shippings", 0, -1)
	if err := zRange.Err(); err != nil {
		return nil, err
	}

	ids, err := zRange.Result()
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		if shipping, err := rm.LoadCustomerShippingById(id); err != nil {
			return nil, err
		} else {
			shippings = append(shippings, shipping)
		}
	}
	return shippings, nil
}

func (rm *RedisManager) SaveCustomerShipping(shipping types.CustomerShipping) error {
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

func (rm *RedisManager) DeleteCustomerShipping(id string) error {
	shipping, err := rm.LoadCustomerShippingById(id)
	if err != nil {
		return err
	}

	if shipping.Address != nil {
		if err := rm.rdb.Unlink(makeCustomerShippingAddressId(shipping.Address.Id)).Err(); err != nil {
			return err
		}
	}

	if err := rm.rdb.Unlink(makeCustomerShippingId(id)).Err(); err != nil {
		return err
	}

	if err := rm.rdb.ZRem("customer_shippings", id).Err(); err != nil {
		return err
	}

	return nil
}

func makeCustomerShippingId(id string) string {
	return fmt.Sprintf("customer_shipping:%s", id)
}

func makeCustomerShippingAddressId(id string) string {
	return fmt.Sprintf("customer_shipping_address:%s", id)
}
