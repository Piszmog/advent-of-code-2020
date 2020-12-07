use std::fs::File;
use std::io::{BufRead, BufReader};

pub fn get_input<T, F>(file_name: String, map_function: F) -> Vec<T>
    where
        F: FnMut(String) -> Option<T>
{
    let file = File::open(file_name).expect("could not find file");
    let reader = BufReader::new(file);
    reader.lines()
        .map(|line| line.expect("could not read line"))
        .map(map_function)
        .filter(|o| o.is_some())
        .map(|o| o.unwrap())
        .collect()
}