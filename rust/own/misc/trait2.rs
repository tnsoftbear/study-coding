#[derive(Debug)]
struct Foo<T>(T);

fn main() {
    let foo_str = Foo("Hello, world!");
    println!("foo_str: {:?}", foo_str); // вывод: foo_str: Foo("Hello, world!")
    println!("foo_i32: {:?}", Foo(42)); // вывод: foo_i32: Foo(42)
}
