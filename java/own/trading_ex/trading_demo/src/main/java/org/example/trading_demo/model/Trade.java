package org.example.trading_demo.model;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Entity
@Table(name = "trade")
@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
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

    @Override
    public String toString() {
        return "Trade{" +
                "id=" + id +
                ", price=" + price +
                ", quantity=" + quantity +
                ", sellOrderId=" + sellOrderId +
                ", buyOrderId=" + buyOrderId +
                '}';
    }
}
