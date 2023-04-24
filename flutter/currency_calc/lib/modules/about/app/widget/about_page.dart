import 'package:currency_calc/modules/front/app/widget/front_header_bar.dart';
import 'package:currency_calc/modules/front/app/widget/front_main_menu.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class AboutPage extends StatelessWidget {
  const AboutPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final appLoc = AppLocalizations.of(context);
    return Scaffold(
      appBar: FrontHeaderBar(titleText: appLoc.aboutTitle),
      drawer: FrontMainMenu(),
      body: Container(
        padding: const EdgeInsets.all(20),
        child: Center(
          child: Text(
            appLoc.aboutContent,
            style: TextStyle(
              fontSize: 20,
            ),
          ),
        ),
      ),
    );
  }
}
