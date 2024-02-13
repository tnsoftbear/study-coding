package org.example.trading_demo.test;

import org.junit.platform.suite.api.ConfigurationParameter;
import org.junit.platform.suite.api.IncludeEngines;
import org.junit.platform.suite.api.SelectClasspathResource;
import org.junit.platform.suite.api.Suite;
import static io.cucumber.junit.platform.engine.Constants.PLUGIN_PROPERTY_NAME;

@Suite
@IncludeEngines("cucumber")
@SelectClasspathResource("features")
@ConfigurationParameter(key = PLUGIN_PROPERTY_NAME, value = "pretty")
public class RunCucumberTests {
}


//
//import io.cucumber.junit.Cucumber;
//import io.cucumber.junit.CucumberOptions;
//import org.junit.jupiter.api.Test;
//import org.junit.runner.RunWith;

//@RunWith(Cucumber.class)
//@CucumberOptions(
//        features = {"src/test/resources/features"}
//        glue = {"org.example.trading_demo.test"}
//        // glue = {"steps"}
//)
//public class RunCucumberTests {
//}