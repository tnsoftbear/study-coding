import 'package:CurrencyCalc/modules/conversion/infra/CurrencyRateFetchingInput.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

Future<double> fetchExchangeRate(
    CurrencyRateFetchingInput input) async {
  final apiKey = '1yUWc2Kb2Bzr13w7hryFnkKBCxGV38Ia';
  final fromCurrency = input.from;
  final toCurrency = input.to;
  final url =
      'https://api.apilayer.com/fixer/convert?from=$fromCurrency&to=$toCurrency&amount=1';
  final response = await http.get(Uri.parse(url), headers: {'apikey': apiKey});
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
