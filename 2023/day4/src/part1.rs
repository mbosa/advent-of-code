use crate::{calc_card_result, Scratchcard};

pub fn part1(input: &Vec<Scratchcard>) -> u32 {
    input
        .iter()
        .map(calc_card_result)
        .filter(|&n| n > 0)
        .map(|n| 2u32.pow(n - 1))
        .sum()
}

#[cfg(test)]
mod test {
    use crate::parse_input;

    use super::*;

    #[test]
    fn test_part1() {
        let input = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11";

        let parsed = parse_input(&input);

        let res = part1(&parsed);

        assert_eq!(res, 13);
    }
}
