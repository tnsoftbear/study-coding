package org.example.trading_demo.test.common;

import io.cucumber.java.After;
import io.cucumber.java.Before;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class StepHooks {

    @Before
    public void setUp() {
        log.info("Up: --------------------------------------------------------------");
    }

    @After
    public void tearDown() {
        log.info("Down: --------------------------------------------------------------");
    }
}
