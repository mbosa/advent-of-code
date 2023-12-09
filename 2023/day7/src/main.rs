mod data;
mod part1;
mod part2;

use part1::part1;
use part2::part2;

fn main() {
    let input = include_str!("../../inputs/day7.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Vec<(&str, u32)> {
    input
        .lines()
        .map(|line| line.split_once(" ").unwrap())
        .map(|(a, b)| (a, b.parse().unwrap()))
        .collect()
}
