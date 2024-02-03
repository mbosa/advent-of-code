use crate::{calc_ways_to_win, Race};

pub fn part2(input: &str) -> u64 {
    let parsed = parse_input_part2(input);

    calc_ways_to_win(&parsed)
}

fn parse_input_part2(input: &str) -> Race {
    let mut parsed_lines = input
        .lines()
        .map(|line| {
            line.split_whitespace()
                .skip(1)
                .fold(String::new(), |acc, el| acc + el)
        })
        .map(|n| n.parse::<u64>().unwrap());

    let time = parsed_lines.next().unwrap();
    let distance = parsed_lines.next().unwrap();

    Race { time, distance }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = "Time:      7  15   30
Distance:  9  40  200";

        let res = part2(&input);

        assert_eq!(res, 71503);
    }
}
