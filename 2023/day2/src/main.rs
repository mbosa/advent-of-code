use data::{ParseGameError, Rgb};

mod data;
mod part1;
mod part2;

use part1::part1;
use part2::part2;

use data::Game;

fn main() {
    let input = include_str!("../../inputs/day2.txt");

    let parsed = parse_input(&input).unwrap();

    let part1 = part1(&parsed);
    let part2 = part2(&parsed);

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Result<Vec<Game>, ParseGameError> {
    input.lines().map(|line| line.parse::<Game>()).collect()
}
