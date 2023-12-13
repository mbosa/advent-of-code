mod part1;
mod part2;

use part1::part1;
use part2::part2;

type Input<'a> = Vec<Vec<&'a [u8]>>;

fn main() {
    let input = include_str!("../../inputs/day13.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Input {
    input
        .split("\n\n")
        .map(|pattern| pattern.lines().map(|line| line.as_bytes()).collect())
        .collect()
}

fn find_row(pattern: &Vec<&[u8]>, error_limit: usize) -> Option<usize> {
    for i in 1..pattern.len() {
        let mut errors = 0;

        'delta_loop: for di in 0..usize::min(i, pattern.len() - i) {
            let up = i - 1 - di;
            let down = i + di;

            for j in 0..pattern[0].len() {
                if pattern[up][j] != pattern[down][j] {
                    errors += 1;

                    if errors > error_limit {
                        break 'delta_loop;
                    }
                }
            }
        }
        if errors == error_limit {
            return Some(i);
        }
    }

    None
}

fn find_col(pattern: &Vec<&[u8]>, error_limit: usize) -> Option<usize> {
    for j in 1..pattern[0].len() {
        let mut errors = 0;

        'delta_loop: for dj in 0..usize::min(j, pattern[0].len() - j) {
            let left = j - 1 - dj;
            let right = j + dj;

            for i in 0..pattern.len() {
                if pattern[i][left] != pattern[i][right] {
                    errors += 1;

                    if errors > error_limit {
                        break 'delta_loop;
                    }
                }
            }
        }
        if errors == error_limit {
            return Some(j);
        }
    }

    None
}
