package org.example.trading_demo.service;

import org.example.trading_demo.model.Order;
import org.springframework.stereotype.Service;

@Service
public class ExchangeService {
    public Order exchange(Order buyOrder, Order sellOrder){
        Order result = new Order();
        result.quantity = Math.min(buyOrder.quantity, sellOrder.quantity);
        result.price = Math.min(buyOrder.price, sellOrder.price);
        return result;
    }
}
