use anyhow::Result;

use crate::parse_input;

pub fn part2(input: &str) -> Result<i32> {
    let input = parse_input(input)?;

    let res = input
        .iter()
        .cloned()
        .map(|mut values| {
            let mut all_zeroes = false;
            let mut start = 0;

            while !all_zeroes {
                all_zeroes = true;

                for i in (start + 1..values.len()).rev() {
                    values[i] = values[i] - values[i - 1];

                    if values[i] != 0 {
                        all_zeroes = false;
                    }
                }
                start += 1;
            }
            values.iter().rev().fold(0, |acc, &v| v - acc)
        })
        .sum();

    Ok(res)
}

pub fn part2_recursive(input: &str) -> Result<i32> {
    let input = parse_input(input)?;

    let res = input
        .iter()
        .map(|a| {
            let q = calc_recursive_part2(a);
            q
        })
        .sum();

    Ok(res)
}

pub fn calc_recursive_part2(sequence: &Vec<i32>) -> i32 {
    if sequence.iter().all(|&v| v == 0) {
        return 0;
    }

    let mut new_sequence = Vec::new();

    for (i, v) in sequence.iter().enumerate().skip(1) {
        new_sequence.push(v - sequence[i - 1]);
    }

    sequence.first().unwrap() - calc_recursive_part2(&new_sequence)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45";

        let res = part2(input).unwrap();

        assert_eq!(res, 2);
    }

    #[test]
    fn test_part2_recursive() {
        let input = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45";

        let res = part2_recursive(input).unwrap();

        assert_eq!(res, 2);
    }
}
