import 'package:currency_calc/modules/front/app/widget/header_bar.dart';
import 'package:flutter/material.dart';

import 'currency_conversion_widget.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class CurrencyConversionPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: HeaderBar(titleText: AppLocalizations.of(context).conversionTitle),
      body: CurrencyConversionWidget(),
    );
  }
}
