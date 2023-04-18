import 'package:CurrencyCalc/modules/conversion/app/config/CurrencyConversionConfig.dart';
import 'package:CurrencyCalc/modules/conversion/infra/fetch/FawazAhmedCurrencyRateFetcher.dart';
import 'package:CurrencyCalc/modules/conversion/infra/fetch/FixerIoCurrencyRateFetcher.dart';
import 'CurrencyRateFetcherConstants.dart';
import 'CurrencyRateFetcher.dart';

class CurrencyRateFetcherFactory {
  static CurrencyRateFetcher create(
      {required CurrencyConversionConfig config,
      String type = CurrencyRateFetcherConstants.RF_FAWAZ_AHMED}) {
    if (type == CurrencyRateFetcherConstants.RF_FIXER_IO) {
      return FixerIoCurrencyRateFetcher(
          url: config.FixerIoApiBaseUrl, apiKey: config.FixerIoApiKey);
    }

    return FawazAhmedCurrencyConversionFetcher(
        url: config.FawazAhmedApiBaseUrl);
  }
}
