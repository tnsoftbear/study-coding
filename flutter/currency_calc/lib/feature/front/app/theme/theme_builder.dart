import 'package:currency_calc/feature/setting/app/constant/appearance_constant.dart';
import 'package:flutter/material.dart';

class ThemeBuilder {
  static ThemeData buildTheme(String? themeType, String? fontFamily) {
    fontFamily ??= AppearanceConstant.FF_DEFAULT;
    if (themeType == AppearanceConstant.THEME_GREEN) {
      return ThemeData(
        fontFamily: fontFamily,
        scaffoldBackgroundColor: Colors.green[300],
        colorScheme: ColorScheme.fromSwatch(
          primarySwatch: Colors.green,
        ),
        drawerTheme: DrawerThemeData(
          backgroundColor: Colors.green[300],
        ),
        appBarTheme: AppBarTheme(
          backgroundColor: Colors.green[300],
          titleTextStyle: TextStyle(
            fontSize: 26,
            color: Colors.white,
          ),
        ),
        iconTheme: const IconThemeData(color: Colors.white),
        listTileTheme: const ListTileThemeData(
          iconColor: Colors.white,
          style: ListTileStyle.drawer,
        ),
      );
    }

    if (themeType == AppearanceConstant.THEME_RED) {
      return ThemeData(
        fontFamily: fontFamily,
        scaffoldBackgroundColor: Colors.red[300],
        colorScheme: ColorScheme.fromSwatch(
          primarySwatch: Colors.red,
        ),
        drawerTheme: DrawerThemeData(
          backgroundColor: Colors.red[300],
        ),
        appBarTheme: AppBarTheme(
            backgroundColor: Colors.red[300],
            titleTextStyle: TextStyle(
              fontSize: 26,
              color: Colors.white,
            )),
        iconTheme: const IconThemeData(color: Colors.white),
        listTileTheme: const ListTileThemeData(
          iconColor: Colors.white,
          style: ListTileStyle.drawer,
        ),
      );
    }

    return ThemeData(
      fontFamily: fontFamily,
      scaffoldBackgroundColor: Colors.blue[300],
      colorScheme: ColorScheme.fromSwatch(
        primarySwatch: Colors.blue,
      ),
      drawerTheme: DrawerThemeData(
        backgroundColor: Colors.blue[300],
      ),
      appBarTheme: AppBarTheme(
          backgroundColor: Colors.blue[300],
          titleTextStyle: TextStyle(
            fontSize: 26,
            color: Colors.white,
          )),
      iconTheme: const IconThemeData(color: Colors.white),
      listTileTheme: const ListTileThemeData(
        iconColor: Colors.white,
        style: ListTileStyle.drawer,
      ),
    );
  }
}