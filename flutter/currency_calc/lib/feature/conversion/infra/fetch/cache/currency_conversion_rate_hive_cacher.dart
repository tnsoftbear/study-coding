import 'package:currency_calc/feature/conversion/domain/fetch/cache/currency_conversion_rate_cacher.dart';
import 'package:currency_calc/feature/conversion/domain/model/fetch/currency_conversion_rate_fetch_record.dart';
import 'package:currency_calc/feature/conversion/infra/repository/fetch/currency_conversion_rate_fetch_record_repository.dart';

class CurrencyConversionRateHiveCacher extends CurrencyConversionRateCacher {
  int ttl = 0;

  CurrencyConversionRateHiveCacher(this.ttl);

  Future<double?> get(String from, String to) async {
    final key = _makeKey(from, to);
    final repo = CurrencyConversionRateFetchRecordRepository();
    await repo.init();
    CurrencyConversionRateFetchRecord? record = await repo.loadByKey(key);
    if (record != null) {
      DateTime expiredAt = record.createdAt.add(Duration(seconds: ttl));
      print('Date now: ' +
          DateTime.now().toString() +
          '; Expired at: ' +
          expiredAt.toString());
      if (DateTime.now().isBefore(expiredAt)) {
        print("Cache hit: " + record.createdAt.toString());
        return Future.value(record.exchangeRate);
      }

      print("Cache expired");
      await repo.deleteByKey(key);
      return Future.value(null);
    }

    print("Cache miss");
    return Future.value(null);
  }

  Future<void> set(String from, String to, double rate) async {
    final key = _makeKey(from, to);
    print("Set: " + DateTime.now().toString());
    final record = CurrencyConversionRateFetchRecord(
      sourceCurrency: from,
      targetCurrency: to,
      exchangeRate: rate,
      createdAt: DateTime.now(),
    );
    final repo = CurrencyConversionRateFetchRecordRepository();
    await repo.init();
    await repo.saveByKey(key, record);
  }

  String _makeKey(String from, String to) {
    return from + to;
  }
}
