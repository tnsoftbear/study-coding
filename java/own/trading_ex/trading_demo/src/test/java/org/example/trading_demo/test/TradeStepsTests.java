package org.example.trading_demo.test;

import io.cucumber.java.en.And;
import io.cucumber.java.en.Given;
import io.cucumber.java.en.Then;
import io.cucumber.java.en.When;
import io.cucumber.spring.CucumberContextConfiguration;
import org.example.trading_demo.model.*;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.*;
import org.springframework.web.client.RestTemplate;

@CucumberContextConfiguration
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.DEFINED_PORT)
public class TradeStepsTests {

    private final RestTemplate restTemplate = new RestTemplate();
    private Trade trade;

    @Given("^one security \"([^\"]*)\" and two users \"([^\"]*)\" and \"([^\"]*)\" exist$")
    public void one_security_and_two_users_exist(String securityName, String username1, String username2) {
        User user1 = User.builder().username(username1).password("").build();
        this.postSaveUser(user1);

        User user2 = User.builder().username(username2).password("").build();
        this.postSaveUser(user2);

        Security security = new Security();
        security.setName(securityName);
        this.postSaveSecurity(security);
    }

    private void postSaveUser(User user) {
        String url = "http://localhost:8080/api/v1/users/save";
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        HttpEntity<User> request = new HttpEntity<>(user, headers);
        ResponseEntity<User> response = restTemplate.postForEntity(url, request, User.class);
        if (response.getStatusCode() != HttpStatus.OK) {
            throw new AssertionError("Unexpected status code: " + response.getStatusCode());
        }
    }

    private void postSaveSecurity(Security security) {
        String url = "http://localhost:8080/api/v1/security/save";
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        HttpEntity<Security> request = new HttpEntity<>(security, headers);
        ResponseEntity<Security> response = restTemplate.postForEntity(url, request, Security.class);
        if (response.getStatusCode() != HttpStatus.OK) {
            throw new AssertionError("Unexpected status code: " + response.getStatusCode());
        }
    }

    @When("^user \"([^\"]*)\" puts a buy order for security \"([^\"]*)\" with a price of (\\d+) and quantity of (\\d+)$")
    public void user_puts_a_buy_order(String user, String security, int price, int quantity) {
        CustomerOrder buyerOrder = new CustomerOrder(security, user, price, quantity);
        String url = "http://localhost:8080/api/v1/order/buy";
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        HttpEntity<CustomerOrder> request = new HttpEntity<>(buyerOrder, headers);
        ResponseEntity<CustomerOrder> response = restTemplate.postForEntity(url, request, CustomerOrder.class);
        if (response.getStatusCode() != HttpStatus.OK) {
            throw new AssertionError("Unexpected status code: " + response.getStatusCode());
        }
    }

    @And("^user \"([^\"]*)\" puts a sell order for security \"([^\"]*)\" with a price of (\\d+) and a quantity of (\\d+)$")
    public void user_puts_a_sell_order(String user, String security, int price, int quantity) {
        CustomerOrder sellerOrder = new CustomerOrder(security, user, price, quantity);
        String url = "http://localhost:8080/api/v1/order/sell_and_trade";
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        HttpEntity<CustomerOrder> request = new HttpEntity<>(sellerOrder, headers);
        ResponseEntity<Trade> response = restTemplate.postForEntity(url, request, Trade.class);
        if (response.getStatusCode() != HttpStatus.OK) {
            throw new AssertionError("Unexpected status code: " + response.getStatusCode());
        }
        this.trade = response.getBody();
    }

    @Then("^a trade occurs with the price of (\\d+) and quantity of (\\d+)$")
    public void trade_occurs(int expectedPrice, int expectedQuantity) {
        if (this.trade.getPrice() != expectedPrice) {
            throw new AssertionError("Trade price is not expected");
        }
        if (this.trade.getQuantity() != expectedQuantity) {
            throw new AssertionError("Trade quantity is not expected");
        }
    }

}