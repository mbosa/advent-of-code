use crate::{parse_input, run_beam, Direction, Position, State};

pub fn part1(input: &str) -> usize {
    let input = parse_input(input);

    let start = State {
        position: Position { row: 0, col: 0 },
        direction: Direction::Right,
    };

    run_beam(&input, start)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
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

        let res = part1(input);

        assert_eq!(res, 46);
    }
}
