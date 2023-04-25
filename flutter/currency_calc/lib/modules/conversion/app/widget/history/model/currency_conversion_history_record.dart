import 'package:hive/hive.dart';

part 'currency_conversion_history_record.g.dart';

@HiveType(typeId: 0)
class CurrencyConversionHistoryRecord extends HiveObject {
  @HiveField(0)
  String sourceCurrency = '';

  @HiveField(1)
  double sourceAmount = 0.0;

  @HiveField(2)
  String targetCurrency = '';

  @HiveField(3)
  double targetAmount = 0.0;

  @HiveField(4)
  double rate = 0.0;

  @HiveField(5)
  DateTime? date;

  toList() {
    return [
      sourceCurrency,
      sourceAmount,
      targetCurrency,
      targetAmount,
      rate,
      date
    ];
  }
}
