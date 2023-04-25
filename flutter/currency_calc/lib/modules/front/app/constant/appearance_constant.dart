class AppearanceConstant {
  // Language Codes
  static const String LC_DEFAULT = LC_EN;
  static const String LC_EN = 'en';
  static const String LC_RU = 'ru';

  static const Map<String, Map<String, String>> CONFIG = {
    LC_EN: {'countryCode': 'US'},
    LC_RU: {'countryCode': 'RU'},
  };

  static const String FF_DEFAULT = 'Roboto';

  static const List<String> FONT_FAMILIES = [
    'Roboto',
    'Montserrat',
    'IndieFlower',
  ];
}