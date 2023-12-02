const NUMS_WORDS: [&str; 9] = [
    "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
];

pub fn part2(input: &str) -> u32 {
    input
        .lines()
        .flat_map(|line| find_first_digit(line).zip(find_last_digit(line)))
        .map(|(first_letter, last_letter)| first_letter * 10 + last_letter)
        .sum::<u32>()
}

fn find_first_digit(line: &str) -> Option<u32> {
    for i in 0..line.len() {
        let substr = &line[i..];

        if let Some(first_digit) = substr.chars().next().map(|c| c.to_digit(10)).flatten() {
            return Some(first_digit);
        }

        for (i, &word) in NUMS_WORDS.iter().enumerate() {
            if substr.starts_with(word) {
                return Some(1 + i as u32);
            }
        }
    }
    None
}

fn find_last_digit(line: &str) -> Option<u32> {
    for i in (0..line.len() + 1).rev() {
        let substr = &line[..i];

        if let Some(first_digit) = substr.chars().last().map(|c| c.to_digit(10)).flatten() {
            return Some(first_digit);
        }

        for (i, &word) in NUMS_WORDS.iter().enumerate() {
            if substr.ends_with(word) {
                return Some(1 + i as u32);
            }
        }
    }
    None
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";

        let res = part2(input);

        assert_eq!(res, 281);
    }
}
