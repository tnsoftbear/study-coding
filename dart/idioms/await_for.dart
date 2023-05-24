Future<void> main() async {
  await for (final number in numbers.map((n) => n * n)) {
    print(number);
  }
}

Stream<int> get numbers async* {
  for (int i = 1; i <= 5; i++) {
    await Future.delayed(Duration(milliseconds: 500));
    yield i;
  }
}