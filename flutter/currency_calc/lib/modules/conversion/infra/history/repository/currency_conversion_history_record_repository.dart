import 'package:currency_calc/modules/conversion/domain/history/model/currency_conversion_history_record.dart';
import 'package:hive_flutter/hive_flutter.dart';

class CurrencyConversionHistoryRecordRepository {
  static const BOX_NAME = 'currencyConversionHistory';

  Box<CurrencyConversionHistoryRecord>? box;

  Future<void> init() async {
    box = await Hive.openBox<CurrencyConversionHistoryRecord>(BOX_NAME);
  }

  List<CurrencyConversionHistoryRecord> loadAll() {
    return box!.values.toList();
  }

  int countAll() {
    return box!.length;
  }

  Future<void> save(CurrencyConversionHistoryRecord record) async {
    await box!.add(record);
    await box!.close();
  }

  Future<void> deleteByIndex(int index) async {
    await box!.deleteAt(index);
    await box!.close();
  }
}