fn main() {
    let input = "3.1\n\
    654\n\
    278\n";
    for (row, line) in input.split("\n").enumerate() {
        // println!("{} {}", line, row);
        for (col, ch) in line.chars().enumerate() {
            println!("{}_", ch);
        }
    }
}