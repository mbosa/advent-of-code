mod part1;
mod part2;

use std::collections::HashSet;

use part1::part1;
use part2::part2;

type Input = Vec<Vec<char>>;

type Position = [usize; 2];

enum Directions {
    Up,
    Down,
    Right,
    Left,
}

struct DirectionDeltas {
    up: [i32; 2],
    down: [i32; 2],
    right: [i32; 2],
    left: [i32; 2],
}

struct ValidLinks {
    up: [char; 3],
    down: [char; 3],
    right: [char; 3],
    left: [char; 3],
}

const DIRECTION_DELTAS: DirectionDeltas = DirectionDeltas {
    up: [-1, 0],
    down: [1, 0],
    right: [0, 1],
    left: [0, -1],
};

const VALID_LINKS: ValidLinks = ValidLinks {
    up: ['|', '7', 'F'],
    down: ['|', 'L', 'J'],
    right: ['-', '7', 'J'],
    left: ['-', 'L', 'F'],
};

fn main() {
    let input = include_str!("../../inputs/day10.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Input {
    input.lines().map(|line| line.chars().collect()).collect()
}

fn get_path_positions_and_mutate_start(input: &mut Input) -> HashSet<Position> {
    let mut positions = HashSet::new();

    let start = find_start(input);
    positions.insert(start);

    let start_replacement = calc_start_replacement(input, start);
    input[start[0]][start[1]] = start_replacement;

    // first move from start
    let mut direction = match start_replacement {
        '|' | 'L' | 'J' => Directions::Up,
        '7' | 'F' => Directions::Down,
        '-' => Directions::Right,
        _ => unreachable!(),
    };

    let mut curr = match direction {
        Directions::Up => calc_position(input, start, DIRECTION_DELTAS.up).unwrap(),
        Directions::Down => calc_position(input, start, DIRECTION_DELTAS.down).unwrap(),
        _ => unreachable!(),
    };
    positions.insert(curr);

    // walk around the loop
    while curr != start {
        match input[curr[0]][curr[1]] {
            '|' => {}
            'L' => {
                direction = match direction {
                    Directions::Down => Directions::Right,
                    Directions::Left => Directions::Up,
                    _ => unreachable!(),
                };
            }
            'J' => {
                direction = match direction {
                    Directions::Down => Directions::Left,
                    Directions::Right => Directions::Up,
                    _ => unreachable!(),
                };
            }
            '7' => {
                direction = match direction {
                    Directions::Up => Directions::Left,
                    Directions::Right => Directions::Down,
                    _ => unreachable!(),
                };
            }
            'F' => {
                direction = match direction {
                    Directions::Up => Directions::Right,
                    Directions::Left => Directions::Down,
                    _ => unreachable!(),
                };
            }
            '-' => {}
            _ => unreachable!(),
        }

        match direction {
            Directions::Up => curr = calc_position(input, curr, DIRECTION_DELTAS.up).unwrap(),
            Directions::Down => curr = calc_position(input, curr, DIRECTION_DELTAS.down).unwrap(),
            Directions::Right => curr = calc_position(input, curr, DIRECTION_DELTAS.right).unwrap(),
            Directions::Left => curr = calc_position(input, curr, DIRECTION_DELTAS.left).unwrap(),
        }
        positions.insert(curr);
    }

    positions
}

fn find_start(input: &Input) -> Position {
    for (i, row) in input.iter().enumerate() {
        for (j, &el) in row.iter().enumerate() {
            if el == 'S' {
                return [i, j];
            }
        }
    }

    unreachable!()
}

fn calc_position(input: &Input, curr_pos: [usize; 2], delta: [i32; 2]) -> Option<Position> {
    if curr_pos[0] as i32 + delta[0] < 0 {
        return None;
    }
    if curr_pos[0] as i32 + delta[0] > input.len() as i32 {
        return None;
    }
    if curr_pos[1] as i32 + delta[1] < 0 {
        return None;
    }
    if curr_pos[1] as i32 + delta[1] > input[0].len() as i32 {
        return None;
    }

    let next_row = (curr_pos[0] as i32 + delta[0]) as usize;
    let next_col = (curr_pos[1] as i32 + delta[1]) as usize;

    Some([next_row, next_col])
}

fn calc_start_replacement(input: &Input, start: Position) -> char {
    let linked_up = calc_position(input, start, DIRECTION_DELTAS.up)
        .map_or(false, |[i, j]| VALID_LINKS.up.contains(&input[i][j]));
    let linked_down = calc_position(input, start, DIRECTION_DELTAS.down)
        .map_or(false, |[i, j]| VALID_LINKS.down.contains(&input[i][j]));
    let linked_right = calc_position(input, start, DIRECTION_DELTAS.right)
        .map_or(false, |[i, j]| VALID_LINKS.right.contains(&input[i][j]));
    let linked_left = calc_position(input, start, DIRECTION_DELTAS.left)
        .map_or(false, |[i, j]| VALID_LINKS.left.contains(&input[i][j]));

    match (linked_up, linked_down, linked_right, linked_left) {
        (true, true, _, _) => '|',
        (true, false, true, _) => 'L',
        (true, false, false, true) => 'J',
        (false, true, true, _) => 'F',
        (false, true, false, true) => '7',
        (false, false, true, true) => '-',
        _ => unreachable!(),
    }
}
