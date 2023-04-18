import 'package:CurrencyCalc/modules/conversion/infra/CurrencyRateFetchingInput.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

Future<double> fetchExchangeRate(CurrencyRateFetchingInput input) async {
  final toCurrencyLower = input.to.toLowerCase();
  final fromCurrencyLower = input.from.toLowerCase();
  final url = 'https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/' +
      'currencies/$fromCurrencyLower/$toCurrencyLower.json';
  final response = await http.get(Uri.parse(url));
  if (response.statusCode == 200) {
    final data = json.decode(response.body);
    print(data);
    return data[toCurrencyLower];
  } else {
    throw Exception('Failed to load exchange rate');
  }
}
