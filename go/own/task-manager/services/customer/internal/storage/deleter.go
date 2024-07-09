package storage

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