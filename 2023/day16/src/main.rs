mod part1;
mod part2;

use part1::part1;
use part2::part2;

type Input = Vec<Vec<u8>>;

#[derive(Debug, Clone, Copy, Eq, PartialEq, Hash)]
enum Direction {
    Down,
    Right,
    Up,
    Left,
}

#[derive(Debug, Clone, Copy, Eq, PartialEq, Hash)]
struct Position {
    row: usize,
    col: usize,
}

#[derive(Debug, Clone, Copy, Eq, PartialEq, Hash)]
struct State {
    position: Position,
    direction: Direction,
}

fn main() {
    let input = include_str!("../../inputs/day16.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Input {
    input.lines().map(|line| line.bytes().collect()).collect()
}

fn run_beam(input: &Input, start: State) -> usize {
    let mut stack = Vec::new();
    stack.push(start);

    let mut seen: Vec<Vec<[bool; 4]>> = vec![vec![[false; 4]; input[0].len()]; input.len()];

    while let Some(curr) = stack.pop() {
        match seen[curr.position.row][curr.position.col][curr.direction as usize] {
            true => continue,
            false => seen[curr.position.row][curr.position.col][curr.direction as usize] = true,
        }

        match input[curr.position.row][curr.position.col] {
            b'.' => {
                if let Some(next_pos) = step(input, curr.position, curr.direction) {
                    stack.push(State {
                        position: next_pos,
                        direction: curr.direction,
                    })
                }
            }
            b'/' => {
                let next_dir = [
                    Direction::Left,
                    Direction::Up,
                    Direction::Right,
                    Direction::Down,
                ][curr.direction as usize];
                if let Some(next_pos) = step(input, curr.position, next_dir) {
                    stack.push(State {
                        position: next_pos,
                        direction: next_dir,
                    })
                }
            }
            b'\\' => {
                let next_dir = [
                    Direction::Right,
                    Direction::Down,
                    Direction::Left,
                    Direction::Up,
                ][curr.direction as usize];
                if let Some(next_pos) = step(input, curr.position, next_dir) {
                    stack.push(State {
                        position: next_pos,
                        direction: next_dir,
                    })
                }
            }
            b'-' => match curr.direction {
                Direction::Up | Direction::Down => {
                    let next_dir = Direction::Right;
                    if let Some(next_pos) = step(input, curr.position, next_dir) {
                        stack.push(State {
                            position: next_pos,
                            direction: next_dir,
                        })
                    }
                    let next_dir = Direction::Left;
                    if let Some(next_pos) = step(input, curr.position, next_dir) {
                        stack.push(State {
                            position: next_pos,
                            direction: next_dir,
                        })
                    }
                }
                _ => {
                    if let Some(next_pos) = step(input, curr.position, curr.direction) {
                        stack.push(State {
                            position: next_pos,
                            direction: curr.direction,
                        })
                    }
                }
            },
            b'|' => match curr.direction {
                Direction::Right | Direction::Left => {
                    let next_dir = Direction::Up;
                    if let Some(next_pos) = step(input, curr.position, next_dir) {
                        stack.push(State {
                            position: next_pos,
                            direction: next_dir,
                        })
                    }
                    let next_dir = Direction::Down;
                    if let Some(next_pos) = step(input, curr.position, next_dir) {
                        stack.push(State {
                            position: next_pos,
                            direction: next_dir,
                        })
                    }
                }
                _ => {
                    if let Some(next_pos) = step(input, curr.position, curr.direction) {
                        stack.push(State {
                            position: next_pos,
                            direction: curr.direction,
                        })
                    }
                }
            },
            _ => unreachable!(),
        }
    }

    seen.iter()
        .flat_map(|row| row)
        .filter(|&el| el != &[false, false, false, false])
        .count()
}

fn step(input: &Input, position: Position, direction: Direction) -> Option<Position> {
    let (dx, dy): (isize, isize) = match direction {
        Direction::Down => (1, 0),
        Direction::Right => (0, 1),
        Direction::Up => (-1, 0),
        Direction::Left => (0, -1),
    };

    let new_row = position.row.checked_add_signed(dx);
    let new_col = position.col.checked_add_signed(dy);

    match (new_row, new_col) {
        (None, _) | (_, None) => return None,
        (Some(new_row), Some(new_col)) => {
            if new_row >= input.len() || new_col >= input[0].len() {
                return None;
            }

            let next_position = Position {
                row: new_row,
                col: new_col,
            };

            return Some(next_position);
        }
    }
}
