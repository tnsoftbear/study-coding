import 'package:currency_calc/modules/conversion/app/widget/history/currency_conversion_history_data_table_widget.dart';
import 'package:currency_calc/modules/front/app/widget/front_header_bar.dart';
import 'package:currency_calc/modules/front/app/widget/front_main_menu.dart';
import 'package:flutter/material.dart';
import 'calculator/currency_conversion_calculator_widget.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class CurrencyConversionPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final appLoc = AppLocalizations.of(context);
    return Scaffold(
      appBar: FrontHeaderBar(titleText: appLoc.conversionTitle),
      drawer: FrontMainMenu(),
      body: Container(
        padding: EdgeInsets.only(bottom: 16),
        decoration: BoxDecoration(
          image: DecorationImage(
            image: AssetImage("assets/images/portugal-sea.jpg"),
            fit: BoxFit.cover,
          ),
        ),
        child: Column(
          children: [
            Expanded(
              flex: 1,
              child: CurrencyConversionCalculatorWidget(),
            ),
            Expanded(
                flex: 2,
                child: CurrencyConversionHistoryDataTableWidget()
            ),
          ],
        ),
      ),
    );
  }
}
