package org.example.trading_demo.repository;

import org.example.trading_demo.model.StoredOrder;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface StoredOrderRepository extends JpaRepository<StoredOrder, Long> {
    List<StoredOrder> findByTypeAndSecurityIdAndFulfilledIsFalse(int type, long securityId);
}
