import 'package:currency_calc/feature/conversion/domain/model/fetch/currency_conversion_rate_fetch_record.dart';
import 'package:hive_flutter/hive_flutter.dart';

class CurrencyConversionRateFetchRecordRepository {
  static const BOX_NAME = 'CurrencyConversionRateFetchRecord';

  Box<CurrencyConversionRateFetchRecord>? box;

  Future<CurrencyConversionRateFetchRecordRepository> init() async {
    box = await Hive.openBox<CurrencyConversionRateFetchRecord>(BOX_NAME);
    return this;
  }

  Future<CurrencyConversionRateFetchRecord?> loadByKey(String key) async {
    return await box!.get(key);
  }

  Future<void> saveByKey(String key, CurrencyConversionRateFetchRecord record) async {
    await box!.put(key, record);
    await box!.close();
  }

  Future<void> deleteByKey(String key) async {
    await box!.delete(key);
    await box!.close();
  }
}