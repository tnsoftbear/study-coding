import 'package:currency_calc/modules/conversion/app/widget/currency_conversion_page.dart';
import 'package:flutter/material.dart';

class CurrencyCalcApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Currency Calculator',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: CurrencyConversionPage(),
    );
  }
}