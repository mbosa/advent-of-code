use crate::{calc_card_result, parse_input};

pub fn part2(input: &str) -> u32 {
    let input = parse_input(input);

    let values = input.iter().map(calc_card_result).collect::<Vec<u32>>();
    let mut stack = input.iter().map(|card| card.id).rev().collect::<Vec<u32>>();

    let mut res = 0;

    while let Some(el) = stack.pop() {
        res += 1;
        let val = values[(el - 1) as usize];

        for i in 1..val + 1 {
            stack.push(el + i);
        }
    }

    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11";

        let res = part2(input);

        assert_eq!(res, 30);
    }
}
