import 'package:currency_calc/feature/conversion/domain/fetch/load/currency_conversion_rate_fetcher.dart';
import 'package:currency_calc/feature/conversion/domain/fetch/cache/currency_conversion_rate_cacher.dart';

class CurrencyConversionRateCashingFetcher implements CurrencyConversionRateFetcher {
  final CurrencyConversionRateFetcher _currencyRateFetcher;
  final CurrencyConversionRateCacher _currencyRateCacher;

  CurrencyConversionRateCashingFetcher(
      this._currencyRateFetcher, this._currencyRateCacher);

  @override
  Future<double> fetchExchangeRate(String from, String to) async {
    final cachedCurrencyRate = await _currencyRateCacher.get(from, to);
    if (cachedCurrencyRate != null) {
      return cachedCurrencyRate;
    }

    final currencyRate = await _currencyRateFetcher.fetchExchangeRate(from, to);
    await _currencyRateCacher.set(from, to, currencyRate);
    return currencyRate;
  }
}