package org.example.trading_demo.model;

import jakarta.persistence.*;
import lombok.*;

@Entity
@Table(name = "trade")
@Data
public class Trade {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;
    @Column(name = "price")
    private int price;
    @Column(name = "quantity")
    private int quantity;
    @Column(name = "sell_order_id")
    private long sellOrderId;
    @Column(name = "buy_order_id")
    private long buyOrderId;
}
