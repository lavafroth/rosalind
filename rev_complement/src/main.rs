use std::fs::read_to_string;



fn main() {
    let contents =
    read_to_string
        ("./sample.txt")
    .expect("failed to read");

    let revcomp: String =
        contents.trim()
        .chars()
        .rev()
        .map(|character| {
            match character {
                'A' => 'T',
                'T' => 'A',
                'C' => 'G',
                'G' => 'C',
                _ => character
            }
        })
    .collect();
    println!("{revcomp}");
}
