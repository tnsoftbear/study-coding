package storage

import (
	"strconv"
	"task_manager/internal/model"
)

func (rm *RedisManager) LoadCustomerShippingById(id string) (*model.CustomerShipping, error) {
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
	shipping := model.CustomerShipping{
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
			address := model.CustomerAddress{
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

func (rm *RedisManager) LoadAllCustomerShippings() ([]*model.CustomerShipping, error) {
	var shippings []*model.CustomerShipping = make([]*model.CustomerShipping, 0)
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
