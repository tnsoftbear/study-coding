package org.example.trading_demo.service;

import lombok.AllArgsConstructor;
import org.example.trading_demo.model.stored_order.StoredOrder;
import org.example.trading_demo.model.Trade;
import org.example.trading_demo.repository.TradeRepository;
import org.springframework.stereotype.Service;
import org.example.trading_demo.model.stored_order.Type;
import java.util.logging.Logger;

@Service
@AllArgsConstructor
public class TradeService {
    TradeRepository tradeRepository;
    OrderService orderService;
    private static final Logger log = Logger.getLogger(TradeService.class.getName());

    public Trade tradeWithSeller(StoredOrder sellerOrder) {
        StoredOrder buyerOrder = orderService.findFirstByTypeAndSecurityId(Type.BUYER, sellerOrder.getSecurityId());
        if (buyerOrder == null) {
            return null;
        }

        return createTrade(buyerOrder, sellerOrder);
    }

    public Trade tradeWithBuyer(StoredOrder buyerOrder) {
        StoredOrder sellerOrder = orderService.findFirstByTypeAndSecurityId(Type.SELLER, buyerOrder.getSecurityId());
        if (sellerOrder == null) {
            return null;
        }

        return createTrade(buyerOrder, sellerOrder);
    }

    private Trade createTrade(StoredOrder buyerOrder, StoredOrder sellerOrder) {
        Trade trade = new Trade();
        trade.setQuantity(Math.min(buyerOrder.getQuantity(), sellerOrder.getQuantity()));
        trade.setPrice(Math.min(buyerOrder.getPrice(), sellerOrder.getPrice()));
        trade.setBuyOrderId(buyerOrder.getId());
        trade.setSellOrderId(sellerOrder.getId());
        tradeRepository.save(trade);
        orderService.markFulfilled(buyerOrder);
        orderService.markFulfilled(sellerOrder);
        log.info("Trade registered - " + trade);
        return trade;
    }
}
