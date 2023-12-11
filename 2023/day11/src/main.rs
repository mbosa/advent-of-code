mod part1;
mod part2;

use part1::part1;
use part2::part2;

type Input = Vec<Vec<char>>;
type Position = [usize; 2];

fn main() {
    let input = include_str!("../../inputs/day11.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Input {
    input.lines().map(|line| line.chars().collect()).collect()
}

fn find_galaxies_with_expansion(input: &Input, expand_by_times: u32) -> Vec<Position> {
    let mut galaxies: Vec<Position> = Vec::new();

    // x is the expanded row coordinate
    let mut x = 0;
    for row in input.iter() {
        let mut empty_line = true;
        for (j, &el) in row.iter().enumerate() {
            if el == '#' {
                empty_line = false;
                galaxies.push([x, j]);
            }
        }
        x += match empty_line {
            true => expand_by_times as usize,
            false => 1,
        };
    }

    //columns
    for j in (0..input[0].len()).rev() {
        if input.iter().all(|row| row[j] == '.') {
            for k in 0..galaxies.len() {
                if galaxies[k][1] > j {
                    galaxies[k][1] += expand_by_times as usize - 1;
                }
            }
        }
    }

    galaxies
}

fn manhattan_distance(a: &Position, b: &Position) -> u64 {
    (usize::abs_diff(a[0], b[0]) + usize::abs_diff(a[1], b[1])) as u64
}

fn solve(input: &Input, expand_by_times: u32) -> u64 {
    let galaxies = find_galaxies_with_expansion(&input, expand_by_times);

    let mut res = 0;

    for i in 0..galaxies.len() {
        for j in i + 1..galaxies.len() {
            res += manhattan_distance(&galaxies[i], &galaxies[j])
        }
    }

    res
}
