import 'package:currency_calc/feature/setting/app/constant/appearance_constant.dart';
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class SettingManager {
  static Future<String> detectLanguageCode() async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    final languageCode =
        await prefs.getString("languageCode") ?? AppearanceConstant.LC_DEFAULT;
    return languageCode;
  }

  static Future<Locale> detectLocale() async {
    final languageCode = await detectLanguageCode();
    final locale = Locale(languageCode);
    return locale;
  }

  static Future<void> saveLanguageCode(String languageCode) async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setString("languageCode", languageCode);
  }

  static Future<String> detectFontFamily() async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    final fontFamily =
        await prefs.getString("fontFamily") ?? AppearanceConstant.FF_DEFAULT;
    return fontFamily;
  }

  static Future<void> saveFontFamily(String fontFamily) async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setString("fontFamily", fontFamily);
  }

  static Future<String> detectThemeType() async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    final themeType = await prefs.getString("themeType") ?? AppearanceConstant.THEME_DEFAULT;
    return themeType;
  }

  static Future<void> saveThemeType(String themeType) async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setString("themeType", themeType);
  }
}
