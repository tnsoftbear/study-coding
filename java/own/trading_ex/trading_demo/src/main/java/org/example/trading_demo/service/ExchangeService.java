package org.example.trading_demo.service;

import lombok.AllArgsConstructor;
import org.example.trading_demo.model.CustomerOrder;
import org.example.trading_demo.model.Security;
import org.example.trading_demo.model.User;
import org.example.trading_demo.repository.SecurityRepository;
import org.example.trading_demo.service.user.UserService;
import org.springframework.stereotype.Service;

@Deprecated
@Service
@AllArgsConstructor
final public class ExchangeService {
    private UserService userService;
    private SecurityRepository securityRepository;

    public CustomerOrder exchange(CustomerOrder buyOrder, CustomerOrder sellOrder){
        CustomerOrder result = new CustomerOrder();
        result.quantity = Math.min(buyOrder.quantity, sellOrder.quantity);
        result.price = Math.min(buyOrder.price, sellOrder.price);
        result.securityName = buyOrder.securityName;
        return result;
    }

    public String validate(CustomerOrder buyOrder, CustomerOrder sellOrder){
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
        User buyer = this.userService.findByUsername(buyOrder.userName);
        if (buyer == null) {
            return "Buyer not found by username: " + buyOrder.userName;
        }
        User seller = this.userService.findByUsername(sellOrder.userName);
        if (seller == null) {
            return "Seller not found by username: " + sellOrder.userName;
        }
        Security security = this.securityRepository.findByName(buyOrder.securityName);
        if (security == null) {
            return "Security not found by name: " + buyOrder.securityName;
        }

        return "";
    }
}
