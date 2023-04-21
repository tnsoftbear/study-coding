import 'package:currency_calc/modules/conversion/domain/validate/currency_conversion_validation_result.dart';

class CurrencyConversionValidationTranslator {
  static String get(String key) {
    return _translations[key] ?? key;
  }

  static const Map<int, String> _translations = {
    CurrencyConversionValidationResult.ERR_SOURCE_CURRENCY_INVALID: 'Invalid source currency',
    CurrencyConversionValidationResult.ERR_TARGET_CURRENCY_INVALID: 'Invalid target currency',
    CurrencyConversionValidationResult.ERR_SOURCE_AND_TARGET_CURRENCY_SAME: 'Source and target currencies cannot be the same',
    CurrencyConversionValidationResult.ERR_SOURCE_AMOUNT_NOT_NUMERIC: 'Numeric amount expected',
    CurrencyConversionValidationResult.ERR_SOURCE_AMOUNT_NOT_POSITIVE: 'Source amount must be positive',
  };

  static List<String> translate(List<int> errors) {
    final List<String> errorMessages = [];
    errors.forEach((error) {
      errorMessages.add(_translations[error] ?? '');
    });
    return errorMessages;
  }
}