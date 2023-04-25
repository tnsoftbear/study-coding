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

    return Container(
        width: 400,
        height: 200,
        decoration: BoxDecoration(
          color: Color.fromRGBO(255, 255, 255, 0.5),
          borderRadius: BorderRadius.circular(20),
        ),
        child: Theme(
            data: Theme.of(context).copyWith(
                cardColor: Color.fromRGBO(255, 255, 255, 0.5)),
            child: PaginatedDataTable(
              rowsPerPage: 5,
              columnSpacing: 8,
              horizontalMargin: 8,
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

class CurrencyConversionHistoryDataTableSource extends DataTableSource {
  List<Map<String, String>> _historyRecords;

  CurrencyConversionHistoryDataTableSource(this._historyRecords);

  @override
  bool get isRowCountApproximate => false;

  @override
  int get rowCount => _historyRecords.length;

  @override
  int get selectedRowCount => 0;

  DataRow getRow(int index) {
    index = _historyRecords.length - index - 1;
    return DataRow(
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
    );
  }

  _onDeletePressed(int index) async {
    var box = await Hive.openBox('currencyConversionHistory');
    await box.deleteAt(index);
    await box.close();
  }
}
