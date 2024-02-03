pub fn part1(input: &str) -> u32 {
    input
        .lines()
        .map(|line| line.chars().filter_map(|c| c.to_digit(10)))
        .map(|iter| iter.clone().zip(iter.rev()))
        .flat_map(|mut iter| iter.next())
        .map(|(a, b)| a * 10 + b)
        .sum::<u32>()
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";

        let res = part1(input);

        assert_eq!(res, 142);
    }
}
