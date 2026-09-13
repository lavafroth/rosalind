use std::fs::{OpenOptions, read_to_string};
use std::io::Write;

use anyhow::Result;
use reqwest::blocking::get;

fn find_n_glycosylation(s: &str) -> Vec<usize> {
    let window_size = 4;
    let iters = s.len() - window_size + 1;
    let s_bytes = s.as_bytes();
    let mut indices = vec![];
    for start in 0..iters {
        if let [b'N', not_p, b'S' | b'T', not_p_eiter] = s_bytes[start..start + window_size]
            && not_p != b'P'
            && not_p_eiter != b'P'
        {
            indices.push(start + 1);
        }
    }
    indices
}

fn parse_fasta_strand(fasta: &str) -> String {
    let mut strand = String::new();
    for line in fasta.lines() {
        if line.starts_with('>') {
            continue;
        }
        strand.push_str(line);
    }
    strand
}

fn main() -> Result<()> {
    let mut log = OpenOptions::new()
        .create(true)
        .write(true)
        .truncate(true)
        .open("run.log")?;
    for line in read_to_string("sample.txt")?.lines() {
        let id = line.split('_').next().unwrap_or(line);
        let response = get(format!("https://rest.uniprot.org/uniprotkb/{}.fasta", id))?;
        let text = response.text()?;
        let strand = parse_fasta_strand(&text);
        writeln!(log, "{strand}")?;
        let indices = find_n_glycosylation(&strand)
            .iter()
            .map(|v| v.to_string())
            .collect::<Vec<_>>()
            .join(" ");
        if !indices.is_empty() {
            println!("{line}");
            println!("{indices}");
        }
    }
    Ok(())
}
