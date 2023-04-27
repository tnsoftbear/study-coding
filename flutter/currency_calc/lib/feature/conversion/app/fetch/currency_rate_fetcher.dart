import 'currency_rate_fetching_input.dart';

abstract class CurrencyRateFetcher {
  Future<double> fetchExchangeRate(CurrencyRateFetchingInput input);
}