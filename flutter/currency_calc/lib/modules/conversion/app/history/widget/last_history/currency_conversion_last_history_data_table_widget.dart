import 'package:currency_calc/modules/conversion/app/constant/currency_constant.dart';
import 'package:currency_calc/modules/conversion/app/history/model/currency_conversion_history_output_data.dart';
import 'package:currency_calc/modules/conversion/infra/history/repository/currency_conversion_history_record_repository.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';
import 'package:intl/intl.dart';

class CurrencyConversionHistoryDataTableWidget extends StatefulWidget {
  @override
  _CurrencyConversionHistoryDataTableWidget createState() =>
      _CurrencyConversionHistoryDataTableWidget();
}

class _CurrencyConversionHistoryDataTableWidget
    extends State<CurrencyConversionHistoryDataTableWidget> {
  late List<CurrencyConversionHistoryOutputData> _historyRecords;

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
            DataCell(Text(_historyRecords[index].date,
                style: TextStyle(fontSize: 12))),
            DataCell(Text(_historyRecords[index].from)),
            DataCell(Text(_historyRecords[index].to)),
            DataCell(Text(_historyRecords[index].rate)),
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
    final repo = CurrencyConversionHistoryRecordRepository();
    await repo.init();
    final deleteIndex = repo.countAll() - index - 1;
    await repo.deleteByIndex(deleteIndex);
  }

  Future<void> _loadHistoryRecords(BuildContext context) async {
    final localeName = Localizations.localeOf(context).toString();
    final df = DateFormat.yMMMd(localeName);
    final tf = DateFormat.Hms(localeName);
    final nf = NumberFormat.decimalPattern(localeName);
    final repo = CurrencyConversionHistoryRecordRepository();
    await repo.init();
    final totalCount = repo.countAll();
    final skipCount = totalCount > CurrencyConstant.LAST_HISTORY_RECORD_COUNT
        ? totalCount - CurrencyConstant.LAST_HISTORY_RECORD_COUNT
        : 0;
    final historyRecords = repo.loadAll() // box.values
        .skip(skipCount)
        .map((e) => CurrencyConversionHistoryOutputData(
            df.format(e.date) + "\n" + tf.format(e.date),
            _formatCurrency(e.sourceAmount, e.sourceCurrency),
            _formatCurrency(e.targetAmount, e.targetCurrency),
            nf.format(e.rate)))
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
