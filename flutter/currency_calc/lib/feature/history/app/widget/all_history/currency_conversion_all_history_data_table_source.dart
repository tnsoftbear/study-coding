import 'package:currency_calc/feature/history/app/dto/currency_conversion_history_output_dto.dart';
import 'package:currency_calc/feature/history/infra/repository/currency_conversion_history_record_repository.dart';
import 'package:flutter/material.dart';

class CurrencyConversionAllHistoryDataTableSource extends DataTableSource {
  List<CurrencyConversionHistoryOutputDto> _historyRecords;

  CurrencyConversionAllHistoryDataTableSource(this._historyRecords);

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
        DataCell(
            Text(_historyRecords[index].date, style: TextStyle(fontSize: 12))),
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
    );
  }

  _onDeletePressed(int index) async {
    final repo = CurrencyConversionHistoryRecordRepository();
    await repo.init();
    repo.deleteByIndex(index);
  }
}