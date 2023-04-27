import 'package:currency_calc/feature/conversion/domain/validate/currency_conversion_validation_result.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class CurrencyConversionValidationTranslator {
  static List<String> translateErrorMessages(
      {required BuildContext context, required List<int> errors}) {
    final tr = AppLocalizations.of(context);
    final Map<int, String> _translations = {
      CurrencyConversionValidationResult.ERR_SOURCE_CURRENCY_INVALID:
          tr.conversionValidationErrSourceCurrencyInvalid,
      CurrencyConversionValidationResult.ERR_TARGET_CURRENCY_INVALID:
          tr.conversionValidationErrTargetCurrencyInvalid,
      CurrencyConversionValidationResult.ERR_SOURCE_AND_TARGET_CURRENCY_SAME:
          tr.conversionValidationErrSourceAndTargetCurrencySame,
      CurrencyConversionValidationResult.ERR_SOURCE_AMOUNT_NOT_NUMERIC:
          tr.conversionValidationErrSourceAmountNotNumeric,
      CurrencyConversionValidationResult.ERR_SOURCE_AMOUNT_NOT_POSITIVE:
          tr.conversionValidationErrSourceAmountNotPositive
    };

    final List<String> errorMessages = [];
    errors.forEach((error) {
      errorMessages.add(_translations[error] ?? '');
    });
    return errorMessages;
  }

  static String translateConcatenatedErrorMessage(
      {required BuildContext context,
      required CurrencyConversionValidationResult validationResult,
      String separator = "\n"}) {
    final List<String> errorMessages = translateErrorMessages(
        context: context, errors: validationResult.errors);
    return errorMessages.join(separator);
  }
}
