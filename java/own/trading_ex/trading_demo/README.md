# REST API by Spring Boot

* Model: User, Security, Order, Trade
* H2 memory DB. Console at <http://localhost:8080/h2-ui>
* Cucumber BDD test

## Run RestAPI server

```sh
mvnw clean spring-boot:run
# or
mvn clean package
java -jar ./target/trading_demo-0.0.1-SNAPSHOT.jar
# or via docker
docker compose build
docker compose up
# read user list for api checking
curl http://127.0.0.1:8080/api/v1/users 
```

## Run BDD test

```sh
# via local java
mvn test
# or
mvnw test
```

### Test with reports

```sh
mvn verify
# or
mvnw verify
# via docker
docker build -t trading_demo_bdd_test -f ./Dockerfile.runtest .
docker run trading_demo_bdd_test
# enter docker container
docker run -it trading_demo_bdd_test bash
```

Observe HTML reports in

* `target/site/jacoco/index.html` Jacoco coverage report

![20240215-jacoco-report.png](doc/i/20240215-jacoco-report.png)

* `target/generated-report/index.html` Cluecumber report

![20240215-add-user-and-search-report.png](doc/i/20240215-add-user-and-search-report.png)
![20240215-all-steps-report.png](doc/i/20240215-all-steps-report.png)

* [trivago/cluecumber](https://github.com/trivago/cluecumber/)

* `target/cucumber-report/cucumber.html`

![20240216-cucumber-report.png](doc/i/20240216-cucumber-report.png)

## Manual requests

* post buy order and trade: POST: http://localhost:8080/api/v1/order/buy_and_trade

```json
{
  "price": 100,
  "quantity": 80,
  "securityName": "Apple",
  "userName": "user4"
}
```

* post sell order and trade: POST: http://localhost:8080/api/v1/order/sell_and_trade

```json
{
  "price": 150,
  "quantity": 40,
  "securityName": "Apple",
  "userName": "user3"
}
```
* list orders: GET: http://localhost:8080/api/v1/order

* list users: GET: http://localhost:8080/api/v1/users

`curl -X GET http://localhost:8080/api/v1/users`

* add user: POST: http://localhost:8080/api/v1/users/save

```json
{
    "username": "user5",
    "password": "pw5"
}
```

* list securities: GET: http://localhost:8080/api/v1/security
* find by name: GET: http://localhost:8080/api/v1/security/Apple
* add security: POST: http://localhost:8080/api/v1/security/save

```json
{
    "name": "sec-name"
}
```

* delete security: DELETE: http://localhost:8080/api/v1/security/delete/sec-name