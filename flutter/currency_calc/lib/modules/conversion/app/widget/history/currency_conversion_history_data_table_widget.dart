import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class CurrencyConversionHistoryDataTableWidget extends StatelessWidget {
  final List<Map<String, dynamic>> data;

  CurrencyConversionHistoryDataTableWidget(this.data);

  @override
  Widget build(BuildContext context) {
    final appLoc = AppLocalizations.of(context);
    return DataTable(
      columnSpacing: 8,
      horizontalMargin: 16,
      decoration: BoxDecoration(
        color: Color.fromRGBO(255, 255, 255, 0.5),
        borderRadius: BorderRadius.circular(20),
      ),
      columns: [
        DataColumn(label: Text(appLoc.conversionHistoryDateColumnTitle), tooltip: appLoc.conversionHistoryDateColumnTooltip),
        DataColumn(label: Text(appLoc.conversionHistorySourceColumnTitle), tooltip: appLoc.conversionHistorySourceColumnTooltip),
        DataColumn(label: Text(appLoc.conversionHistoryTargetColumnTitle), tooltip: appLoc.conversionHistoryTargetColumnTooltip),
        DataColumn(label: Text(appLoc.conversionHistoryRateColumnTitle), tooltip: appLoc.conversionHistoryRateColumnTooltip),
        DataColumn(label: Text(appLoc.conversionHistoryActionsColumnTitle), tooltip: appLoc.conversionHistoryActionsColumnTooltip),
      ],
      rows: List.generate(
        data.length,
        (index) => DataRow(
          cells: [
            DataCell(Text(data[index]['date'].toString())),
            DataCell(Text(data[index]['source_amount'])),
            DataCell(Text(data[index]['target_amount'])),
            DataCell(Text(data[index]['rate'].toString())),
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

  _onDeletePressed(int index) {
    print('Delete pressed for index $index');
  }
}
