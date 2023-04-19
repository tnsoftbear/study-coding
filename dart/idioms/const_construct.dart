class Circle {
  final double radius;
  static const double pi = 3.14159;

  const Circle(this.radius);

  double get circumference => 2 * pi * radius;
  double get area => pi * radius * radius;
}

void main() {
  const circle = Circle(5);
  print('Radius: ${circle.radius}');
  print('Circumference: ${circle.circumference}');
  print('Area: ${circle.area}');
}
