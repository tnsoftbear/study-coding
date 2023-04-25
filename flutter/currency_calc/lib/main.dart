import 'package:currency_calc/modules/conversion/app/widget/history/model/currency_conversion_history_record.dart';
import 'package:flutter/material.dart';
import 'package:currency_calc/modules/front/app/widget/front_material_app.dart';
import 'package:hive_flutter/hive_flutter.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await Hive.initFlutter();
  Hive.registerAdapter(CurrencyConversionHistoryRecordAdapter());

  runApp(FrontMaterialApp());
}
