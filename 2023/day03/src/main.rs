mod part1;
mod part2;
mod utils;

use part1::part1;
use part2::part2;

fn main() {
    let input = include_str!("../../inputs/day3.txt");

    let part1 = part1(input);
    let part2 = part2(input);

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Vec<Vec<char>> {
    input
        .lines()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>()
}
