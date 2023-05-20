void main() {
  final now = DateTime.now();
  print(now.toString());
  print(now.toUtc().toString());
  final ts = now.millisecondsSinceEpoch;
  print("now.millisecondsSinceEpoch: $ts");
  final ts2 = now.toUtc().millisecondsSinceEpoch;
  print("now.toUtc().millisecondsSinceEpoch: $ts2");
  final now2 = DateTime.fromMillisecondsSinceEpoch(ts);
  print(now2.toString());
  print(now2.toUtc().toString());
}