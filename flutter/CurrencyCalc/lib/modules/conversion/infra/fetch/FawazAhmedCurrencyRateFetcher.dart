import 'package:CurrencyCalc/modules/conversion/app/fetch/CurrencyRateFetcher.dart';
import 'package:CurrencyCalc/modules/conversion/app/fetch/CurrencyRateFetchingInput.dart';
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
    print(url);
    final response = await http.get(Uri.parse(url));
    if (response.statusCode == 200) {
      final data = json.decode(response.body);
      print(data);
      return data[toCurrencyLower];
    } else {
      throw Exception('Failed to load exchange rate');
    }
  }
}
