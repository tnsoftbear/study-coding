import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:intl/intl.dart';

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
    final appLoc = AppLocalizations.of(context);
    _loadHistoryRecords(context);

    return DataTable(
      columnSpacing: 8,
      horizontalMargin: 16,
      decoration: BoxDecoration(
        color: Color.fromRGBO(255, 255, 255, 0.5),
        borderRadius: BorderRadius.circular(20),
      ),
      columns: [
        DataColumn(
            label: Text(appLoc.conversionHistoryDateColumnTitle),
            tooltip: appLoc.conversionHistoryDateColumnTooltip),
        DataColumn(
            label: Text(appLoc.conversionHistorySourceColumnTitle),
            tooltip: appLoc.conversionHistorySourceColumnTooltip),
        DataColumn(
            label: Text(appLoc.conversionHistoryTargetColumnTitle),
            tooltip: appLoc.conversionHistoryTargetColumnTooltip),
        DataColumn(
            label: Text(appLoc.conversionHistoryRateColumnTitle),
            tooltip: appLoc.conversionHistoryRateColumnTooltip),
        DataColumn(
            label: Text(appLoc.conversionHistoryActionsColumnTitle),
            tooltip: appLoc.conversionHistoryActionsColumnTooltip),
      ],
      rows: List.generate(
        _historyRecords.length,
            (index) => DataRow(
          cells: [
            DataCell(Text(_historyRecords[index]['date'].toString(),
                style: TextStyle(fontSize: 12))),
            DataCell(Text(_historyRecords[index]['source_amount'].toString())),
            DataCell(Text(_historyRecords[index]['target_amount'].toString())),
            DataCell(Text(_historyRecords[index]['rate'].toString())),
            DataCell(
              IconButton(
                icon: Icon(Icons.delete, size: 20, color: Colors.red),
                onPressed: () => _onDeletePressed(index),
              ),
            ),
          ],
        ),
      ),
    );
  }

  _onDeletePressed(int index) async {
    var box = await Hive.openBox('currencyConversionHistory');
    box.deleteAt(index);
  }

  Future<void> _loadHistoryRecords(BuildContext context) async {
    final localeName = Localizations.localeOf(context).toString();
    final df = DateFormat.yMMMd(localeName);
    final tf = DateFormat.Hms(localeName);
    final nf = NumberFormat.decimalPattern(localeName);
    List<Map<String, String>> historyRecords = [];
    var box = await Hive.openBox('currencyConversionHistory');
    final totalCount = box.length;
    final skipCount = totalCount > 5 ? totalCount - 5 : 0;
    historyRecords = box.values
        .skip(skipCount)
        .map((e) => {
      'date': df.format(e.date) + "\n" + tf.format(e.date),
      'source_amount':
      _formatCurrency(e.sourceAmount, e.sourceCurrency),
      'target_amount':
      _formatCurrency(e.targetAmount, e.targetCurrency),
      'rate': nf.format(e.rate)
    })
        .toList()
        .reversed
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
