import 'package:currency_calc/modules/conversion/domain/validate/currency_conversion_validation_result.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class CurrencyConversionValidationTranslator {
  static List<String> translateErrorMessages(
      {required BuildContext context, required List<int> errors}) {
    final appLocalizations = AppLocalizations.of(context);
    final Map<int, String> _translations = {
      CurrencyConversionValidationResult.ERR_SOURCE_CURRENCY_INVALID:
          appLocalizations.conversionValidationErrSourceCurrencyInvalid,
      CurrencyConversionValidationResult.ERR_TARGET_CURRENCY_INVALID:
          appLocalizations.conversionValidationErrTargetCurrencyInvalid,
      CurrencyConversionValidationResult.ERR_SOURCE_AND_TARGET_CURRENCY_SAME:
          appLocalizations.conversionValidationErrSourceAndTargetCurrencySame,
      CurrencyConversionValidationResult.ERR_SOURCE_AMOUNT_NOT_NUMERIC:
          appLocalizations.conversionValidationErrSourceAmountNotNumeric,
      CurrencyConversionValidationResult.ERR_SOURCE_AMOUNT_NOT_POSITIVE:
          appLocalizations.conversionValidationErrSourceAmountNotPositive
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
