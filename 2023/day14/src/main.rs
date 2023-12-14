mod part1;
mod part2;

use part1::part1;
use part2::part2;

fn main() {
    let input = include_str!("../../inputs/day14.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Vec<Vec<u8>> {
    input.lines().map(|line| line.as_bytes().to_vec()).collect()
}

fn tilt_north(input: &mut Vec<Vec<u8>>) {
    for i in 1..input.len() {
        for j in 0..input[0].len() {
            if input[i][j] != b'O' {
                continue;
            }

            for k in (0..i).rev() {
                if input[k][j] == b'.' {
                    input[k][j] = b'O';
                    input[k + 1][j] = b'.';
                } else {
                    break;
                }
            }
        }
    }
}
fn tilt_west(input: &mut Vec<Vec<u8>>) {
    for j in 1..input[0].len() {
        for i in 0..input.len() {
            if input[i][j] != b'O' {
                continue;
            }

            for k in (0..j).rev() {
                if input[i][k] == b'.' {
                    input[i][k] = b'O';
                    input[i][k + 1] = b'.';
                } else {
                    break;
                }
            }
        }
    }
}
fn tilt_south(input: &mut Vec<Vec<u8>>) {
    for i in (0..input.len() - 1).rev() {
        for j in 0..input[0].len() {
            if input[i][j] != b'O' {
                continue;
            }

            for k in i + 1..input.len() {
                if input[k][j] == b'.' {
                    input[k][j] = b'O';
                    input[k - 1][j] = b'.';
                } else {
                    break;
                }
            }
        }
    }
}
fn tilt_east(input: &mut Vec<Vec<u8>>) {
    for j in (0..input[0].len() - 1).rev() {
        for i in 0..input.len() {
            if input[i][j] != b'O' {
                continue;
            }

            for k in j + 1..input[0].len() {
                if input[i][k] == b'.' {
                    input[i][k] = b'O';
                    input[i][k - 1] = b'.';
                } else {
                    break;
                }
            }
        }
    }
}
fn calc_north_load(input: &Vec<Vec<u8>>) -> usize {
    let mut res = 0;
    for i in 0..input.len() {
        for j in 0..input[0].len() {
            if input[i][j] == b'O' {
                res += input.len() - i;
            }
        }
    }

    res
}
