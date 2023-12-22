mod part1;
mod part2;

use std::collections::{HashSet, VecDeque};

use part1::part1;
use part2::part2;

#[derive(Debug, Copy, Clone, PartialEq, Eq, Hash)]
struct Position {
    row: isize,
    col: isize,
}

#[derive(Debug, Copy, Clone, PartialEq, Eq, Hash)]
struct QueueItem {
    count: usize,
    pos: Position,
}

fn main() {
    let input = include_str!("../../inputs/day21.txt");

    let part1 = part1(input, 64);
    println!("part1: {}", part1);

    let part2 = part2(input, 26501365);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Vec<Vec<char>> {
    input.lines().map(|line| line.chars().collect()).collect()
}

fn flood_bfs(input: &Vec<Vec<char>>, start: Position, target_steps: usize) -> usize {
    let rows = input.len() as isize;
    let cols = input[0].len() as isize;

    let is_end_position = |count: usize| -> bool { target_steps % 2 == count % 2 };

    let mut res = 0;
    let mut seen = HashSet::new();
    let mut q = VecDeque::new();
    q.push_back(QueueItem {
        count: 0,
        pos: start,
    });

    while let Some(item) = q.pop_front() {
        if !seen.insert(item.pos) {
            continue;
        }

        if is_end_position(item.count) {
            res += 1;
        }

        if item.count > target_steps {
            continue;
        }

        for [drow, dcol] in [[-1, 0], [1, 0], [0, -1], [0, 1]] {
            let next_position = Position {
                row: item.pos.row + drow,
                col: item.pos.col + dcol,
            };

            let original_row = ((next_position.row % rows) + rows) % rows;
            let original_col = ((next_position.col % cols) + cols) % cols;

            if input[original_row as usize][original_col as usize] != '#' {
                let new_item = QueueItem {
                    count: item.count + 1,
                    pos: next_position,
                };
                q.push_back(new_item);
            }
        }
    }

    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = "...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........";

        let parsed_input = parse_input(input);
        let start = Position { row: 5, col: 5 };

        let res = flood_bfs(&parsed_input, start, 1);
        assert_eq!(res, 2);
        let res = flood_bfs(&parsed_input, start, 3);
        assert_eq!(res, 6);
        let res = flood_bfs(&parsed_input, start, 6);
        assert_eq!(res, 16);
        let res = flood_bfs(&parsed_input, start, 10);
        assert_eq!(res, 50);
        let res = flood_bfs(&parsed_input, start, 50);
        assert_eq!(res, 1594);
        let res = flood_bfs(&parsed_input, start, 100);
        assert_eq!(res, 6536);
        let res = flood_bfs(&parsed_input, start, 500);
        assert_eq!(res, 167004);
    }
}
