import 'package:currency_calc/modules/about/app/screen/about_screen.dart';
import 'package:currency_calc/modules/conversion/app/screen/currency_conversion_screen.dart';
import 'package:currency_calc/modules/front/app/constant/route_constant.dart';
import 'package:currency_calc/modules/setting/app/manage/setting_manager.dart';
import 'package:currency_calc/modules/setting/app/screen/setting_primary_screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class FrontMaterialApp extends StatefulWidget {
  const FrontMaterialApp({Key? key}) : super(key: key);

  State<FrontMaterialApp> createState() => _FrontMaterialAppState();

  static void assignLocale(BuildContext context, Locale newLocale) async {
    _FrontMaterialAppState state =
        context.findAncestorStateOfType<_FrontMaterialAppState>()!;
    state.setLocale(newLocale);

    SettingManager.saveLanguageCode(newLocale.languageCode);
  }

  static void assignFontFamily(
      BuildContext context, String? newFontFamily) async {
    if (newFontFamily == null) {
      return;
    }

    _FrontMaterialAppState state =
        context.findAncestorStateOfType<_FrontMaterialAppState>()!;
    state.setFontFamily(newFontFamily);

    SettingManager.saveFontFamily(newFontFamily);
  }

  static String getFontFamily(BuildContext context) {
    _FrontMaterialAppState state =
        context.findAncestorStateOfType<_FrontMaterialAppState>()!;
    return state._fontFamily!;
  }
}

class _FrontMaterialAppState extends State<FrontMaterialApp> {
  Locale? _locale;
  String? _fontFamily;

  @override
  Widget build(BuildContext context) {
    SettingManager.detectLocale().then((locale) => _locale = locale);
    SettingManager.detectFontFamily()
        .then((fontFamily) => _fontFamily = fontFamily);

    return MaterialApp(
      debugShowCheckedModeBanner: false,
      onGenerateTitle: (context) => AppLocalizations.of(context).appTitle,
      theme: ThemeData(
        primarySwatch: Colors.blue,
        fontFamily: _fontFamily,
      ),
      localizationsDelegates: AppLocalizations.localizationsDelegates,
      supportedLocales: AppLocalizations.supportedLocales,
      locale: _locale,
      home: CurrencyConversionScreen(),
      initialRoute: RouteConstant.currencyConversionRoute,
      routes: {
        RouteConstant.currencyConversionRoute: (context) =>
            CurrencyConversionScreen(),
        RouteConstant.aboutRoute: (context) => AboutScreen(),
        RouteConstant.settingRoute: (context) => SettingPrimaryScreen(),
      },
    );
  }

  void setLocale(Locale newLocale) {
    setState(() {
      _locale = newLocale;
    });
  }

  void setFontFamily(String newFontFamily) {
    setState(() {
      _fontFamily = newFontFamily;
    });
  }
}