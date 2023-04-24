import 'package:currency_calc/modules/conversion/app/config/currency_conversion_config.dart';
import 'package:currency_calc/modules/conversion/app/constant/currency_constant.dart';
import 'package:currency_calc/modules/conversion/app/fetch/currency_rate_fetcher_factory.dart';
import 'package:currency_calc/modules/conversion/app/fetch/currency_rate_fetching_input.dart';
import 'package:currency_calc/modules/conversion/app/translate/currency_conversion_validation_translator.dart';
import 'package:currency_calc/modules/conversion/domain/calculator/currency_converter.dart';
import 'package:currency_calc/modules/conversion/domain/validate/currency_conversion_validator.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';
import 'package:intl/intl.dart';

class CurrencyConversionCalculatorWidget extends StatefulWidget {
  @override
  _CurrencyConversionCalculatorWidgetState createState() =>
      _CurrencyConversionCalculatorWidgetState();
}

class _CurrencyConversionCalculatorWidgetState
    extends State<CurrencyConversionCalculatorWidget> {
  late String _sourceCurrency;
  late String _targetCurrency;
  late String _sourceAmount;
  late String _resultMessage;
  late String _rateMessage;
  late bool _isLoading;
  late bool _isSaveButtonVisible;

  final _amountController = TextEditingController();

  @override
  void initState() {
    super.initState();
    _sourceCurrency = CurrencyConstant.CURRENCIES[0];
    _targetCurrency = CurrencyConstant.CURRENCIES[1];
    _sourceAmount = '';
    _resultMessage = '';
    _rateMessage = '';
    _isLoading = false;
    _isSaveButtonVisible = false;
  }

  @override
  Widget build(BuildContext context) {
    final appLoc = AppLocalizations.of(context);
    return Padding(
      key: ValueKey("currencyConversionCalculatorWidget"),
      padding: EdgeInsets.all(16.0),
      child: Container(
        decoration: BoxDecoration(
          color: Color.fromRGBO(255, 255, 255, 0.1),
          borderRadius: BorderRadius.circular(8.0),
        ),
        child: Column(
          children: [
            Row(mainAxisAlignment: MainAxisAlignment.spaceEvenly, children: [
              DropdownButton<String>(
                value: _sourceCurrency,
                onChanged: (String? newValue) {
                  setState(() {
                    _sourceCurrency = newValue!;
                    _updateConversion();
                  });
                },
                items: CurrencyConstant.CURRENCIES
                    .map<DropdownMenuItem<String>>((String value) {
                  return DropdownMenuItem<String>(
                    value: value,
                    child: Text(value),
                  );
                }).toList(),
              ),
              DropdownButton<String>(
                value: _targetCurrency,
                onChanged: (String? newValue) {
                  setState(() {
                    _targetCurrency = newValue!;
                    _updateConversion();
                  });
                },
                items: CurrencyConstant.CURRENCIES
                    .map<DropdownMenuItem<String>>((String value) {
                  return DropdownMenuItem<String>(
                    value: value,
                    child: Text(value),
                  );
                }).toList(),
              ),
            ]),
            Container(
                width: 200,
                alignment: Alignment.center,
                decoration: BoxDecoration(
                  color: Color.fromRGBO(255, 255, 255, 0.05),
                  borderRadius: BorderRadius.circular(8.0),
                ),
                child: TextField(
                  key: Key('sourceAmount'),
                  textAlign: TextAlign.center,
                  textAlignVertical: TextAlignVertical.center,
                  autofocus: true,
                  style: TextStyle(
                    fontSize: 24.0,
                    fontWeight: FontWeight.bold,
                  ),
                  controller: _amountController,
                  keyboardType: TextInputType.number,
                  decoration: InputDecoration(
                    floatingLabelAlignment: FloatingLabelAlignment.center,
                    labelText:
                        appLoc.conversionEnterAmount,
                  ),
                  onChanged: (text) {
                    setState(() {
                      _sourceAmount = text;
                      _updateConversion();
                    });
                  },
                )),
            Container(
              padding: EdgeInsets.all(16.0),
              child: _isLoading
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
                        Visibility(
                            visible: _isSaveButtonVisible,
                            child: ElevatedButton(
                                key: Key('saveCurrencyConversionButton'),
                                onPressed: _onSavePressed,
                                child: Text('Save'))),
                      ],
                    ),
            ),
          ],
        ),
      ),
    );
  }

  void _updateConversion() {
    final appLoc = AppLocalizations.of(context);
    final validationResult = CurrencyConversionValidator.validate(
        sourceCurrency: _sourceCurrency,
        targetCurrency: _targetCurrency,
        amount: _sourceAmount);
    if (!validationResult.isSuccess()) {
      setState(() {
        _resultMessage = CurrencyConversionValidationTranslator
            .translateConcatenatedErrorMessage(
                context: context, validationResult: validationResult);
        _rateMessage = '';
        _isSaveButtonVisible = false;
      });
      return;
    }

    setState(() {
      _isLoading = true;
    });

    final input =
        CurrencyRateFetchingInput(from: _sourceCurrency, to: _targetCurrency);
    final rateFetcher =
        CurrencyRateFetcherFactory.create(config: CurrencyConversionConfig());
    rateFetcher.fetchExchangeRate(input).then((rate) {
      setState(() {
        double sourceAmount = double.parse(_sourceAmount);
        final ccResult = CurrencyConverter.convert(sourceAmount, rate);
        final currencyFormatter = NumberFormat.simpleCurrency(
            locale: Localizations.localeOf(context).toString(),
            name: _targetCurrency);
        final targetAmount = currencyFormatter.format(ccResult.targetAmount);
        _resultMessage = appLoc
            .conversionCalculationResult(targetAmount);
        final numberFormatter = NumberFormat.decimalPattern(
            Localizations.localeOf(context).toString());
        final rateFormatted = numberFormatter.format(ccResult.rate);
        _rateMessage = appLoc.conversionRateResult(
            rateFormatted, _sourceCurrency, _targetCurrency);
        _isLoading = false;
        _isSaveButtonVisible = true;
      });
    }).catchError((error) {
      setState(() {
        _resultMessage = error.toString();
        _isLoading = false;
        _isSaveButtonVisible = false;
      });
    });
  }

  void _onSavePressed() {

  }
}
