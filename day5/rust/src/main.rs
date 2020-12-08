fn main() {
    let file_name = "../input.txt".to_string();
    let mut input = file_utils::get_input(file_name, |line| {
        Some(BoardingPass {
            value: line.split("").map(String::from).collect(),
            row: 0,
            column: 0,
            seat_id: 0,
        })
    });

    // part 1
    let mut highest_id = 0;
    for p in &mut input {
        let value: &Vec<String> = &p.value;
        let decoded_values = decode_value(value);
        let seat_id = decoded_values.0 * 8 + decoded_values.1;
        p.seat_id = seat_id;
        p.row = decoded_values.0;
        p.column = decoded_values.1;
        if highest_id < seat_id {
            highest_id = seat_id;
        }
    }
    println!("Part 1: Highest seat id: {}", highest_id);

    // part 2
    input.sort_by(|a, b| a.seat_id.cmp(&b.seat_id));
    let mut missing_id = 0;
    for (i, p) in input.iter().enumerate() {
        if i == input.len() - 1 {
            continue;
        }
        if p.seat_id != input[i + 1].seat_id - 1 {
            missing_id = p.seat_id + 1;
            break;
        }
    }
    println!("Part 2: Missing seat id: {}", missing_id);
}

#[derive(Debug, Eq, Ord, PartialEq, PartialOrd)]
struct BoardingPass {
    value: Vec<String>,
    row: i32,
    column: i32,
    seat_id: i32,
}

fn decode_value(v: &Vec<String>) -> (i32, i32) {
    let mut current_min_row = 0;
    let mut current_max_row = 127;
    let mut current_min_column = 0;
    let mut current_max_column = 7;
    v.iter().for_each(|e| {
        match e.as_str() {
            "F" => current_max_row = (current_max_row - current_min_row) / 2 + current_min_row,
            "B" => current_min_row = (current_max_row + current_min_row) / 2 + 1,
            "L" => current_max_column = (current_max_column - current_min_column) / 2 + current_min_column,
            "R" => current_min_column = (current_max_column + current_min_column) / 2 + 1,
            _ => (),
        }
    });
    (current_max_row, current_min_column)
}