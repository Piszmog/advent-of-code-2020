use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let passwords = get_passwords();

    // part 1
    let valid_password_count = get_valid_passwords_count_1(&passwords);
    println!("Part 1: There are {} valid passwords", valid_password_count);

    // part 2
    let valid_password_count = get_valid_passwords_count_2(&passwords);
    println!("Part 2: There are {} valid passwords", valid_password_count);
}

struct Password {
    min: i32,
    max: i32,
    letter: String,
    value: String,
}

fn get_passwords() -> Vec<Password> {
    let file = File::open("../passwords.txt").expect("could not find file");
    let reader = BufReader::new(file);
    reader.lines()
        .map(|line| line.expect("could not read line"))
        .map(|line| {
            let parts: Vec<&str> = line.split(" ").collect();
            let min_max: Vec<&str> = parts[0].split("-").collect();
            let min = min_max[0].parse::<i32>().unwrap();
            let max = min_max[1].parse::<i32>().unwrap();
            let letter = parts[1].trim_end_matches(":").to_string();
            Password {
                min,
                max,
                letter,
                value: parts[2].to_string(),
            }
        })
        .collect()
}

fn get_valid_passwords_count_1(passwords: &Vec<Password>) -> i32 {
    let mut valid_password_count = 0;
    passwords.iter().for_each(|p| {
        let count = p.value.matches(&p.letter).count();
        if count >= p.min as usize && count <= p.max as usize {
            valid_password_count += 1;
        }
    });
    valid_password_count
}

fn get_valid_passwords_count_2(passwords: &Vec<Password>) -> i32 {
    let mut valid_password_count = 0;
    passwords.iter().for_each(|p| {
        let r = p.letter.chars().next().unwrap();

        let has_pos_1 = p.value.len() >= p.min as usize
            && p.value.chars().nth((p.min - 1) as usize).unwrap() == r;

        let has_pos_2 = p.value.len() >= p.max as usize
            && p.value.chars().nth((p.max - 1) as usize).unwrap() == r;

        if (has_pos_1 && !has_pos_2) || (!has_pos_1 && has_pos_2) {
            valid_password_count += 1;
        }
    });
    valid_password_count
}
