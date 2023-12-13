use crate::{find_col, find_row, parse_input};

pub fn part1(input: &str) -> usize {
    let input = parse_input(input);

    let mut res = 0;
    for pattern in input {
        let col = find_col(&pattern, 0);

        if let Some(c) = col {
            res += c;
            continue;
        }

        let row = find_row(&pattern, 0);

        if let Some(r) = row {
            res += 100 * r;
        }
    }

    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#";

        let res = part1(input);

        assert_eq!(res, 405);
    }
}
