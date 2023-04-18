import 'package:currency_calc/modules/conversion/domain/calculator/CurrencyConversionResult.dart';

class CurrencyConverter {
  static CurrencyConversionResult convert(double amount, double rate) {
    double resultAmount = amount * rate;
    return CurrencyConversionResult(targetAmount: resultAmount, rate: rate);
  }
}
