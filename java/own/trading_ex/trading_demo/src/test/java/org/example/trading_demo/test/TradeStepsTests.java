package org.example.trading_demo.test;

import io.cucumber.java.en.And;
import io.cucumber.java.en.Given;
import io.cucumber.java.en.Then;
import io.cucumber.java.en.When;
import io.cucumber.spring.CucumberContextConfiguration;
import org.example.trading_demo.TradingDemoApplication;
import org.example.trading_demo.model.ExchangeRequest;
import org.example.trading_demo.model.Order;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.*;
import org.springframework.web.client.RestTemplate;

@CucumberContextConfiguration
@SpringBootTest(classes = TradingDemoApplication.class, webEnvironment = SpringBootTest.WebEnvironment.DEFINED_PORT)
public class TradeStepsTests {

    private final RestTemplate restTemplate = new RestTemplate();

    private ExchangeRequest exchangeRequest;

    @Given("^one security \"([^\"]*)\" and two users \"([^\"]*)\" and \"([^\"]*)\" exist$")
    public void one_security_and_two_users_exist(String security, String user1, String user2) {
        // Implementation to create security and users
    }

    @When("^user \"([^\"]*)\" puts a buy order for security \"([^\"]*)\" with a price of (\\d+) and quantity of (\\d+)$")
    public void user_puts_a_buy_order(String user, String security, int price, int quantity) {
        Order buyerOrder = new Order(security, user, price, quantity);
        this.exchangeRequest = new ExchangeRequest();
        this.exchangeRequest.setBuyerOrder(buyerOrder);
    }

    @And("^user \"([^\"]*)\" puts a sell order for security \"([^\"]*)\" with a price of (\\d+) and a quantity of (\\d+)$")
    public void user_puts_a_sell_order(String user, String security, int price, int quantity) {
        Order sellerOrder = new Order(security, user, price, quantity);
        this.exchangeRequest.setSellerOrder(sellerOrder);
    }

    // Метод для проверки текущего ответа на торговлю
    @Then("^a trade occurs with the price of (\\d+) and quantity of (\\d+)$")
    public void trade_occurs(int expectedPrice, int expectedQuantity) {
        String url = "http://localhost:8080/exchange";
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        HttpEntity<ExchangeRequest> request = new HttpEntity<>(this.exchangeRequest, headers);
        ResponseEntity<Order> response = restTemplate.postForEntity(url, request, Order.class);
        if (response.getStatusCode() == HttpStatus.OK) {
            if (response.getBody() != null) {
                Order order = response.getBody();
                if (order.price == expectedPrice && order.quantity == expectedQuantity) {
                    // Сделка соответствует ожидаемым параметрам, ничего не делаем
                } else {
                    // throw new AssertionError("Trade parameters do not match expected values");
                    String errorMessage = String.format("Trade parameters do not match expected values. Actual price: %s, expected price: %s, actual quantity: %s, expected quantity: %s",
                            order.price, expectedPrice, order.quantity, expectedQuantity);
                    throw new AssertionError(errorMessage);
                }
            } else {
                throw new AssertionError("No order data found in response body");
            }
        } else {
            throw new AssertionError("Unexpected status code: " + response.getStatusCode());
        }
    }

}