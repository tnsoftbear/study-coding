package org.example.trading_demo.model;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Setter
@Getter
@NoArgsConstructor
public class ExchangeRequest {
    private Order buyerOrder;
    private Order sellerOrder;
}
