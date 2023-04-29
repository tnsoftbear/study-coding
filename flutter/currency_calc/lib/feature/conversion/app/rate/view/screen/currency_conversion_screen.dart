import 'package:currency_calc/feature/conversion/app/history/view/widget/last_history/conversion_last_history_data_table_widget.dart';
import 'package:currency_calc/feature/conversion/app/rate/view/widget/calculator/conversion_calculator_widget.dart';
import 'package:currency_calc/feature/front/app/view/widget/front_header_bar.dart';
import 'package:currency_calc/feature/front/app/view/widget/front_main_menu.dart';
import 'package:currency_calc/feature/front/app/constant/appearance_constant.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class CurrencyConversionScreen extends StatefulWidget {
  @override
  CurrencyConversionScreenState createState() =>
      CurrencyConversionScreenState();
}

class CurrencyConversionScreenState extends State<CurrencyConversionScreen> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final tr = AppLocalizations.of(context);
    return Scaffold(
      appBar: FrontHeaderBar(titleText: tr.conversionTitle),
      drawer: FrontMainMenu(),
      body: Container(
        padding: EdgeInsets.only(bottom: 16),
        decoration: const BoxDecoration(
          image: const DecorationImage(
            image: const AssetImage(
                AppearanceConstant.BG_IMAGE_FOR_CURRENCY_CONVERSION_SCREEN),
            fit: BoxFit.cover,
          ),
        ),
        child: Column(
          children: [
            Container(
              child: ConversionCalculatorWidget(),
            ),
            Container(
                height: 300,
                child: ConversionHistoryDataTableWidget()),
          ],
        ),
      ),
    );
  }
}
