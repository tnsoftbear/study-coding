import 'dart:mirrors';

class MyClass {
  String property1 = 'value1';
  String get myProperty => 'Hello, world!';
}

void main() {
  print(reflect(MyClass()).getField(Symbol('property1')).reflectee);
  print(reflect(MyClass()).getField(Symbol('myProperty')).reflectee);
}
