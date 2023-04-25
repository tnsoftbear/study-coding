import 'package:currency_calc/modules/front/app/constant/appearance_constant.dart';
import 'package:currency_calc/modules/front/app/widget/front_header_bar.dart';
import 'package:currency_calc/modules/front/app/widget/front_material_app.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class SettingPrimaryPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: FrontHeaderBar(
        titleText: AppLocalizations.of(context).settingTitle,
        isSettingMenu: false,
      ),
      body: Container(
        padding: const EdgeInsets.all(16),
        decoration: BoxDecoration(
          image: DecorationImage(
            image: AssetImage("assets/images/riga-cloudy-sky-landscape.jpg"),
            fit: BoxFit.cover,
          ),
        ),
        child: Center(
          child: Container(
            padding: const EdgeInsets.all(16),
            decoration: BoxDecoration(
              color: Color.fromRGBO(255, 255, 255, 0.8),
              borderRadius: BorderRadius.circular(8.0),
            ),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              mainAxisAlignment: MainAxisAlignment.end,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                _buildLocaleSetting(context),
                _buildFontFamilySetting(context),
              ],
            ),
          ),
        ),
      ),
    );
  }

  _onLocaleChange(BuildContext context, String? languageCode) {
    languageCode = languageCode ?? AppearanceConstant.LC_DEFAULT;
    final countryCode = AppearanceConstant.CONFIG[languageCode]!['countryCode'];
    final locale = Locale(languageCode, countryCode);
    FrontMaterialApp.setLocale(context, locale);
  }

  _buildLocaleSetting(BuildContext context) {
    final appLoc = AppLocalizations.of(context);
    List<Map<String, dynamic>> languages = [
      {
        'title': appLoc.settingLocaleEnLabel,
        'value': AppearanceConstant.LC_EN,
      },
      {
        'title': appLoc.settingLocaleRuLabel,
        'value': AppearanceConstant.LC_RU,
      },
    ];
    final List<Widget> localeList = languages.map((language) {
      return Expanded(
        child: RadioListTile(
          title: Text(language['title']),
          value: language['value'],
          groupValue: Localizations.localeOf(context).languageCode,
          onChanged: (languageCode) => _onLocaleChange(context, languageCode),
        ),
      );
    }).toList();
    final List<Widget> localeWidgetList =
        <Widget>[Text(appLoc.settingSelectLanguage)] + localeList;
    return Row(children: localeWidgetList);
  }

  _buildFontFamilySetting(BuildContext context) {
    final appLoc = AppLocalizations.of(context);
    final List<Widget> fontFamilyWidgetList = [
      Text(appLoc.settingSelectFontFamily),
      DropdownButton<String>(
        value: FrontMaterialApp.getFontFamily(context),
        onChanged: (String? fontFamily) =>
            FrontMaterialApp.setFontFamily(context, fontFamily),
        items: AppearanceConstant.FONT_FAMILIES
            .map<DropdownMenuItem<String>>((String value) {
          return DropdownMenuItem<String>(
            value: value,
            child: Text(value),
          );
        }).toList(),
      ),
    ];
    return Row(
        children: fontFamilyWidgetList,
        mainAxisAlignment: MainAxisAlignment.spaceBetween);
  }
}
