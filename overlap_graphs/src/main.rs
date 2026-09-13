use std::fs::read_to_string;

#[derive(Debug)]
pub struct Summary<'a> {
    prefix: &'a str,
    suffix: &'a str,
    name: &'a str,
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let text = read_to_string("./sample.txt")?;

    let mut strands = vec![];

    for line in text.lines() {
        if line.starts_with('>') {
            strands.push(Summary {
                prefix: "",
                suffix: "",
                name: line.strip_prefix('>').unwrap_or(line),
            });

            continue;
        }

        let Some(last_strand) = strands.last_mut() else {
            continue;
        };

        if last_strand.prefix.is_empty() {
            last_strand.prefix = &line[..3];
        }

        last_strand.suffix = &line[line.len() - 3..];
    }

    for (i, left) in strands.iter().enumerate() {
        for (j, right) in strands.iter().enumerate() {
            if left.suffix == right.prefix && i != j {
                println!("{} {}", left.name, right.name);
            }
        }
    }

    Ok(())
}
