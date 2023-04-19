void main() {
    // Пример List
    var myList = [1, 2, 3];
    print(myList[1]); // выводит "2"
    myList.add(4); // добавляет элемент "4" в конец списка
    myList.removeAt(1); // удаляет элемент с индексом "1"
    print(myList); // выводит [1, 3, 4]
    
    // Пример Map
    var myMap = {1: 'apple', 2: 'banana', 3: 'orange'};
    print(myMap[2]); // выводит "banana"
    myMap[4] = 'grape'; // добавляет пару ключ-значение в Map
    myMap.remove(2); // удаляет пару ключ-значение с ключом "2" из Map
    print(myMap); // {1: apple, 3: orange, 4: grape}

    // Пример Set
    var mySet = {1, 2, 3};
    print(mySet.contains(2)); // выводит "true"
    mySet.add(4); // добавляет элемент "4" в множество
    mySet.remove(2); // удаляет элемент "2" из множества
    print(mySet); // {1, 3, 4}

    // Операции для множеств
    // Для Set
    var set1 = {1, 2, 3};
    var set2 = {2, 3, 4};
    var differenceForSet = set1.difference(set2); // возвращает {1}
    var intersectionForSet = set1.intersection(set2); // возвращает {2, 3}
    var unionForSet = set1.union(set2); // возвращает {1, 2, 3, 4}
    print(differenceForSet);
    print(intersectionForSet);
    print(unionForSet);

    // Для List
    var list1 = [1, 2, 3];
    var list2 = [2, 3, 4];
    var differenceForList = list1.toSet().difference(list2.toSet()).toList(); // возвращает [1]
    var intersectionForList = list1.toSet().intersection(list2.toSet()).toList(); // возвращает [2, 3]
    var unionForList = [...list1, ...list2].toSet().toList(); // возвращает [1, 2, 3, 4]
    print(differenceForList);
    print(intersectionForList);
    print(unionForList);

    // Для Map
    var map1 = {'a': 1, 'b': 2, 'c': 3};
    var map2 = {'b': 2, 'c': 3, 'd': 4};
    var differenceForMap = map1.keys.toSet().difference(map2.keys.toSet()).toList(); // возвращает ['a']
    var intersectionForMap = map1.keys.toSet().intersection(map2.keys.toSet()).toList(); // возвращает ['b', 'c']
    var unionForMap = {...map1, ...map2}; // возвращает {a: 1, b: 2, c: 3, d: 4}
    var unionKeysForMap = {...map1, ...map2}.keys.toList(); // возвращает ['a', 'b', 'c', 'd']
    print(differenceForMap);
    print(intersectionForMap);
    print(unionForMap);
    print(unionKeysForMap);
}