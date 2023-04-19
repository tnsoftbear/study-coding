# Currency Calculator

My 1st experiment with flutter and dart in April of 2023.

Module logic is separated to application, infrastructure and domain layers.

Infrastructure layer is responsible for the currency exchange rate loading by API. 
Rate fetchers implement common interface and are provided by factory.
Application layer operates by this interface and provides the currency exchange rate to the UI.
Domain layer calculates the currency conversion.

Few tests are added.
