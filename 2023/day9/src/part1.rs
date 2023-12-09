use anyhow::Result;

use crate::parse_input;

pub fn part1(input: &str) -> Result<i32> {
    let input = parse_input(input)?;

    let res = input
        .iter()
        .cloned()
        .map(|mut values| {
            let mut all_zeroes = false;
            let mut len = values.len();

            while !all_zeroes {
                all_zeroes = true;

                for i in 0..len - 1 {
                    values[i] = values[i + 1] - values[i];

                    if values[i] != 0 {
                        all_zeroes = false;
                    }
                }
                len -= 1;
            }

            values.iter().sum::<i32>()
        })
        .sum();

    Ok(res)
}

pub fn part1_recursive(input: &str) -> Result<i32> {
    let input = parse_input(input)?;

    let res = input.iter().map(calc_recursive_part1).sum();

    Ok(res)
}

fn calc_recursive_part1(sequence: &Vec<i32>) -> i32 {
    if sequence.iter().all(|&v| v == 0) {
        return 0;
    }

    let mut new_sequence = Vec::new();

    for (i, v) in sequence.iter().enumerate().skip(1) {
        new_sequence.push(v - sequence[i - 1]);
    }

    sequence.last().unwrap() + calc_recursive_part1(&new_sequence)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45";

        let res = part1(input).unwrap();

        assert_eq!(res, 114);
    }

    #[test]
    fn test_part1_recursive() {
        let input = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45";

        let res = part1_recursive(input).unwrap();

        assert_eq!(res, 114);
    }
}
