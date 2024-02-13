package org.example.trading_demo.service;

import org.example.trading_demo.model.Order;
import org.springframework.stereotype.Service;

@Service
final public class ExchangeService {
    public Order exchange(Order buyOrder, Order sellOrder){
        Order result = new Order();
        result.quantity = Math.min(buyOrder.quantity, sellOrder.quantity);
        result.price = Math.min(buyOrder.price, sellOrder.price);
        return result;
    }

    public String validate(Order buyOrder, Order sellOrder){
        if (!buyOrder.securityName.equals(sellOrder.securityName)) {
            return "Orders are for different security names";
        }
        if (buyOrder.userName.equals(sellOrder.userName)) {
            return "Buyer and seller are the same users";
        }
        if (buyOrder.quantity <= 0) {
            return "Buyer quantity must be positive";
        }
        if (sellOrder.quantity <= 0) {
            return "Seller quantity must be positive";
        }
        if (buyOrder.price <= 0) {
            return "Buyer price must be positive";
        }
        if (sellOrder.price <= 0) {
            return "Sell price must be positive";
        }
        return "";
    }
}
