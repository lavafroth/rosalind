use std::fs::read_to_string;

fn main() {
    let contents =
        read_to_string
        ("./sample.txt")
    .expect("failed to read");

    let mut split =
    contents
        .split_whitespace();

    let n: u64 = split.next()
        .expect("bad input")
        .parse()
        .expect(
            "failed to parse"
        );

    let k: u64 = split.next()
        .expect("bad input")
        .parse()
        .expect(
            "failed to parse"
        );

    let mut f1 = 1;
    let mut f2 = 1;

    if n < 3 {
        println!("1");
        return
    }

    for _ in 2..n {
        let offsprings
         = f1 * k + f2;

         f1 = f2;
         f2 = offsprings;
    }
    println!("{f2}");
}

// 1 -> 1
// 1 -> 1 + 1 = 2
// 2 -> 2 + 1 = 3
// 3 -> 3 + 2 = 5

// f1, f2 = 1, 1
// 1 -> 1 + f1 * 3 = 4
// 4 -> 4 + 1 * 3 = 7
// 7 -> 7 + 4 * 3 = 19
