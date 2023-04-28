void main() {
  List<int> numbers = [1, 2, 3];

// Этот код не будет работать, потому что метод add() возвращает null
//   numbers.add(4).forEach((number) {
//     print(number);
//   });


// Этот код будет работать, потому что мы используем каскадную операцию
  numbers..add(4)..forEach((number) {
    print(number);
  });

  // Как если бы мы написали так:
  numbers.add(4); // add возвращает null
  numbers.forEach((number) {
    print(number);
  });
}
