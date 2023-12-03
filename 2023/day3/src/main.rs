use std::{env, fs};

mod part1;
mod part2;
mod utils;

use part1::part1;
use part2::part2;

fn main() {
    let input_path = env::current_dir().unwrap().join("inputs/day3.txt");

    let input = fs::read_to_string(input_path).unwrap();

    let parsed = parse_input(&input);

    let part1 = part1(&parsed);
    let part2 = part2(&parsed);

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Vec<Vec<char>> {
    input
        .lines()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>()
}
