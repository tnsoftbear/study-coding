import 'package:currency_calc/modules/front/app/constant/route_constant.dart';
import 'package:flutter/material.dart';

import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class FrontMainMenu extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final appLoc = AppLocalizations.of(context);
    final menuItemStyle = TextStyle(
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
                child: const CircleAvatar(backgroundImage: AssetImage("assets/images/portugal-sea.jpg")),
              ),
            ),
            ListTile(
              leading: const Icon(
                Icons.info,
                color: Colors.white,
                size: 24,
              ),
              title: Text(
                appLoc.aboutTitle,
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
                appLoc.conversionTitle,
                style: menuItemStyle,
              ),
              onTap: () {
                // To close the Drawer
                Navigator.pop(context);
                // Navigating to About Page
                Navigator.pushNamed(context, RouteConstant.currencyConversionRoute);
              },
            ),
            ListTile(
              leading: const Icon(
                Icons.settings,
                color: Colors.white,
                size: 24,
              ),
              title: Text(
                appLoc.settingTitle,
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