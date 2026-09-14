fn main() {
    let file_contents = std::fs::read_to_string("sample.txt").expect("failed to read file");
    let mut max_id = "";
    let mut max = 0.0;
    for chunk in file_contents.split('>') {
        if chunk.is_empty() {
            continue;
        }
        let Some((id, rest)) = chunk.split_once('\n') else {
            break;
        };

        let mut len = rest.as_bytes().len();
        let mut gc = 0;
        for b in rest.bytes() {
            match b {
                b'G' | b'C' => gc += 1,
                b'\n' => len -= 1,
                _ => {}
            }
        }

        let gc_content = gc as f32 / len as f32;
        if gc_content > max {
            max = gc_content;
            max_id = id;
        }
    }
    println!("{max_id}");
    println!("{max}");
}
