use std::fs::read_to_string;

fn nucleotide_to_index(ch: char) -> usize {
    match ch {
        'A' => 0,
        'C' => 1,
        'G' => 2,
        'T' => 3,
        _ => 4,
    }
}

const SEQ: [char; 4] = ['A', 'C', 'G', 'T'];

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let contents = read_to_string("sample.txt")?;
    let mut consensus = vec![];

    let mut strands = vec![];

    for line in contents.lines() {
        if line.starts_with('>') {
            let strand = String::new();
            strands.push(strand);
            continue;
        }

        if let Some(strand) = strands.last_mut() {
            strand.push_str(&line.trim());
        }
    }

    for strand in strands {
        while consensus.len() < strand.len() {
            consensus.push([0, 0, 0, 0]);
        }

        for (index, ch) in strand.char_indices() {
            let summary_index = nucleotide_to_index(ch);
            if summary_index > 3 {
                continue;
            }
            consensus[index][summary_index] += 1;
        }
    }

    for summary in consensus.iter() {
        let index = summary
            .iter()
            .enumerate()
            .max_by_key(|(_, c)| **c)
            .expect("there must be a max")
            .0;

        print!("{}", SEQ[index]);
    }
    println!();

    for nl in 0..4 {
        let mut consensus_str = vec![];
        for summary in consensus.iter() {
            let thing = summary[nl];
            consensus_str.push(thing.to_string());
        }
        println!("{}: {}", SEQ[nl], consensus_str.join(" "));
    }

    Ok(())
}
