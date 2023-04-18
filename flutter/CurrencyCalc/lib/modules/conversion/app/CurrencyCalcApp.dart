import 'package:CurrencyCalc/modules/conversion/app/CurrencyConversionPage.dart';
import 'package:flutter/material.dart';

class CurrencyCalcApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Currency Calculator',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: CurrencyConversionPage(title: 'Currency Calculator'),
    );
  }
}