import 'package:currency_calc/feature/conversion/app/config/currency_conversion_config.dart';
import 'package:currency_calc/feature/conversion/infra/fetch/fawaz_ahmed_currency_rate_fetcher.dart';
import 'package:currency_calc/feature/conversion/infra/fetch/fixer_io_currency_rate_fetcher.dart';
import 'currency_rate_fetcher_constants.dart';
import 'currency_rate_fetcher.dart';

class CurrencyRateFetcherFactory {
  static CurrencyRateFetcher create(
      {required CurrencyConversionConfig config,
      String type = CurrencyRateFetcherConstants.RF_FAWAZ_AHMED}) {
    if (type == CurrencyRateFetcherConstants.RF_FIXER_IO) {
      return FixerIoCurrencyRateFetcher(
          url: config.fixerIoApiBaseUrl, apiKey: config.fixerIoApiKey);
    }

    return FawazAhmedCurrencyConversionFetcher(
        url: config.fawazAhmedApiBaseUrl);
  }
}
