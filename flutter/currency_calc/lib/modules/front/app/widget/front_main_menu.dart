import 'package:currency_calc/modules/front/app/constant/route_constant.dart';
import 'package:currency_calc/modules/setting/app/constant/appearance_constant.dart';
import 'package:flutter/material.dart';

import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class FrontMainMenu extends StatelessWidget {
  const FrontMainMenu({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final tr = AppLocalizations.of(context);
    final menuItemStyle = const TextStyle(
      fontSize: 24,
      color: Colors.white,
    );
    return Drawer(
      child: Container(
        color: Theme.of(context).primaryColor,
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
                color: Colors.white,
                size: 24,
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
                color: Colors.white,
                size: 24,
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
                Icons.settings,
                color: Colors.white,
                size: 24,
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
