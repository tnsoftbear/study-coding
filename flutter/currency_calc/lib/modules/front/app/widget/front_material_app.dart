import 'package:currency_calc/modules/about/app/widget/about_page.dart';
import 'package:currency_calc/modules/conversion/app/widget/currency_conversion_page.dart';
import 'package:currency_calc/modules/front/app/route/constant/route_constant.dart';
import 'package:currency_calc/modules/setting/app/widget/primary/setting_primary_page.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';
import 'package:shared_preferences/shared_preferences.dart';

class FrontMaterialApp extends StatefulWidget {
  const FrontMaterialApp({Key? key}) : super(key: key);

  State<FrontMaterialApp> createState() => _FrontMaterialAppState();

  static void setLocale(BuildContext context, Locale newLocale) async {
    _FrontMaterialAppState state =
        context.findAncestorStateOfType<_FrontMaterialAppState>()!;
    state.setLocale(newLocale);

    final SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setString("languageCode", newLocale.languageCode);
  }

  static void setFontFamily(BuildContext context, String newFontFamily) async {
    _FrontMaterialAppState state =
        context.findAncestorStateOfType<_FrontMaterialAppState>()!;
    state.setFontFamily(newFontFamily);

    final SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setString("fontFamily", newFontFamily);
  }

  Future<Locale> getLocale(BuildContext context) async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    final languageCode = await prefs.getString("languageCode") ?? 'en';
    final locale = Locale(languageCode);
    return locale;
  }

  Future<String> getFontFamily(BuildContext context) async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    final fontFamily = await prefs.getString("fontFamily") ?? 'Roboto';
    return fontFamily;
  }
}

class _FrontMaterialAppState extends State<FrontMaterialApp> {
  Locale _locale = Locale('en');
  String _fontFamily = 'Roboto';

  @override
  Widget build(BuildContext context) {
    widget.getLocale(context).then((locale) => _locale = locale);
    widget
        .getFontFamily(context)
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
      home: CurrencyConversionPage(),
      initialRoute: RouteConstant.currencyConversionRoute,
      routes: {
        RouteConstant.currencyConversionRoute: (context) =>
            CurrencyConversionPage(),
        RouteConstant.aboutRoute: (context) => AboutPage(),
        RouteConstant.settingRoute: (context) => SettingPrimaryPage(),
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
