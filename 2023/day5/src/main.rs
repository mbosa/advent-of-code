use std::{env, fs};

mod data;
mod part1;
mod part2;

use anyhow::{Error, Result};
use data::{Input, MapItem};
use part1::part1;
use part2::part2;

fn main() {
    let input_path = env::current_dir().unwrap().join("inputs/day5.txt");

    let input = fs::read_to_string(input_path).unwrap();

    let parsed = parse_input(&input).unwrap();

    let part1 = part1(&parsed);
    let part2 = part2(&parsed);

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Result<Input> {
    let mut i_iter = input.split("\n\n");
    let (_, seeds) = i_iter
        .next()
        .ok_or(Error::msg("Error parsing seeds"))?
        .split_once(": ")
        .ok_or(Error::msg("Error parsing seeds"))?;
    let seeds = seeds
        .split_whitespace()
        .map(|s| s.parse())
        .collect::<Result<Vec<i64>, _>>()?;

    let maps = i_iter
        .map(|map| {
            map.lines()
                .skip(1)
                .map(|line| line.parse::<MapItem>())
                .collect::<Result<Vec<_>>>()
        })
        .collect::<Result<Vec<_>>>()?;

    let i = Input { seeds, maps };

    Ok(i)
}
