import 'package:currency_calc/modules/conversion/app/config/currency_conversion_config.dart';
import 'package:currency_calc/modules/conversion/app/constants/currency_constants.dart';
import 'package:currency_calc/modules/conversion/app/fetch/currency_rate_fetcher_factory.dart';
import 'package:currency_calc/modules/conversion/app/fetch/currency_rate_fetching_input.dart';
import 'package:currency_calc/modules/conversion/app/translate/currency_conversion_validation_translator.dart';
import 'package:currency_calc/modules/conversion/domain/calculator/currency_converter.dart';
import 'package:currency_calc/modules/conversion/domain/validate/currency_conversion_validator.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';
import 'package:intl/intl.dart';

class CurrencyConversionWidget extends StatefulWidget {
  @override
  _CurrencyConversionWidgetState createState() =>
      _CurrencyConversionWidgetState();
}

class _CurrencyConversionWidgetState extends State<CurrencyConversionWidget> {
  late String _sourceCurrency;
  late String _targetCurrency;
  late String _sourceAmount;
  late String _resultMessage;
  late String _rateMessage;
  late bool _isLoading;

  final _amountController = TextEditingController();

  @override
  void initState() {
    super.initState();
    _sourceCurrency = CurrencyConstants.CURRENCIES[0];
    _targetCurrency = CurrencyConstants.CURRENCIES[1];
    _sourceAmount = '';
    _resultMessage = '';
    _rateMessage = '';
    _isLoading = false;
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          DropdownButton<String>(
            value: _sourceCurrency,
            onChanged: (String? newValue) {
              setState(() {
                _sourceCurrency = newValue!;
                _updateConversion();
              });
            },
            items: CurrencyConstants.CURRENCIES
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
            items: CurrencyConstants.CURRENCIES
                .map<DropdownMenuItem<String>>((String value) {
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
              labelText: AppLocalizations.of(context).conversionEnterAmount,
            ),
            onChanged: (text) {
              setState(() {
                _sourceAmount = text;
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
    );
  }

  void _updateConversion() {
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
        _resultMessage = AppLocalizations.of(context)
            .conversionCalculationResult(targetAmount);
        final numberFormatter = NumberFormat.decimalPattern(
            Localizations.localeOf(context).toString());
        final rateFormatted = numberFormatter.format(ccResult.rate);
        _rateMessage = AppLocalizations.of(context).conversionRateResult(
            rateFormatted, _sourceCurrency, _targetCurrency);
        _isLoading = false;
      });
    }).catchError((error) {
      setState(() {
        _resultMessage = error.toString();
        _isLoading = false;
      });
    });
  }
}
