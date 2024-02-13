package org.example.trading_demo.model;

public class Order {
    public String securityName;
    public String userName;
    public double price;
    public int quantity;
    // Геттеры и сеттеры

    public Order(){}

    public Order(String securityName, String userName, double price, int quantity) {
        this.securityName = securityName;
        this.userName = userName;
        this.price = price;
        this.quantity = quantity;
    }

    @Override
    public String toString() {
        return "securityName: " + securityName + ", userName: " + userName + ", price: " + price + ", quantity: " + quantity;
    }
}
