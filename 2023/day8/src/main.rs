use std::collections::HashMap;

mod part1;
mod part2;

use anyhow::{Error, Result};
use part1::part1;
use part2::part2;

#[derive(Debug)]
struct Input<'a> {
    instructions: &'a str,
    nodes: HashMap<&'a str, [&'a str; 2]>,
}

fn main() {
    let input = include_str!("../../inputs/day8.txt");

    let part1 = part1(input).unwrap();
    println!("part1: {}", part1);

    let part2 = part2(input).unwrap();
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Result<Input> {
    let (instructions, node_list) = input
        .split_once("\n\n")
        .ok_or(Error::msg("Error parsing the input"))?;

    let nodes = node_list
        .lines()
        .map(|line| {
            let (node, connections) = line
                .split_once(" = ")
                .ok_or(Error::msg("Error parsing the input"))?;

            let (left, right) = connections
                .split_once(", ")
                .ok_or(Error::msg("Error parsing the input"))?;

            let left = &left[1..];
            let right = &right[..3];

            Ok::<_, Error>((node, [left, right]))
        })
        .collect::<Result<HashMap<_, _>>>()?;

    let i = Input {
        instructions,
        nodes,
    };

    Ok(i)
}
