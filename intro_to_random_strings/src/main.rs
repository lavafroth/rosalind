use std::fs::read_to_string;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let contents = read_to_string("sample.txt")?;
    let mut lines = contents.lines();

    let strand = lines.next().expect("no DNA strand supplied").trim();
    let gc_contents: Vec<f64> = lines
        .next()
        .expect("no floating point probabilities")
        .trim()
        .split_whitespace()
        .map(|v| v.parse())
        .collect::<Result<_, _>>()?;

    let mut probs = vec![];
    // println!("{strand:?}");
    for gc in gc_contents {
        let claim_c = gc / 2.0;
        // print!("claim C/G = {claim_c:0.03} ");
        let claim_a = (1.0 - gc) / 2.0;
        // print!("A/T = {:0.03} ", claim_a);
        let prob = strand.chars().fold(1.0, |acc, ch| {
            acc * match ch {
                'A' | 'T' => claim_a,
                'C' | 'G' => claim_c,
                _ => 1.0,
            }
        });
        probs.push(prob.log10());
        // println!("prob = {:0.03} ", prob.log10());
    }

    let output = probs
        .into_iter()
        .map(|v| format!("{v:0.03}"))
        .collect::<Vec<_>>()
        .join(" ");
    println!("{output}");
    Ok(())
}
