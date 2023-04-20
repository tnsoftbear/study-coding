import 'package:currency_calc/modules/front/app/widget/HeaderBar.dart';
import 'package:flutter/material.dart';

import 'CurrencyConversionWidget.dart';

class CurrencyConversionPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: HeaderBar(titleText: 'Currency Calculator'),
      body: CurrencyConversionWidget(),
    );
  }
}
