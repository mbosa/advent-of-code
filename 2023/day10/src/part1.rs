use crate::{get_path_positions_and_mutate_start, parse_input};

pub fn part1(input: &str) -> u32 {
    let mut input = parse_input(input);

    let positions = get_path_positions_and_mutate_start(&mut input);

    (positions.len() / 2) as u32
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ";

        let res = part1(input);

        assert_eq!(res, 8);
    }
}
