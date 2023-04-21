import 'package:currency_calc/modules/conversion/app/fetch/currency_rate_fetcher.dart';
import 'package:currency_calc/modules/conversion/app/fetch/currency_rate_fetching_input.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class FixerIoCurrencyRateFetcher extends CurrencyRateFetcher {
  final String url;
  final String apiKey;

  FixerIoCurrencyRateFetcher({required this.url, required this.apiKey});

  Future<double> fetchExchangeRate(CurrencyRateFetchingInput input) async {
    final fromCurrency = input.from;
    final toCurrency = input.to;
    final url = this.url + '?from=$fromCurrency&to=$toCurrency&amount=1';
    final response =
        await http.get(Uri.parse(url), headers: {'apikey': this.apiKey});
    if (response.statusCode == 200) {
      final data = json.decode(response.body);
      // print(data);
      if (data['success'] == false) {
        throw Exception(data['error']['info']);
      }
      return data['info']['rate'];
    } else {
      throw Exception('Failed to load exchange rate');
    }
  }
}
