use crate::{calc_ways_to_win, Race};

pub fn part1(input: &str) -> u64 {
    let parsed = parse_input_part1(input);

    parsed.iter().map(calc_ways_to_win).product()
}

fn parse_input_part1(input: &str) -> Vec<Race> {
    let mut parsed_lines = input.lines().map(|line| {
        line.split_whitespace()
            .skip(1)
            .map(|n| n.parse::<u64>().unwrap())
    });

    let times = parsed_lines.next().unwrap();
    let distances = parsed_lines.next().unwrap();

    times
        .zip(distances)
        .map(|(time, distance)| Race { time, distance })
        .collect()
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "Time:      7  15   30
Distance:  9  40  200";

        let res = part1(&input);

        assert_eq!(res, 288);
    }
}
