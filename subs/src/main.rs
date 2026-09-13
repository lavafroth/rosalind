use std::fs::read_to_string;

fn eat(_: String) {}

fn uses_ref(s: &str) {}
fn mutates_ref(s: &mut str) {}

fn main() {

let mut owned = "hello"
    .to_string();

eat(owned);

// let refer = owned.as_str();


// let another_tenant
//     = owned.as_str();

// mutates_ref(&mut owned);
// uses_ref(refer);

























    
    let content =
    read_to_string
    ("./sample.txt")
    .expect("woooo");

    let mut lines =
        content.lines();

    let haystack
        = lines.next()
        .expect("bruh");

    let needle
        = lines.next()
        .expect("bruh");

    let n_windows
        = haystack.len()
        - needle.len()
        + 1;

    // let owned = String::new();

    let mut output = vec![];

    for shift
        in 0..n_windows {
        let start = shift;
        let end = start +
            needle.len();
        if &haystack
            [start..end]
        == needle {
        output
            .push(
            (shift + 1)
            .to_string()
            )
        }
    }

    println!("{}",
        output.join(" "));
}
