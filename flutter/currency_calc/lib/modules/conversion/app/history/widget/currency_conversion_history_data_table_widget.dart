import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:intl/intl.dart';

import 'currency_conversion_history_data_table_source.dart';

class CurrencyConversionHistoryDataTableWidget extends StatefulWidget {
  @override
  _CurrencyConversionHistoryDataTableWidget createState() =>
      _CurrencyConversionHistoryDataTableWidget();
}

class _CurrencyConversionHistoryDataTableWidget
    extends State<CurrencyConversionHistoryDataTableWidget> {
  late List<Map<String, String>> _historyRecords;

  @override
  void initState() {
    super.initState();
    _historyRecords = [];
  }

  @override
  Widget build(BuildContext context) {
    final tr = AppLocalizations.of(context);
    _loadHistoryRecords(context);

    return Container(
        width: 400,
        height: 200,
        decoration: BoxDecoration(
          color: Color.fromRGBO(255, 255, 255, 0.5),
          borderRadius: BorderRadius.circular(20),
        ),
        child: Theme(
            data: Theme.of(context)
                .copyWith(cardColor: Color.fromRGBO(255, 255, 255, 0.5)),
            child: PaginatedDataTable(
              rowsPerPage: 5,
              columnSpacing: 8,
              horizontalMargin: 8,
              columns: [
                DataColumn(
                    label: Text(tr.conversionHistoryDateColumnTitle),
                    tooltip: tr.conversionHistoryDateColumnTooltip),
                DataColumn(
                    label: Text(tr.conversionHistorySourceColumnTitle),
                    tooltip: tr.conversionHistorySourceColumnTooltip),
                DataColumn(
                    label: Text(tr.conversionHistoryTargetColumnTitle),
                    tooltip: tr.conversionHistoryTargetColumnTooltip),
                DataColumn(
                    label: Text(tr.conversionHistoryRateColumnTitle),
                    tooltip: tr.conversionHistoryRateColumnTooltip),
                DataColumn(
                    label: Text(tr.conversionHistoryActionsColumnTitle),
                    tooltip: tr.conversionHistoryActionsColumnTooltip),
              ],
              source: CurrencyConversionHistoryDataTableSource(_historyRecords),
            )));
  }

  Future<void> _loadHistoryRecords(BuildContext context) async {
    final localeName = Localizations.localeOf(context).toString();
    final df = DateFormat.yMMMd(localeName);
    final tf = DateFormat.Hms(localeName);
    final nf = NumberFormat.decimalPattern(localeName);
    List<Map<String, String>> historyRecords = [];
    var box = await Hive.openBox('currencyConversionHistory');
    historyRecords = box.values
        .toList()
        .map((e) => {
              'date': df.format(e.date) + "\n" + tf.format(e.date),
              'source_amount':
                  _formatCurrency(e.sourceAmount, e.sourceCurrency),
              'target_amount':
                  _formatCurrency(e.targetAmount, e.targetCurrency),
              'rate': nf.format(e.rate)
            })
        .toList();
    setState(() {
      _historyRecords = historyRecords;
    });
  }

  String _formatCurrency(double amount, String currencyCode) {
    final localeName = Localizations.localeOf(context).toString();
    final format = NumberFormat.simpleCurrency(
      locale: localeName,
      name: currencyCode,
    );
    return format.format(amount);
  }
}
