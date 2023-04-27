import 'package:currency_calc/modules/conversion/app/calculate/widget/currency_conversion_calculator_widget.dart';
import 'package:currency_calc/modules/conversion/app/history/widget/last_history/currency_conversion_last_history_data_table_widget.dart';
import 'package:currency_calc/modules/front/app/widget/front_header_bar.dart';
import 'package:currency_calc/modules/front/app/widget/front_main_menu.dart';
import 'package:currency_calc/modules/setting/app/constant/appearance_constant.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class CurrencyConversionScreen extends StatefulWidget {
  @override
  CurrencyConversionScreenState createState() => CurrencyConversionScreenState();
}

class CurrencyConversionScreenState extends State<CurrencyConversionScreen> {
  late bool _isCurrencyConversionHistoryVisible;

  set isCurrencyConversionHistoryVisible(bool enable) =>
      setState(() => _isCurrencyConversionHistoryVisible = enable);

  @override
  void initState() {
    super.initState();
    _isCurrencyConversionHistoryVisible = true;
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
            Expanded(
              flex: 1,
              child: CurrencyConversionCalculatorWidget(),
            ),
            Visibility(
              visible: _isCurrencyConversionHistoryVisible,
              child: Expanded(
                  flex: 2, child: CurrencyConversionHistoryDataTableWidget()),
            ),
          ],
        ),
      ),
    );
  }
}
