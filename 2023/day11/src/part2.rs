use crate::{parse_input, solve};

pub fn part2(input: &str) -> u64 {
    let input = parse_input(input);

    solve(&input, 1_000_000)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
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

        let res = part2(input);

        assert_eq!(res, 82000210);
    }
}
