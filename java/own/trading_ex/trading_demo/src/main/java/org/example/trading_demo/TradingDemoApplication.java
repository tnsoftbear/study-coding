package org.example.trading_demo;

import jakarta.annotation.PostConstruct;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import java.util.Arrays;

@SpringBootApplication
public class TradingDemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(TradingDemoApplication.class, args);
	}

//	@PostConstruct
//	private void initDb() {
//		String sqlStatements[] = {
//				"drop table employees if exists",
//				"create table employees(id serial,first_name varchar(255),last_name varchar(255))",
//				"insert into employees(first_name, last_name) values('Eugen','Paraschiv')",
//				"insert into employees(first_name, last_name) values('Scott','Tiger')"
//		};
//
//		Arrays.asList(sqlStatements).forEach(sql -> {
//			jdbcTemplate.execute(sql);
//		});
//
//		// Query test data and print results
//	}

}
