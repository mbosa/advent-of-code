use crate::{parse_input, solve};

pub fn part1(input: &str) -> u64 {
    let input = parse_input(input);

    solve(&input, 2)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....";

        let res = part1(input);

        assert_eq!(res, 374);
    }
}
