import 'package:currency_calc/feature/conversion/app/config/currency_conversion_config.dart';
import 'package:currency_calc/feature/conversion/domain/fetch/currency_conversion_rate_caching_fetcher.dart';
import 'package:currency_calc/feature/conversion/domain/fetch/load/currency_conversion_rate_fetcher.dart';
import 'package:currency_calc/feature/conversion/infra/fetch/cache/currency_conversion_rate_hive_cacher.dart';
import 'package:currency_calc/feature/conversion/infra/fetch/constant/currency_conversion_fetching_constant.dart';
import 'package:currency_calc/feature/conversion/infra/fetch/load/fawaz_ahmed_currency_rate_fetcher.dart';
import 'package:currency_calc/feature/conversion/infra/fetch/load/fixer_io_currency_rate_fetcher.dart';

class CurrencyConversionRateFetcherFactory {
  static CurrencyConversionRateFetcher create(CurrencyConversionConfig config) {
    CurrencyConversionRateFetcher? fetcher;
    if (config.currencyConversionRateFetcherType ==
        CurrencyConversionFetchingConstant.FT_FIXER_IO) {
      fetcher = FixerIoCurrencyRateFetcher(
          url: config.fixerIoApiBaseUrl, apiKey: config.fixerIoApiKey);
    } else if (config.currencyConversionRateFetcherType ==
        CurrencyConversionFetchingConstant.FT_FAWAZ_AHMED) {
      fetcher =
          FawazAhmedCurrencyConversionFetcher(url: config.fawazAhmedApiBaseUrl);
    }

    if (fetcher == null) {
      throw Exception('Invalid currency rate fetcher type');
    }

    final cacher = CurrencyConversionRateHiveCacher(
        config.currencyConversionRateCacheExpiryInSeconds);
    return CurrencyConversionRateCashingFetcher(fetcher, cacher);
  }
}
