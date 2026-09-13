use std::fs::read_to_string;
fn comb(n: u128, mut k: u128) -> u128 {
    if k > n {
        return 0;
    }
    if k * 2 > n {
        k = n - k;
    }

    let mut result: u128 = 1;
    for i in 1..=k {
        result = result * (n - k + i) / i;
    }
    result
}
fn main() {
    let contents = read_to_string("./sample.txt").expect("failed to read");

    let mut split = contents.split_whitespace();

    let generation_number: u64 = split
        .next()
        .expect("bad input")
        .parse()
        .expect("failed to parse");

    let want_successes: u64 = split
        .next()
        .expect("bad input")
        .parse()
        .expect("failed to parse");

    let population = 1 << generation_number;
    let p_success: f32 = 0.25;

    // cases where there are fewer successful trials than wanted
    let p_unsatisfied: f32 = (0..want_successes)
        .into_iter()
        .map(|i| {
            comb(population as u128, i as u128) as f32
                * p_success.powi(i.try_into().unwrap())
                * (1.0 - p_success).powi((population - i).try_into().unwrap())
        })
        .sum();

    println!("{}", 1.0 - p_unsatisfied);
}
