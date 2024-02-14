package org.example.trading_demo.model;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Setter
@Getter
@NoArgsConstructor
public class ExchangeRequest {
    private CustomerOrder buyerOrder;
    private CustomerOrder sellerOrder;
}
