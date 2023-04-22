import 'my_helper.dart';

void main() {
  final myClass = MyClass();
  myClass.sayHello();

  final myHelper = MyHelper();
  myHelper.doSomething(myClass);

  myClass.sayHello();
}
