import 'package:currency_calc/feature/front/app/widget/front_header_bar.dart';
import 'package:currency_calc/feature/front/app/widget/front_main_menu.dart';
import 'package:currency_calc/feature/setting/app/constant/appearance_constant.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/all_localizations.dart';

class AboutScreen extends StatelessWidget {
  const AboutScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final tr = AppLocalizations.of(context);
    return Scaffold(
      appBar: FrontHeaderBar(titleText: tr.aboutTitle),
      drawer: const FrontMainMenu(),
      body: Container(
        padding: const EdgeInsets.all(16),
        decoration: const BoxDecoration(
          image: const DecorationImage(
            image: const AssetImage(AppearanceConstant.BG_IMAGE_FOR_ABOUT_SCREEN),
            fit: BoxFit.cover,
          ),
        ),
        child: Center(
          child: Container(
            padding: const EdgeInsets.all(16),
            decoration: const BoxDecoration(
              color: const Color.fromRGBO(255, 255, 255, 0.8),
              borderRadius: const BorderRadius.all(Radius.circular(8.0)),
            ),
            child: Text(
              tr.aboutContent,
              style: const TextStyle(
                fontSize: 20,
              ),
            ),
          ),
        ),
      ),
    );
  }
}