import 'package:currency_calc/modules/conversion/app/config/CurrencyConversionConfig.dart';
import 'package:currency_calc/modules/conversion/infra/fetch/FawazAhmedCurrencyRateFetcher.dart';
import 'package:currency_calc/modules/conversion/infra/fetch/FixerIoCurrencyRateFetcher.dart';
import 'CurrencyRateFetcherConstants.dart';
import 'CurrencyRateFetcher.dart';

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
