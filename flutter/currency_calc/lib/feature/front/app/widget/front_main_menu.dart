import 'package:currency_calc/feature/front/app/constant/route_constant.dart';
import 'package:currency_calc/feature/setting/app/constant/appearance_constant.dart';
import 'package:flutter/material.dart';

import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class FrontMainMenu extends StatelessWidget {
  const FrontMainMenu({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final tr = AppLocalizations.of(context);
    final menuItemStyle = const TextStyle(
      fontSize: 20,
      color: Colors.white,
    );
    return Drawer(
      child: Container(
        child: ListView(
          padding: EdgeInsets.zero,
          children: <Widget>[
            DrawerHeader(
              child: Container(
                height: 100,
                child: const CircleAvatar(
                    backgroundImage: const AssetImage(
                        AppearanceConstant.BG_IMAGE_FOR_MAIN_MENU_AVATAR)),
              ),
            ),
            ListTile(
              leading: const Icon(
                Icons.info,
                size: 20,
              ),
              title: Text(
                tr.aboutTitle,
                style: menuItemStyle,
              ),
              onTap: () {
                // To close the Drawer
                Navigator.pop(context);
                // Navigating to About Page
                Navigator.pushNamed(context, RouteConstant.aboutRoute);
              },
            ),
            ListTile(
              leading: const Icon(
                Icons.calculate,
                size: 20,
              ),
              title: Text(
                tr.conversionTitle,
                style: menuItemStyle,
              ),
              onTap: () {
                // To close the Drawer
                Navigator.pop(context);
                // Navigating to About Page
                Navigator.pushNamed(
                    context, RouteConstant.currencyConversionRoute);
              },
            ),
            ListTile(
              leading: const Icon(
                Icons.history,
                size: 20,
              ),
              title: Text(
                tr.conversionHistoryTitle,
                style: menuItemStyle,
              ),
              onTap: () {
                // To close the Drawer
                Navigator.pop(context);
                // Navigating to About Page
                Navigator.pushNamed(context, RouteConstant.currencyConversionAllHistoryRoute);
              },
            ),
            ListTile(
              leading: const Icon(
                Icons.settings,
                size: 20,
              ),
              title: Text(
                tr.settingTitle,
                style: menuItemStyle,
              ),
              onTap: () {
                // To close the Drawer
                Navigator.pop(context);
                // Navigating to About Page
                Navigator.pushNamed(context, RouteConstant.settingRoute);
              },
            ),
          ],
        ),
      ),
    );
  }
}
