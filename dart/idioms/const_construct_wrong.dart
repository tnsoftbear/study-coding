class Point<T> {
  final T x;
  final T y;
  const Point(T x, T y)
    : this.x = x,
      this.y = y;
}

void main() {
  const point = Point(1, 1);
  // const point2 = Point(DateTime.now(), DateTime.now());
}