import 'package:currency_calc/modules/conversion/domain/calculator/currency_conversion_result.dart';

class CurrencyConverter {
  static CurrencyConversionResult convert(double amount, double rate) {
    double resultAmount = amount * rate;
    return CurrencyConversionResult(targetAmount: resultAmount, rate: rate);
  }
}
