package org.example.trading_demo.controller;

import org.example.trading_demo.model.ExchangeRequest;
import org.example.trading_demo.model.Order;
import org.example.trading_demo.service.ExchangeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ExchangeController {

    private ExchangeService exchangeService;

    @Autowired
    public ExchangeController(ExchangeService exchangeService){
        this.exchangeService = exchangeService;
    }

    @GetMapping("/list")
    public ResponseEntity<String> listOrder() {
        return ResponseEntity.ok("Order list");
    }

    @PostMapping("/exchange")
    public ResponseEntity<Order> exchange(@RequestBody ExchangeRequest request) {
        Order resultOrder = this.exchangeService.exchange(request.getBuyerOrder(), request.getSellerOrder());
        return ResponseEntity.ok(resultOrder);
    }
}
