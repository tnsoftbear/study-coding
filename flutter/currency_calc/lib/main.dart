import 'package:currency_calc/modules/conversion/domain/history/model/currency_conversion_history_record.dart';
import 'package:flutter/material.dart';
import 'package:currency_calc/modules/front/app/widget/front_material_app.dart';
import 'package:flutter/services.dart';
import 'package:hive_flutter/hive_flutter.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await Hive.initFlutter();
  Hive.registerAdapter(CurrencyConversionHistoryRecordAdapter());

  await SystemChrome.setPreferredOrientations(
      [DeviceOrientation.portraitUp, DeviceOrientation.portraitDown]);

  runApp(FrontMaterialApp());
}
