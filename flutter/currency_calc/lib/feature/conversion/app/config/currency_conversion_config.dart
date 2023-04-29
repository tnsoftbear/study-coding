import 'package:currency_calc/feature/conversion/infra/fetch/constant/currency_conversion_fetching_constant.dart';

class CurrencyConversionConfig {
  final int currencyConversionRateCacheExpiryInSeconds = 60;
  final int currencyConversionRateCacheType =
      CurrencyConversionFetchingConstant.CT_HIVE;
  final int currencyConversionRateFetcherType =
      CurrencyConversionFetchingConstant.FT_FAWAZ_AHMED;

  final String fixerIoApiBaseUrl = 'https://api.apilayer.com/fixer/convert';
  final String fixerIoApiKey = '1yUWc2Kb2Bzr13w7hryFnkKBCxGV38Ia';

  final String fawazAhmedApiBaseUrl =
      'https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest';
}
