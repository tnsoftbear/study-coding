import 'package:currency_calc/modules/conversion/app/fetch/currency_rate_fetcher.dart';
import 'package:currency_calc/modules/conversion/app/fetch/currency_rate_fetching_input.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class FawazAhmedCurrencyConversionFetcher extends CurrencyRateFetcher {
  final String url;

  FawazAhmedCurrencyConversionFetcher({required this.url});

  Future<double> fetchExchangeRate(CurrencyRateFetchingInput input) async {
    final toCurrencyLower = input.to.toLowerCase();
    final fromCurrencyLower = input.from.toLowerCase();
    final url =
        this.url + '/currencies/$fromCurrencyLower/$toCurrencyLower.json';
    final response = await http.get(Uri.parse(url));
    if (response.statusCode == 200) {
      final data = json.decode(response.body);
      return data[toCurrencyLower];
    } else {
      throw Exception('Failed to load exchange rate');
    }
  }
}
