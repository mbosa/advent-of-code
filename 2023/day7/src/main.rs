use std::{env, fs};

mod data;
mod part1;
mod part2;

use part1::part1;
use part2::part2;

fn main() {
    let input_path = env::current_dir().unwrap().join("inputs/day7.txt");

    let input = fs::read_to_string(input_path).unwrap();

    let parsed = parse_input(&input);

    let part1 = part1(&parsed);
    println!("part1: {}", part1);

    let part2 = part2(&parsed);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Vec<(&str, u32)> {
    input
        .lines()
        .map(|line| line.split_once(" ").unwrap())
        .map(|(a, b)| (a, b.parse().unwrap()))
        .collect()
}
