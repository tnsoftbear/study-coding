import 'dart:collection';

void main() {
    var hashSet = HashSet<String>();
    hashSet.add('apple');
    hashSet.add('banana');
    hashSet.add('orange');
    print(hashSet); // выводит {apple, banana, orange}
    hashSet.remove('banana');
    print(hashSet); // выводит {apple, orange}

    var linkedSet = LinkedHashSet<String>();
    linkedSet.add('apple');
    linkedSet.add('banana');
    linkedSet.add('orange');
    print(linkedSet); // выводит {apple, banana, orange}
    linkedSet.remove('banana');
    print(linkedSet); // выводит {apple, orange}

    var splayTreeSet = SplayTreeSet<String>();
    splayTreeSet.add('apple');
    splayTreeSet.add('banana');
    splayTreeSet.add('orange');
    print(splayTreeSet); // выводит {apple, banana, orange}
    splayTreeSet.remove('banana');
    print(splayTreeSet); // выводит {apple, orange}
}