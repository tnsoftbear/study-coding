package org.example.trading_demo.model.stored_order;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;


@Entity
@Table(name = "stored_order")
@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class StoredOrder {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;
    @Column(name = "price")
    private int price;
    @Column(name = "quantity")
    private int quantity;
    @Column(name = "type")
    private Type type;
    @Column(name = "fulfilled")
    private Boolean fulfilled;
    @Column(name = "security_id")
    private long securityId;
    @Column(name = "user_id")
    private long userId;

    @Override
    public String toString() {
        return "StoredOrder{" +
                "id=" + id +
                ", price=" + price +
                ", quantity=" + quantity +
                ", type=" + type +
                ", fulfilled=" + fulfilled +
                ", securityId=" + securityId +
                ", userId=" + userId +
                '}';
    }
}
