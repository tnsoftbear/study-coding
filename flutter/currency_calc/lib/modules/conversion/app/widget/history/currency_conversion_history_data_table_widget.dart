import 'package:flutter/material.dart';

class CurrencyConversionHistoryDataTableWidget extends StatelessWidget {
  final List<Map<String, dynamic>> data;

  CurrencyConversionHistoryDataTableWidget(this.data);

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      scrollDirection: Axis.horizontal,
      child: DataTable(
        columnSpacing: 8,
        horizontalMargin: 16,
        decoration: BoxDecoration(
          color: Color.fromRGBO(255, 255, 255, 0.5),
          borderRadius: BorderRadius.circular(20),
        ),
        columns: [
          DataColumn(label: Text('Date')),
          DataColumn(label: Text('From'), tooltip: 'Source Currency'),
          DataColumn(label: Text('To'), tooltip: 'Target Currency'),
          DataColumn(label: Text('Rate')),
          DataColumn(label: Text(''), tooltip: 'Actions'),
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
      ),
    );
  }

  _onDeletePressed(int index) {
    print('Delete pressed for index $index');
  }
}
