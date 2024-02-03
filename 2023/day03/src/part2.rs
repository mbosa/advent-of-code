use crate::{parse_input, utils::slice_char_to_u32};

pub fn part2(input: &str) -> u32 {
    let input = parse_input(input);

    let rows = input.len() as i32;
    let cols = input[0].len() as i32;

    let mut res = 0u32;

    for (i, row) in input.iter().enumerate() {
        let i = i as i32;
        let mut j = 0i32;

        while j < cols {
            let el = row[j as usize];
            if el != '*' {
                j += 1;
                continue;
            }

            let mut nums_around = Vec::<u32>::new();

            let mut segments: [&[char]; 3] = [&[], &[], &[]];

            for (k, row_i) in [i, i - 1, i + 1].into_iter().enumerate() {
                if row_i < 0 || row_i >= rows {
                    continue;
                }

                let row = &input[row_i as usize];

                // find leftmost digit
                let mut l = j - 1;
                while l >= 0 && row[l as usize].is_ascii_digit() {
                    l -= 1;
                }
                l += 1;
                // find rightmost digit
                let mut r = j + 1;
                while r < cols && row[r as usize].is_ascii_digit() {
                    r += 1;
                }
                r -= 1;

                segments[k] = &row[l as usize..=r as usize];
            }

            // extract the numbers from all the segments
            segments
                .into_iter()
                .flat_map(|s| s.split(|&c| !c.is_ascii_digit()))
                .filter(|c| c.len() > 0)
                .map(slice_char_to_u32)
                .for_each(|n| nums_around.push(n));

            if nums_around.len() == 2 {
                let ratio = nums_around[0] * nums_around[1];

                res += ratio;
            }

            j += 1;
        }
    }
    res
}

#[cfg(test)]
mod test {
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

        let res = part2(input);

        assert_eq!(res, 467835);
    }
}
