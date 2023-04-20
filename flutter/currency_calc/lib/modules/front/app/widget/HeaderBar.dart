import 'package:flutter/material.dart';

class HeaderBar extends AppBar {
  HeaderBar({Key? key, required this.titleText}) : super(key: key);

  final String titleText;

  @override
  _HeaderBarState createState() => _HeaderBarState();
}

class _HeaderBarState extends State<HeaderBar> {
  @override
  Widget build(BuildContext context) {
    return AppBar(
      title: Text(widget.titleText),
    );
  }
}
