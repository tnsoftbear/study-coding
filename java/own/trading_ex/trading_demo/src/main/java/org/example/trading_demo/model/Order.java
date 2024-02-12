package org.example.trading_demo.model;

public class Order {
    public String securityName;
    public String userName;
    public double price;
    public int quantity;
    // Геттеры и сеттеры

    @Override
    public String toString() {
        return "securityName: " + securityName + ", userName: " + userName + ", price: " + price + ", quantity: " + quantity;
    }
}
