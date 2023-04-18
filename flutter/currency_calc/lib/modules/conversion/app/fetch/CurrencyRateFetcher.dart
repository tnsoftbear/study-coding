import 'CurrencyRateFetchingInput.dart';

abstract class CurrencyRateFetcher {
  Future<double> fetchExchangeRate(CurrencyRateFetchingInput input);
}