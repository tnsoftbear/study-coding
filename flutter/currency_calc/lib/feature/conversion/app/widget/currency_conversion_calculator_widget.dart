import 'dart:developer';

import 'package:currency_calc/feature/conversion/app/config/currency_conversion_config.dart';
import 'package:currency_calc/feature/conversion/app/fetch/currency_rate_fetcher_factory.dart';
import 'package:currency_calc/feature/conversion/app/fetch/currency_rate_fetching_input.dart';
import 'package:currency_calc/feature/conversion/app/translate/currency_conversion_validation_translator.dart';
import 'package:currency_calc/feature/conversion/domain/calculate/currency_converter.dart';
import 'package:currency_calc/feature/conversion/domain/constant/currency_constant.dart';
import 'package:currency_calc/feature/conversion/domain/validate/currency_conversion_validator.dart';
import 'package:currency_calc/feature/history/domain/model/currency_conversion_history_record.dart';
import 'package:currency_calc/feature/history/infra/repository/currency_conversion_history_record_repository.dart';
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
  late bool _isLoading;
  late bool _areActionButtonsVisible;
  late double _rate;
  late double _sourceAmount;
  late double _targetAmount;
  late String _rateMessage;
  late String _resultMessage;
  late String _sourceAmountInput;
  late String _sourceCurrency;
  late String _targetCurrency;

  final _sourceAmountController = TextEditingController();
  final _sourceAmountTextFieldFocusNode = FocusNode();

  @override
  void initState() {
    super.initState();
    _isLoading = false;
    _areActionButtonsVisible = false;
    _rate = 0.0;
    _rateMessage = '';
    _resultMessage = '';
    _sourceAmount = 0.0;
    _sourceAmountInput = '';
    _sourceCurrency = CurrencyConstant.CURRENCIES[0];
    _targetAmount = 0.0;
    _targetCurrency = CurrencyConstant.CURRENCIES[1];
  }

  @override
  Widget build(BuildContext context) {
    final tr = AppLocalizations.of(context);
    return Padding(
      key: ValueKey("currencyConversionCalculatorWidget"),
      padding: const EdgeInsets.all(16.0),
      child: Container(
        decoration: const BoxDecoration(
          color: const Color.fromRGBO(255, 255, 255, 0.1),
          borderRadius: const BorderRadius.all(Radius.circular(8.0)),
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
                decoration: const BoxDecoration(
                  color: const Color.fromRGBO(255, 255, 255, 0.05),
                  borderRadius: const BorderRadius.all(Radius.circular(8.0)),
                ),
                child: TextField(
                  key: Key('sourceAmount'),
                  textAlign: TextAlign.center,
                  textAlignVertical: TextAlignVertical.center,
                  autofocus: true,
                  focusNode: _sourceAmountTextFieldFocusNode,
                  style: const TextStyle(
                    fontSize: 24.0,
                    fontWeight: FontWeight.bold,
                  ),
                  controller: _sourceAmountController,
                  keyboardType: TextInputType.number,
                  decoration: InputDecoration(
                    floatingLabelAlignment: FloatingLabelAlignment.center,
                    labelText: tr.conversionEnterAmount,
                  ),
                  onChanged: (text) {
                    setState(() {
                      _sourceAmountInput = text;
                      _updateConversion();
                    });
                  },
                )),
            Container(
              padding: const EdgeInsets.all(16.0),
              child: _isLoading
                  ? Center(child: const CircularProgressIndicator())
                  : Column(
                      children: [
                        Text(
                          _resultMessage,
                          style: const TextStyle(
                            fontSize: 24.0,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                        Text(
                          _rateMessage,
                          style: const TextStyle(
                            fontSize: 24.0,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                        Visibility(
                          visible: _areActionButtonsVisible,
                          child: Row(
                            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                            children: [
                              ElevatedButton(
                                  key: Key('saveCurrencyConversionButton'),
                                  onPressed: _onSavePressed,
                                  child: Text(tr.conversionSaveButtonText)),
                              ElevatedButton(
                                  key: Key('cancelCurrencyConversionButton'),
                                  onPressed: _onCancelPressed,
                                  child: Text(tr.conversionCancelButtonText))
                            ],
                          ),
                        ),
                      ],
                    ),
            ),
          ],
        ),
      ),
    );
  }

  void _updateConversion() {
    if (_sourceAmountInput.isEmpty) {
      return;
    }

    final tr = AppLocalizations.of(context);
    final validationResult = CurrencyConversionValidator.validate(
        sourceCurrency: _sourceCurrency,
        targetCurrency: _targetCurrency,
        amount: _sourceAmountInput);
    if (!validationResult.isSuccess()) {
      setState(() {
        _resultMessage = CurrencyConversionValidationTranslator
            .translateConcatenatedErrorMessage(
                context: context, validationResult: validationResult);
        _rateMessage = '';
        _areActionButtonsVisible = false;
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
      final localeName = Localizations.localeOf(context).toString();
      final currencyFormatter = NumberFormat.simpleCurrency(
          locale: localeName, name: _targetCurrency);
      final numberFormatter = NumberFormat.decimalPattern(localeName);
      setState(() {
        _rate = rate;
        _sourceAmount = double.parse(_sourceAmountInput);
        _targetAmount = CurrencyConverter.convert(_sourceAmount, rate);
        _resultMessage = tr.conversionCalculationResult(
            currencyFormatter.format(_targetAmount));
        final rateFormatted = numberFormatter.format(_rate);
        _rateMessage = tr.conversionRateResult(
            rateFormatted, _sourceCurrency, _targetCurrency);
        _isLoading = false;
        _areActionButtonsVisible = true;
      });
    }).catchError((error) {
      setState(() {
        _resultMessage = error.toString();
        _isLoading = false;
        _areActionButtonsVisible = false;
      });
    });
  }

  void _onSavePressed() async {
    var historyRecord = CurrencyConversionHistoryRecord()
      ..sourceCurrency = _sourceCurrency
      ..targetCurrency = _targetCurrency
      ..sourceAmount = _sourceAmount
      ..targetAmount = _targetAmount
      ..rate = _rate
      ..date = DateTime.now();
    final repo = CurrencyConversionHistoryRecordRepository();
    await repo.init();
    await repo.save(historyRecord);
    FocusScope.of(context).requestFocus(_sourceAmountTextFieldFocusNode);
    _resetInputs();
    log(
        'Saved currency conversion record (Source: $_sourceCurrency $_sourceAmount, ' +
            'Target: $_targetCurrency $_targetAmount, Rate: $_rate)',
        time: DateTime.now(),
        name: 'CurrencyConversionScreen');
  }

  void _onCancelPressed() {
    FocusScope.of(context).requestFocus(_sourceAmountTextFieldFocusNode);
    _resetInputs();
  }

  void _resetInputs() {
    setState(() {
      _areActionButtonsVisible = false;
      _rateMessage = '';
      _resultMessage = '';
      _sourceAmountController.clear();
      _sourceAmountInput = '';
    });
  }
}