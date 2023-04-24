import 'package:currency_calc/modules/setting/app/widget/primary/setting_primary_page.dart';
import 'package:flutter/material.dart';

class FrontHeaderBar extends AppBar {
  FrontHeaderBar({required this.titleText, bool this.isSettingMenu = true});

  final String titleText;
  final bool isSettingMenu;

  @override
  _HeaderBarState createState() => _HeaderBarState();
}

class _HeaderBarState extends State<FrontHeaderBar> {
  @override
  Widget build(BuildContext context) {
    final actions = widget.isSettingMenu
        ? <Widget>[
            IconButton(
              icon: Icon(Icons.settings),
              onPressed: () => Navigator.push(
                  context,
                  MaterialPageRoute(
                      builder: (context) => SettingPrimaryPage())),
            )
          ]
        : <Widget>[];

    return AppBar(
      title: Text(widget.titleText,
          style: TextStyle(
            fontSize: 26,
            color: Colors.white,
          )),
      backgroundColor: Colors.blue[300],
      actions: actions,
      iconTheme: IconThemeData(color: Colors.white),
      actionsIconTheme: IconThemeData(color: Colors.white),
    );
  }
}
