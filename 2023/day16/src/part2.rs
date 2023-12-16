use crate::{parse_input, run_beam, Direction, Position, State};

pub fn part2(input: &str) -> usize {
    let input = parse_input(input);

    let left_col_starts = (0..input.len()).map(|i| State {
        position: Position { row: i, col: 0 },
        direction: Direction::Right,
    });
    let right_col_starts = (0..input.len()).map(|i| State {
        position: Position {
            row: i,
            col: input[0].len() - 1,
        },
        direction: Direction::Left,
    });
    let top_row_starts = (0..input[0].len()).map(|j| State {
        position: Position { row: 0, col: j },
        direction: Direction::Down,
    });
    let bottom_row_starts = (0..input.len()).map(|j| State {
        position: Position {
            row: input.len() - 1,
            col: j,
        },
        direction: Direction::Up,
    });

    left_col_starts
        .chain(right_col_starts)
        .chain(top_row_starts)
        .chain(bottom_row_starts)
        .map(|start| run_beam(&input, start))
        .max()
        .unwrap()
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = ".|...\\....
|.-.\\.....
.....|-...
........|.
..........
.........\\
..../.\\\\..
.-.-/..|..
.|....-|.\\
..//.|....";

        let res = part2(input);

        assert_eq!(res, 51);
    }
}
