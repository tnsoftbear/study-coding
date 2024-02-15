package org.example.trading_demo.model;

import lombok.*;

@Setter
@Getter
@Data
public class CustomerOrder {
    public String securityName;
    public String userName;
    public int price;
    public int quantity;
}
