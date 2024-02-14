package org.example.trading_demo.service;

import lombok.AllArgsConstructor;
import org.example.trading_demo.model.CustomerOrder;
import org.example.trading_demo.model.Security;
import org.example.trading_demo.model.StoredOrder;
import org.example.trading_demo.model.User;
import org.example.trading_demo.repository.SecurityRepository;
import org.example.trading_demo.repository.StoredOrderRepository;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@AllArgsConstructor
public class OrderService {
    private StoredOrderRepository orderRepository;
    private UserService userService;
    private SecurityRepository securityRepository;

    public StoredOrder create(int price, int quantity, int type, long securityId, long userId) {
        StoredOrder order = new StoredOrder();
        order.setPrice(price);
        order.setQuantity(quantity);
        order.setType(type);
        order.setSecurityId(securityId);
        order.setUserId(userId);
        order.setFulfilled(false);
        return orderRepository.save(order);
    }

    public StoredOrder create(CustomerOrder customerOrder, Boolean isSellOrder) {
        User user = userService.findByUsername(customerOrder.getUserName());
        if (user == null) {
            throw new IllegalArgumentException("User not found");
        }
        Security security = securityRepository.findByName(customerOrder.getSecurityName());
        if (security == null) {
            throw new IllegalArgumentException("Security not found");
        }

        return this.create(customerOrder.getPrice(), customerOrder.getQuantity(), isSellOrder ? 1 : 0, security.getId(), user.getId());
    }

    public StoredOrder findFirstByTypeAndSecurityId(int type, long securityId) {
        List<StoredOrder> orders = orderRepository.findByTypeAndSecurityIdAndFulfilledIsFalse(type, securityId);
        if (orders.isEmpty()) {
            return null;
        }
        return orders.getFirst();
    }

    public void markFulfilled(StoredOrder order) {
        order.setFulfilled(true);
        orderRepository.save(order);
    }
}
