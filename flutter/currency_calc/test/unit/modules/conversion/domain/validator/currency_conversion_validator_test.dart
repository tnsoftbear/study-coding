import 'package:currency_calc/modules/conversion/domain/validate/currency_conversion_validation_result.dart';
import 'package:currency_calc/modules/conversion/domain/validate/currency_conversion_validator.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:currency_calc/modules/conversion/app/constant/currency_constant.dart';

void main() {
  group('CurrencyConversionValidator', () {
    test('validate should return no errors for valid inputs', () {
      final result = CurrencyConversionValidator.validate(
          sourceCurrency: CurrencyConstant.CURRENCIES[0],
          targetCurrency: CurrencyConstant.CURRENCIES[1],
          amount: '100.00');
      expect(result.errors.length, 0);
    });

    test(
        'validate should return an error when source currency is invalid and others are valid',
            () {
          final result = CurrencyConversionValidator.validate(
              sourceCurrency: 'XYZ',
              targetCurrency: CurrencyConstant.CURRENCIES[1],
              amount: '100.00');
          expect(result.errors.length, 1);
          expect(result.errors[0],
              CurrencyConversionValidationResult.ERR_SOURCE_CURRENCY_INVALID);
        });

    test(
        'validate should return an error when target currency is invalid and others are valid',
            () {
          final result = CurrencyConversionValidator.validate(
              sourceCurrency: CurrencyConstant.CURRENCIES[0],
              targetCurrency: 'XYZ',
              amount: '100.00');
          expect(result.errors.length, 1);
          expect(result.errors[0],
              CurrencyConversionValidationResult.ERR_TARGET_CURRENCY_INVALID);
        });

    test(
        'validate should return an error when source and target currencies are same and others are valid',
            () {
          final result = CurrencyConversionValidator.validate(
              sourceCurrency: CurrencyConstant.CURRENCIES[0],
              targetCurrency: CurrencyConstant.CURRENCIES[0],
              amount: '100.00');
          expect(result.errors.length, 1);
          expect(result.errors[0],
              CurrencyConversionValidationResult
                  .ERR_SOURCE_AND_TARGET_CURRENCY_SAME);
        });

    test('validate should return an error when amount is not numeric', () {
      final result = CurrencyConversionValidator.validate(
          sourceCurrency: CurrencyConstant.CURRENCIES[0],
          targetCurrency: CurrencyConstant.CURRENCIES[1],
          amount: 'not a number');
      expect(result.errors.length, 1);
      expect(result.errors[0],
          CurrencyConversionValidationResult.ERR_SOURCE_AMOUNT_NOT_NUMERIC);
    });

    test('validate should return an error when amount is not positive', () {
      final result = CurrencyConversionValidator.validate(
          sourceCurrency: CurrencyConstant.CURRENCIES[0],
          targetCurrency: CurrencyConstant.CURRENCIES[1],
          amount: '0.00');
      expect(result.errors.length, 1);
      expect(result.errors[0],
          CurrencyConversionValidationResult.ERR_SOURCE_AMOUNT_NOT_POSITIVE);
    });

    test('validate should return errors because all inputs are wrong', () {
      final result = CurrencyConversionValidator.validate(
          sourceCurrency: 'XXX',
          targetCurrency: 'XXX',
          amount: 'XXX');
      expect(result.errors.length, 4);
      expect(result.errors, equals([
        CurrencyConversionValidationResult.ERR_SOURCE_CURRENCY_INVALID,
        CurrencyConversionValidationResult.ERR_TARGET_CURRENCY_INVALID,
        CurrencyConversionValidationResult.ERR_SOURCE_AND_TARGET_CURRENCY_SAME,
        CurrencyConversionValidationResult.ERR_SOURCE_AMOUNT_NOT_NUMERIC
      ]));
    });
  });
}
