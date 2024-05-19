trait Hello {
    fn hello(&self);
}

// Не компилируется, потому что размер типа не известен.
// error[E0277]: the size for values of type `dyn Hello` cannot be known at compilation time
// fn func_not_1(arr: &[dyn Hello]) {
//     for i in arr {
//         i.hello();
//     }
// }

// Скомпилируется, потому что обобщённый параметр T неявно требует ограничение Sized
fn func1<T: Hello>(arr: &[T]) {
    for i in arr {
        i.hello();
    }
}

fn func2(arr: &[&dyn Hello]) {
    for i in arr {
        i.hello();
    }
}

struct Bar;
impl Hello for Bar {
    fn hello(&self) {
        println!("Hello, Bar!")
    }
}

struct Baz;
impl Hello for Baz {
    fn hello(&self) {
        println!("Hello, Baz!")
    }
}

fn main() {
    let bar1_1 = Bar;
    let bar1_2 = Bar;
    let baz1 = Baz;
    func1(&[bar1_1, bar1_2]);
    // func1(&[bar1_1, bar1_2, baz1]); // error[E0308]: mismatched types
    let bar2 = Bar;
    func2(&[&bar2, &baz1]);
}