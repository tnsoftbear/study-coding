# Currency Calculator

My 1st experiment with flutter and dart in April of 2023.

Module logic is separated to application, infrastructure and domain layers. 
I think, the "package by feature" style could look like this :) 

**Infrastructure layer** is responsible for the currency exchange rate loading by API. 

Rate fetchers implement common interface and are provided by factory.
**Application layer** operates by this interface and provides the currency exchange rate to the UI.  
It translates with help of localization package, and format currency amounts and numbers with help of Intl package.

**Domain layer** validates input values and calculates the currency conversion.

We have unit tests for validation and few functional tests for checking application view.

## Links

* [Folder structure for Flutter with clean architecture. How I do.](https://felipeemidio.medium.com/folder-structure-for-flutter-with-clean-architecture-how-i-do-bbe29225774f)
* [Style guide for Flutter repo](https://github.com/flutter/flutter/wiki/Style-guide-for-Flutter-repo)
* [Internationalizing Flutter apps](https://docs.flutter.dev/development/accessibility-and-localization/internationalization)
* [intl](https://pub.dev/packages/intl)