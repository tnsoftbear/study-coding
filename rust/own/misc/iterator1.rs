fn main() {
    iter1();
}

fn iter1() {
    #[derive(Debug)]
    struct Int(usize);

    let vec = [2, 4, 6, 1, 3, 5, 8, 10, 1, 2]
        .into_iter()
        .map(Int)
        .collect::<Vec<_>>();

    println!("{:?}", vec);
}