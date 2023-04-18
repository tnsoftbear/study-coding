import 'package:CurrencyCalc/modules/conversion/domain/calculator/CurrencyConverter.dart';
import 'package:CurrencyCalc/modules/conversion/infra/CurrencyRateFetchingInput.dart';
import 'package:flutter/material.dart';
// import 'package:CurrencyCalc/modules/conversion/infra/FixerIoCurrencyConversionFetcher.dart';
import 'package:CurrencyCalc/modules/conversion/infra/FawazAhmedCurrencyConversionFetcher.dart';

class CurrencyConversionPage extends StatefulWidget {
  CurrencyConversionPage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  _CurrencyConversionPageState createState() => _CurrencyConversionPageState();
}

class _CurrencyConversionPageState extends State<CurrencyConversionPage> {
  late String _fromCurrency;
  late String _toCurrency;
  late double _sourceAmount;
  late String _result;
  late String _rate;
  late bool _isLoading;

  final _amountController = TextEditingController();

  final _currencies = [
    'USD',
    'EUR',
    'GBP',
    'AUD',
    'CAD',
    'JPY',
    'CHF',
    'NZD',
    'CNY',
    'HKD',
  ];

  @override
  void initState() {
    super.initState();
    _fromCurrency = _currencies[0];
    _toCurrency = _currencies[1];
    _sourceAmount = 0.0;
    _result = '';
    _rate = '';
    _isLoading = false;
  }

  void _updateConversion() {
    if (_sourceAmount <= 0) {
      setState(() {
        _result = 'Enter positive non-zero amount';
        _rate = '';
      });
      return;
    }

    setState(() {
      _isLoading = true;
    });

    final input = CurrencyRateFetchingInput(from: _fromCurrency, to: _toCurrency);

    fetchExchangeRate(input).then((rate) {
      setState(() {
        final ccResult = CurrencyConverter.convert(_sourceAmount, rate);
        _result = 'Result: ' + ccResult.targetAmount.toString() + ' ' + _toCurrency;
        _rate = _fromCurrency + ' to ' + _toCurrency + ' rate: ' + ccResult.rate.toString();
        _isLoading = false;
      });
    }).catchError((error) {
      setState(() {
        _result = error.toString();
        _isLoading = false;
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Padding(
        padding: EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            DropdownButton<String>(
              value: _fromCurrency,
              onChanged: (String? newValue) {
                setState(() {
                  _fromCurrency = newValue!;
                  _updateConversion();
                });
              },
              items: _currencies.map<DropdownMenuItem<String>>((String value) {
                return DropdownMenuItem<String>(
                  value: value,
                  child: Text(value),
                );
              }).toList(),
            ),
            DropdownButton<String>(
              value: _toCurrency,
              onChanged: (String? newValue) {
                setState(() {
                  _toCurrency = newValue!;
                  _updateConversion();
                });
              },
              items: _currencies.map<DropdownMenuItem<String>>((String value) {
                return DropdownMenuItem<String>(
                  value: value,
                  child: Text(value),
                );
              }).toList(),
            ),
            TextField(
              controller: _amountController,
              keyboardType: TextInputType.number,
              decoration: InputDecoration(
                labelText: 'Enter amount',
              ),
              onChanged: (value) {
                setState(() {
                  _sourceAmount = double.tryParse(value) ?? 0.0;
                  _updateConversion();
                });
              },
            ),
            SizedBox(height: 20),
            _isLoading
                ? Center(child: CircularProgressIndicator())
                : Column(
              children: [
                Text(
                  _result,
                  style: TextStyle(
                    fontSize: 24.0,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                Text(
                  _rate,
                  style: TextStyle(
                    fontSize: 24.0,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}