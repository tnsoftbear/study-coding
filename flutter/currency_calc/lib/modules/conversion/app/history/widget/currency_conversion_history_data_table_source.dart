import 'package:flutter/material.dart';
import 'package:hive_flutter/hive_flutter.dart';

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