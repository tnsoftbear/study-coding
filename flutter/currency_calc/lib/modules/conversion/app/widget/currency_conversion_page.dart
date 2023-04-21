import 'package:currency_calc/modules/front/app/widget/header_bar.dart';
import 'package:flutter/material.dart';

import 'currency_conversion_widget.dart';

class CurrencyConversionPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: HeaderBar(titleText: 'Currency Calculator'),
      body: CurrencyConversionWidget(),
    );
  }
}
