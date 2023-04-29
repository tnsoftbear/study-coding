import 'package:currency_calc/feature/conversion/domain/fetch/cache/currency_conversion_rate_cacher.dart';
import 'package:currency_calc/feature/conversion/domain/model/fetch/currency_conversion_rate_fetch_record.dart';

class CurrencyConversionRateMemoryCacher extends CurrencyConversionRateCacher {
  int ttl = 0;
  static Map<String, CurrencyConversionRateFetchRecord> _cache = {};

  CurrencyConversionRateMemoryCacher(this.ttl);

  Future<double?> get(String from, String to) {
    final key = _makeKey(from, to);
    if (_cache.containsKey(key)) {
      final record = _cache[key];
      if (record != null) {
        DateTime expiredAt = record.createdAt.add(Duration(seconds: ttl));
        print('Date now: ' + DateTime.now().toString() +
            '; Expired at: ' +
            expiredAt.toString());
        if (DateTime.now().isBefore(expiredAt)) {
          print("Cache hit: " + record.createdAt.toString());
          return Future.value(record.exchangeRate);
        }
      }

      print("Cache expired");
      _cache.remove(key);
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
    _cache[key] = record;
  }

  String _makeKey(String from, String to) {
    return from + to;
  }
}
