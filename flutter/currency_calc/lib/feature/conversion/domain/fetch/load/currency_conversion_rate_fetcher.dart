abstract class CurrencyConversionRateFetcher {
  Future<double> fetchExchangeRate(String from, String to);
}