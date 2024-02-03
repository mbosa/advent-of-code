mod part1;
mod part2;

use anyhow::Result;
use part1::part1;
use part2::part2;

fn main() {
    let input = include_str!("../../inputs/day9.txt");

    let part1 = part1(input).unwrap();
    println!("part1: {}", part1);

    let part2 = part2(input).unwrap();
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Result<Vec<Vec<i32>>> {
    let res = input
        .lines()
        .map(|line| {
            line.split_whitespace()
                .map(|s| s.parse::<i32>())
                .collect::<Result<Vec<_>, _>>()
        })
        .collect::<Result<Vec<_>, _>>()?;

    Ok(res)
}
