import 'package:currency_calc/modules/front/app/widget/front_header_bar.dart';
import 'package:currency_calc/modules/front/app/widget/front_material_app.dart';
import 'package:flutter/material.dart';

import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class SettingPrimaryPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final appLoc = AppLocalizations.of(context);
    return Scaffold(
      appBar: FrontHeaderBar(
        titleText: AppLocalizations.of(context).settingTitle,
        isSettingMenu: false,
      ),
      body: Container(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          mainAxisAlignment: MainAxisAlignment.end,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(children: [
              Expanded(child: Text(appLoc.settingSelectLanguage)),
              Expanded(
                child: RadioListTile(
                  title: Text(appLoc.settingLocaleEnLabel),
                  value: Locale('en'),
                  groupValue: Localizations.localeOf(context),
                  onChanged: (locale) => onLocaleChange(context, locale!),
                ),
              ),
              Expanded(
                child: RadioListTile(
                  title: Text(appLoc.settingLocaleRuLabel),
                  value: Locale('ru'),
                  groupValue: Localizations.localeOf(context),
                  onChanged: (locale) => onLocaleChange(context, locale!),
                ),
              ),
            ])
          ],
        ),
      ),
    );
  }

  onLocaleChange(BuildContext context, Locale locale) {
    FrontMaterialApp.setLocale(context, Locale(locale.toString()));
    if (locale.languageCode == 'en') {
      FrontMaterialApp.setFontFamily(context, 'IndieFlower');
    } else {
      FrontMaterialApp.setFontFamily(context, 'Roboto');
    }
  }
}
