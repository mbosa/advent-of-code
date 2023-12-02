use data::{ParseGameError, Rgb};
use std::{env, fs};

mod data;
mod part1;
mod part2;

use part1::part1;
use part2::part2;

use data::Game;

fn main() {
    let input_path = env::current_dir().unwrap().join("inputs/day2.txt");

    let input = fs::read_to_string(input_path).unwrap();

    let parsed = parse_input(&input).unwrap();

    let part1 = part1(&parsed);
    let part2 = part2(&parsed);

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Result<Vec<Game>, ParseGameError> {
    input.lines().map(|line| line.parse::<Game>()).collect()
}
