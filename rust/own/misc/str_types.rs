fn main() {
    // Пример использования &str
    let my_str: &str = "Привет, &str"; // Строка в кодировке UTF-8
    println!(
        "Length of string: {}, First character: {}",
        my_str.len(),                                 // Получаем длину строки в байтах
        my_str.chars().next().unwrap()                // Получаем первый символ строки
    );
    println!("my_str: {:?}", my_str);
    println!("my_str.as_bytes(): {:?}, len: {}", my_str.as_bytes(), my_str.as_bytes().len());
    println!("my_str.chars(): {:?}, len:", my_str.chars());       // Итератор по байтам

    // Пример использования &[char]
    // let char_slice: &[char] = &['H', 'e', 'l', 'l', 'o']; // Срез символов
    let char_slice: &mut [char] = &mut ['H', 'e', 'l', 'l', 'o']; // Срез символов
    char_slice[4] = '0';
    println!("{:?} - Length of char slice: {}", char_slice, char_slice.len()); // Получаем длину среза символов
    for ch in char_slice {
        println!("Character: {}", ch); // Печатаем каждый символ в срезе
    }

    // Массив фиксированного размера
    let char_arr: [char; 5] = ['H', 'e', 'l', 'l', 'o'];
    println!("{:?} - Length of char array: {}", char_arr, char_arr.len()); // Получаем длину массива

    let string1: String = String::from("Goodbye, String!");
    let string2: String = string1.clone();
    println!("string2: {:?}", string2);
}