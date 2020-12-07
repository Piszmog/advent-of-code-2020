#[macro_use]
extern crate lazy_static;

use regex::Regex;

const FIELD_BIRTH_YEAR: &'static str = "byr";
const FIELD_ISSUE_YEAR: &'static str = "iyr";
const FIELD_EXPIRATION_YEAR: &'static str = "eyr";
const FIELD_HEIGHT: &'static str = "hgt";
const FIELD_HAIR_COLOR: &'static str = "hcl";
const FIELD_EYE_COLOR: &'static str = "ecl";
const FIELD_PASSPORT_ID: &'static str = "pid";
const FIELD_COUNTRY_ID: &'static str = "cid";

#[derive(Clone)]
struct Passport {
    birth_year: String,
    issue_year: String,
    expiration_year: String,
    height: String,
    hair_color: String,
    eye_color: String,
    id: String,
    country_id: String,
}

fn main() {
    let file_name = "../passports.txt".to_string();
    let mut passport = Passport {
        birth_year: "".to_string(),
        issue_year: "".to_string(),
        expiration_year: "".to_string(),
        height: "".to_string(),
        hair_color: "".to_string(),
        eye_color: "".to_string(),
        id: "".to_string(),
        country_id: "".to_string(),
    };
    let mut passports = file_utils::get_input(file_name, |line: String| {
        if line.is_empty() {
            let actual_passport = passport.clone();
            passport.birth_year = "".to_string();
            passport.issue_year = "".to_string();
            passport.expiration_year = "".to_string();
            passport.height = "".to_string();
            passport.hair_color = "".to_string();
            passport.eye_color = "".to_string();
            passport.id = "".to_string();
            passport.country_id = "".to_string();
            return Some(actual_passport);
        }
        let fields: Vec<String> = line.split(" ").map(String::from).collect();
        for f in fields {
            let field_parts: Vec<&str> = f.split(":").collect();
            match field_parts[0] {
                FIELD_BIRTH_YEAR => passport.birth_year = field_parts[1].to_string(),
                FIELD_ISSUE_YEAR => passport.issue_year = field_parts[1].to_string(),
                FIELD_EXPIRATION_YEAR => passport.expiration_year = field_parts[1].to_string(),
                FIELD_HEIGHT => passport.height = field_parts[1].to_string(),
                FIELD_HAIR_COLOR => passport.hair_color = field_parts[1].to_string(),
                FIELD_EYE_COLOR => passport.eye_color = field_parts[1].to_string(),
                FIELD_PASSPORT_ID => passport.id = field_parts[1].to_string(),
                FIELD_COUNTRY_ID => passport.country_id = field_parts[1].to_string(),
                _ => {}
            }
        }
        return None;
    });
    passports.push(passport);

    // part 1
    let valid_count1 = get_valid_passports_count1(&passports);
    println!("Part 1: Valid Passports {}", valid_count1);

    // part 2
    let valid_count2 = get_valid_passports_count2(&passports);
    println!("Part 2: Valid Passports {}", valid_count2);
}

fn get_valid_passports_count1(passports: &Vec<Passport>) -> i32 {
    let mut count = 0;
    for p in passports {
        if has_required_fields(p) {
            count += 1;
        }
    }
    count
}

fn has_required_fields(p: &Passport) -> bool {
    p.birth_year.len() != 0 && p.issue_year.len() != 0 && p.expiration_year.len() != 0
        && p.height.len() != 0 && p.hair_color.len() != 0 && p.eye_color.len() != 0 && p.id.len() != 0
}

fn get_valid_passports_count2(passports: &Vec<Passport>) -> i32 {
    let mut count = 0;
    for p in passports {
        if has_valid_fields(p) {
            count += 1;
        }
    }
    count
}

fn has_valid_fields(p: &Passport) -> bool {
    lazy_static! {
        static ref BIRTH_RE: Regex = Regex::new(r"^[0-9]{4}$").unwrap();
        static ref ISSUE_RE: Regex = Regex::new(r"^[0-9]{4}$").unwrap();
        static ref EXPIRATION_RE: Regex = Regex::new(r"^[0-9]{4}$").unwrap();
        static ref HEIGHT_RE: Regex = Regex::new(r"^([0-9]{2,3})((cm)|(in))$").unwrap();
        static ref HAIR_COLOR_RE: Regex = Regex::new(r"#[0-9a-f]{6}").unwrap();
        static ref EYE_COLOR_RE: Regex = Regex::new(r"^(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)$").unwrap();
        static ref ID_RE: Regex = Regex::new(r"^[0-9]{9}$").unwrap();
    }
    if !BIRTH_RE.is_match(p.birth_year.as_str()) || !ISSUE_RE.is_match(p.issue_year.as_str()) ||
        !EXPIRATION_RE.is_match(p.expiration_year.as_str()) || !HEIGHT_RE.is_match(p.height.as_str()) ||
        !HAIR_COLOR_RE.is_match(p.hair_color.as_str()) || !EYE_COLOR_RE.is_match(p.eye_color.as_str()) ||
        !ID_RE.is_match(p.id.as_str()) {
        return false;
    }
    let birth_year = p.birth_year.parse::<i32>().unwrap();
    let issue_year = p.issue_year.parse::<i32>().unwrap();
    let expiration_year = p.expiration_year.parse::<i32>().unwrap();
    if birth_year < 1920 || birth_year > 2002 {
        return false;
    }
    if issue_year < 2010 || issue_year > 2020 {
        return false;
    }
    if expiration_year < 2020 || expiration_year > 2030 {
        return false;
    }
    let height_cap = HEIGHT_RE.captures(p.height.as_str()).unwrap();
    let height = height_cap.get(1).unwrap().as_str().parse::<i32>().unwrap();
    let metric_type = height_cap.get(2).unwrap().as_str();
    match metric_type {
        "in" => height >= 59 && height <= 76,
        "cm" => height >= 150 && height <= 193,
        _ => false
    }
}