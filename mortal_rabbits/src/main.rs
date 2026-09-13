use std::{fmt::Debug, fs::read_to_string};

#[derive(Clone, Copy)]
pub struct Rabbits {
    mature: u64,
    babies: u64,
}

impl Rabbits {
    fn total(&self) -> u64 {
        self.mature + self.babies
    }
}

impl Debug for Rabbits {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "{} (mature={}, babies={})",
            self.total(),
            self.mature,
            self.babies
        )
    }
}

fn main() {
    let contents = read_to_string("./sample.txt").expect("failed to read");

    let mut split = contents.split_whitespace();

    let n: u64 = split
        .next()
        .expect("bad input")
        .parse()
        .expect("failed to parse");

    let k: usize = split
        .next()
        .expect("bad input")
        .parse()
        .expect("failed to parse");

    let mut f = vec![
        Rabbits {
            mature: 0,
            babies: 1,
        },
        Rabbits {
            mature: 1,
            babies: 0,
        },
    ];

    if n < 3 {
        println!("1");
        return;
    }

    // println!("{f:?}");

    for _ in 2..n {
        let last = |x| f.len() - x;
        let last_generation = f[last(1)];

        let mut mature = last_generation.babies + last_generation.mature;
        let babies = last_generation.mature;

        if k <= f.len() {
            let k_last_generation = f[last(k)];
            // println!("kth last generation: {k_last_generation:?}");
            mature -= k_last_generation.babies;
        }
        let this_generation = Rabbits { mature, babies };
        f.push(this_generation);
        // println!("{this_generation:?}");
    }

    println!(
        "{}",
        f.last().expect("the last generation should exist").total()
    );
}
