package org.example.trading_demo.controller;

import org.example.trading_demo.model.ExchangeRequest;
import org.example.trading_demo.model.CustomerOrder;
import org.example.trading_demo.service.ExchangeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ExchangeController {

    final private ExchangeService exchangeService;

    @Autowired
    public ExchangeController(ExchangeService exchangeService){
        this.exchangeService = exchangeService;
    }

    @PostMapping("/api/v1/exchange")
    public ResponseEntity<?> exchange(@RequestBody ExchangeRequest request) {
        String errorMessage = this.exchangeService.validate(request.getBuyerOrder(), request.getSellerOrder());
        if (!errorMessage.isEmpty()) {
            return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(errorMessage);
        }
        CustomerOrder resultOrder = this.exchangeService.exchange(request.getBuyerOrder(), request.getSellerOrder());
        return ResponseEntity.ok(resultOrder);
    }
}
