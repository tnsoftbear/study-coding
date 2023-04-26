class AppearanceConstant {
  // Language Codes
  static const String LC_DEFAULT = LC_EN;
  static const String LC_EN = 'en';
  static const String LC_RU = 'ru';

  static const Map<String, Map<String, String>> CONFIG = {
    LC_EN: {'countryCode': 'US'},
    LC_RU: {'countryCode': 'RU'},
  };

  // Font Families
  static const String FF_DEFAULT = 'Roboto';
  static const List<String> FONT_FAMILIES = [
    'Roboto',
    'Montserrat',
    'IndieFlower',
  ];

  static const String BG_IMAGE_FOR_ABOUT_PAGE = 'assets/images/white-tree-portrait.jpg';
  static const String BG_IMAGE_FOR_SETTING_PAGE = 'assets/images/riga-cloudy-sky-landscape.jpg';
  static const String BG_IMAGE_FOR_CURRENCY_CONVERSION_PAGE = 'assets/images/portugal-sea.jpg';
  static const String BG_IMAGE_FOR_MAIN_MENU_AVATAR = 'assets/images/portugal-sea.jpg';
}