import 'package:currency_calc/modules/conversion/app/constant/currency_constant.dart';

import 'currency_conversion_validation_result.dart';

class CurrencyConversionValidator {
  static CurrencyConversionValidationResult validate(
      {required String sourceCurrency,
      required String targetCurrency,
      required String amount}) {
    final result = CurrencyConversionValidationResult();

    if (!CurrencyConstant.CURRENCIES.contains(sourceCurrency)) {
      result.addError(
          CurrencyConversionValidationResult.ERR_SOURCE_CURRENCY_INVALID);
    }

    if (!CurrencyConstant.CURRENCIES.contains(targetCurrency)) {
      result.addError(
          CurrencyConversionValidationResult.ERR_TARGET_CURRENCY_INVALID);
    }

    if (sourceCurrency == targetCurrency) {
      result.addError(CurrencyConversionValidationResult
          .ERR_SOURCE_AND_TARGET_CURRENCY_SAME);
    }

    double? sourceAmount = double.tryParse(amount);
    if (sourceAmount == null) {
      result.addError(
          CurrencyConversionValidationResult.ERR_SOURCE_AMOUNT_NOT_NUMERIC);
    } else if (sourceAmount <= 0) {
      result.addError(
          CurrencyConversionValidationResult.ERR_SOURCE_AMOUNT_NOT_POSITIVE);
    }

    return result;
  }
}
