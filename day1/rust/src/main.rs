fn main() {
    let file_name = "../expenses.csv".to_string();
    let map_function = |line: String| Some(line.parse().unwrap());
    let expenses = file_utils::get_input(file_name, map_function);

    // part 1
    'outer1: for (i, expense1) in expenses.iter().enumerate() {
        for (j, expense2) in expenses.iter().enumerate() {
            if i == j {
                continue;
            }
            if is_expected_sum_1(expense1, expense2) {
                println!("Part 1: Expense 1: {}, Expense 2: {}, Result {}", expense1, expense2,
                         expense1 * expense2);
                break 'outer1;
            }
        }
    }

    // part 2
    'outer2: for (i, expense1) in expenses.iter().enumerate() {
        for (j, expense2) in expenses.iter().enumerate() {
            for (k, expense3) in expenses.iter().enumerate() {
                if i == j || i == k || j == k {
                    continue;
                }
                if is_expected_sum_2(expense1, expense2, expense3) {
                    println!("Part 2: Expense 1: {}, Expense 2: {}, Expense 3: {}, Result {}",
                             expense1, expense2, expense3, expense1 * expense2 * expense3);
                    break 'outer2;
                }
            }
        }
    }
}

const EXPECTED: i32 = 2020;

fn is_expected_sum_1(x: &i32, y: &i32) -> bool {
    x + y == EXPECTED
}

fn is_expected_sum_2(x: &i32, y: &i32, z: &i32) -> bool {
    x + y + z == EXPECTED
}