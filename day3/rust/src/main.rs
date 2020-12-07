use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let terrain = get_terrain();

    // part 1
    let encounter1 = get_encounters(&terrain, Slope { down: 1, right: 1 });
    println!("Part 1: Tree Squares: {}", encounter1.trees);

    // part 2
    let encounter2 = get_encounters(&terrain, Slope { down: 1, right: 1 });
    let encounter3 = get_encounters(&terrain, Slope { down: 1, right: 5 });
    let encounter4 = get_encounters(&terrain, Slope { down: 1, right: 7 });
    let encounter5 = get_encounters(&terrain, Slope { down: 2, right: 1 });
    let result = encounter1.trees * encounter2.trees * encounter3.trees * encounter4.trees * encounter5.trees;
    println!("Part 2: Result: {}", result)
}

fn get_terrain() -> Vec<Vec<String>> {
    let file = File::open("../terrain.txt").expect("could not find file");
    let reader = BufReader::new(file);
    reader.lines()
        .map(|line| line.expect("could not read line"))
        .map(|line| line.split("").map(String::from).collect())
        .collect()
}

struct Slope {
    down: i32,
    right: i32,
}

struct Encounter {
    trees: i128,
}

const CHARACTER_TREE: &'static str = "#";

fn get_encounters(terrain: &Vec<Vec<String>>, slope: Slope) -> Encounter {
    let mut row = 0;
    let mut column = 0;
    let mut tree_squares = 0;
    while row + slope.down < (terrain.len() as i32) {
        row += slope.down;
        let next_terrain = &terrain[row as usize];
        column = (column + slope.right) % (next_terrain.len() as i32);
        if next_terrain[column as usize] == CHARACTER_TREE {
            tree_squares += 1;
        }
    }
    Encounter {
        trees: tree_squares,
    }
}