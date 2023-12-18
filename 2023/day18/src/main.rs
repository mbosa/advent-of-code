mod part1;
mod part2;

use part1::part1;
use part2::part2;

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
enum Direction {
    Up,
    Down,
    Right,
    Left,
}

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct Position {
    row: i64,
    col: i64,
}

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct Dig<'a> {
    direction: Direction,
    steps: i64,
    color: &'a str,
}

fn main() {
    let input = include_str!("../../inputs/day18.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Vec<Dig> {
    input
        .lines()
        .map(|line| {
            let mut s = line.split_whitespace();

            let direction = match s.next().unwrap() {
                "U" => Direction::Up,
                "D" => Direction::Down,
                "R" => Direction::Right,
                "L" => Direction::Left,
                _ => unreachable!(),
            };

            let steps = s.next().unwrap().parse::<i64>().unwrap();
            let color = s.next().map(|s| &s[1..s.len() - 1]).unwrap();

            Dig {
                direction,
                steps,
                color,
            }
        })
        .collect()
}
