use crate::{calc_north_load, parse_input, tilt_north};

#[derive(Debug, Eq, PartialEq, Hash)]
struct Position {
    row: usize,
    col: usize,
}

pub fn part1(input: &str) -> usize {
    let mut input = parse_input(input);

    tilt_north(&mut input);

    calc_north_load(&input)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....";

        let res = part1(input);

        assert_eq!(res, 136);
    }
}
