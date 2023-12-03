use std::cmp;

use crate::utils::contains_special_char;

pub fn part1(input: &Vec<Vec<char>>) -> u32 {
    let rows = input.len() as i32;
    let cols = input[0].len() as i32;

    let mut res: u32 = 0;

    for (i, row) in input.iter().enumerate() {
        let mut j = 0i32;

        while j < cols {
            let el = row[j as usize];
            if !el.is_ascii_digit() {
                j += 1;
                continue;
            }

            let mut num = 0u32;
            let start = j;

            while j < cols && row[j as usize].is_ascii_digit() {
                let digit = row[j as usize].to_digit(10).unwrap();

                num = num * 10 + digit;

                j += 1;
            }

            let right_boundary = cmp::min(j + 1, cols) as usize;
            let left_boundary = cmp::max(start - 1, 0) as usize;

            let mut segments: [&[char]; 3] = [&row[left_boundary..right_boundary], &[], &[]];

            if i > 0 {
                // segment above
                segments[1] = &input[i - 1][left_boundary..right_boundary]
            }
            if (i as i32) < rows - 1 {
                // segment below
                segments[2] = &input[i + 1][left_boundary..right_boundary]
            }

            if segments.into_iter().any(contains_special_char) {
                res += num;
            }
        }
    }
    res
}

#[cfg(test)]
mod test {
    use crate::parse_input;

    use super::*;

    #[test]
    fn test_part1() {
        let input = "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..";

        let parsed = parse_input(&input);

        let res = part1(&parsed);

        assert_eq!(res, 4361);
    }
}
