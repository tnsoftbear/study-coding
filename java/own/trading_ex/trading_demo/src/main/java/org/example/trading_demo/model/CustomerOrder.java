package org.example.trading_demo.model;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Setter
@Getter
@NoArgsConstructor
@AllArgsConstructor
public class CustomerOrder {
    public String securityName;
    public String userName;
    public int price;
    public int quantity;

    @Override
    public String toString() {
        return "securityName: " + securityName + ", userName: " + userName + ", price: " + price + ", quantity: " + quantity;
    }
}
