import 'package:currency_calc/modules/conversion/app/config/CurrencyConversionConfig.dart';
import 'package:currency_calc/modules/conversion/app/constants/CurrencyConstants.dart';
import 'package:currency_calc/modules/conversion/app/fetch/CurrencyRateFetcherFactory.dart';
import 'package:currency_calc/modules/conversion/domain/calculator/CurrencyConverter.dart';
import 'package:currency_calc/modules/conversion/app/fetch/CurrencyRateFetchingInput.dart';
import 'package:flutter/material.dart';

class CurrencyConversionPage extends StatefulWidget {
  CurrencyConversionPage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  _CurrencyConversionPageState createState() => _CurrencyConversionPageState();
}

class _CurrencyConversionPageState extends State<CurrencyConversionPage> {
  late String _fromCurrency;
  late String _toCurrency;
  late String _sourceAmount;
  late String _resultMessage;
  late String _rateMessage;
  late bool _isLoading;

  final _amountController = TextEditingController();

  @override
  void initState() {
    super.initState();
    _fromCurrency = CurrencyConstants.CURRENCIES[0];
    _toCurrency = CurrencyConstants.CURRENCIES[1];
    _sourceAmount = '';
    _resultMessage = '';
    _rateMessage = '';
    _isLoading = false;
  }

  void _updateConversion() {
    double? sourceAmount = double.tryParse(_sourceAmount);
    if (sourceAmount == null) {
      setState(() {
        _resultMessage = 'Numeric amount expected';
        _rateMessage = '';
      });
      return;
    }

    if (sourceAmount <= 0) {
      setState(() {
        _resultMessage = 'Enter positive non-zero amount';
        _rateMessage = '';
      });
      return;
    }

    setState(() {
      _isLoading = true;
    });

    final input =
        CurrencyRateFetchingInput(from: _fromCurrency, to: _toCurrency);
    final rateFetcher =
        CurrencyRateFetcherFactory.create(config: CurrencyConversionConfig());
    rateFetcher.fetchExchangeRate(input).then((rate) {
      setState(() {
        final ccResult = CurrencyConverter.convert(sourceAmount, rate);
        _resultMessage =
            'Result: ${ccResult.targetAmount.toString()} $_toCurrency';
        _rateMessage =
            '$_fromCurrency to $_toCurrency rate: ${ccResult.rate.toString()}';
        _isLoading = false;
      });
    }).catchError((error) {
      setState(() {
        _resultMessage = error.toString();
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
              items: CurrencyConstants.CURRENCIES.map<DropdownMenuItem<String>>((String value) {
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
              items: CurrencyConstants.CURRENCIES.map<DropdownMenuItem<String>>((String value) {
                return DropdownMenuItem<String>(
                  value: value,
                  child: Text(value),
                );
              }).toList(),
            ),
            TextField(
              key: Key('sourceAmount'),
              controller: _amountController,
              keyboardType: TextInputType.number,
              decoration: InputDecoration(
                labelText: 'Enter amount',
              ),
              onChanged: (value) {
                setState(() {
                  _sourceAmount = value;
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
                        _resultMessage,
                        style: TextStyle(
                          fontSize: 24.0,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                      Text(
                        _rateMessage,
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
