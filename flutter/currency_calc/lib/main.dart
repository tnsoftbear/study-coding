import 'package:currency_calc/feature/conversion/domain/model/fetch/currency_conversion_rate_fetch_record.dart';
import 'package:currency_calc/feature/conversion/domain/model/history/currency_conversion_history_record.dart';
import 'package:flutter/material.dart';
import 'package:currency_calc/feature/front/app/widget/front_material_app.dart';
import 'package:flutter/services.dart';
import 'package:hive_flutter/hive_flutter.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await Hive.initFlutter();
  Hive.registerAdapter(CurrencyConversionHistoryRecordAdapter());
  Hive.registerAdapter(CurrencyConversionRateFetchRecordAdapter());

  await SystemChrome.setPreferredOrientations(
      [DeviceOrientation.portraitUp, DeviceOrientation.portraitDown]);

  runApp(FrontMaterialApp());
}
