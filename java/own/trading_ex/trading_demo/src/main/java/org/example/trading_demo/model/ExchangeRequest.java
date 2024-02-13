package org.example.trading_demo.model;

public class ExchangeRequest {
    private Order buyerOrder;
    private Order sellerOrder;

    public Order getBuyerOrder() {
        return buyerOrder;
    }

    public void setBuyerOrder(Order buyerOrder) {
        this.buyerOrder = buyerOrder;
    }

    public Order getSellerOrder() {
        return sellerOrder;
    }

    public void setSellerOrder(Order sellerOrder) {
        this.sellerOrder = sellerOrder;
    }
}
