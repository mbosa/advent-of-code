use crate::{flood_bfs, parse_input, Position};

#[derive(Debug, Copy, Clone, PartialEq, Eq, Hash)]
struct Item {
    count: usize,
    pos: Position,
}

pub fn part1(input: &str, target_steps: usize) -> usize {
    let input = parse_input(input);

    let mut start: Position = Position { row: 0, col: 0 };
    for i in 0..input.len() {
        for j in 0..input[0].len() {
            if input[i][j] == 'S' {
                start = Position {
                    row: i as isize,
                    col: j as isize,
                };
            }
        }
    }

    flood_bfs(&input, start, target_steps)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
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

        let res = part1(input, 6);

        assert_eq!(res, 16);
    }
}
