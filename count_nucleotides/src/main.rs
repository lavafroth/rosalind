fn main() {
    let file_contents = std::fs::read_to_string("sample.txt").expect("failed to read file");

    let mut count_a = 0;
    let mut count_c = 0;
    let mut count_g = 0;
    let mut count_t = 0;

    for c in file_contents.chars() {
        if c == 'A' {
            count_a += 1;
        }
        if c == 'C' {
            count_c += 1;
        }
        if c == 'G' {
            count_g += 1;
        }
        if c == 'T' {
            count_t += 1;
        }
    }

    println!("{} {} {} {}", count_a, count_c, count_g, count_t);
}
