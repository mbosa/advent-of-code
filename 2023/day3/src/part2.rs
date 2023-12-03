use crate::utils::{calc_num_left, calc_num_right, Pos};

pub fn part2(input: &Vec<Vec<char>>) -> i32 {
    let rows = input.len() as i32;
    let cols = input[0].len() as i32;

    let mut res: i32 = 0;

    for (i, row) in input.iter().enumerate() {
        let mut j = cols - 1;

        while j >= 0 {
            let el = row[j as usize];
            if el != '*' {
                j -= 1;
                continue;
            }

            // el is a gear

            let mut nums_around = Vec::<i32>::new();

            // find num on the left
            if let Some(num) = calc_num_left(
                input,
                Pos {
                    row: i as i32,
                    col: j - 1,
                },
            ) {
                nums_around.push(num);
            }

            // find num on the right
            if let Some(num) = calc_num_right(
                input,
                Pos {
                    row: i as i32,
                    col: j + 1,
                },
            ) {
                nums_around.push(num);
            }

            // up
            if i > 0 {
                let row_up = &input[i - 1];
                let left = 0.max(j - 1) as usize;
                let right = (j + 2).min(cols) as usize;
                let segment_up = &input[i - 1][left..right];

                match (
                    segment_up[0].is_ascii_digit(),
                    segment_up[1].is_ascii_digit(),
                    segment_up[2].is_ascii_digit(),
                ) {
                    (false, false, false) => {}
                    (true, false, true) => {
                        // find num on the left
                        if let Some(num) = calc_num_left(
                            input,
                            Pos {
                                row: (i - 1) as i32,
                                col: j - 1,
                            },
                        ) {
                            nums_around.push(num);
                        }

                        // find num on the right
                        if let Some(num) = calc_num_right(
                            input,
                            Pos {
                                row: (i - 1) as i32,
                                col: j + 1,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                    (false, false, true) => {
                        // find num on the right
                        if let Some(num) = calc_num_right(
                            input,
                            Pos {
                                row: (i - 1) as i32,
                                col: j + 1,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                    (false, true, _) => {
                        // find num on the left
                        // find num on the right
                        if let Some(num) = calc_num_right(
                            input,
                            Pos {
                                row: (i - 1) as i32,
                                col: j,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                    (true, false, false) => {
                        // find num on the left
                        if let Some(num) = calc_num_left(
                            input,
                            Pos {
                                row: (i - 1) as i32,
                                col: j - 1,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                    (_, true, false) => {
                        // find num on the left
                        if let Some(num) = calc_num_left(
                            input,
                            Pos {
                                row: (i - 1) as i32,
                                col: j,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                    (true, true, true) => {
                        // find leftmost digit
                        let mut k = j + 1;
                        while k >= 0 && row_up[k as usize].is_ascii_digit() {
                            k -= 1;
                        }

                        // find num on the right
                        if let Some(num) = calc_num_right(
                            input,
                            Pos {
                                row: (i - 1) as i32,
                                col: k + 1,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                }
            }

            // down
            if i < (rows as usize) - 1 {
                let row_down = &input[i + 1];
                let left = 0.max(j - 1) as usize;
                let right = (j + 2).min(cols) as usize;
                let segment_down = &row_down[left..right];

                match (
                    segment_down[0].is_ascii_digit(),
                    segment_down[1].is_ascii_digit(),
                    segment_down[2].is_ascii_digit(),
                ) {
                    (false, false, false) => {}
                    (true, false, true) => {
                        // find num on the left
                        if let Some(num) = calc_num_left(
                            input,
                            Pos {
                                row: (i + 1) as i32,
                                col: j - 1,
                            },
                        ) {
                            nums_around.push(num);
                        }

                        // find num on the right
                        if let Some(num) = calc_num_right(
                            input,
                            Pos {
                                row: (i + 1) as i32,
                                col: j + 1,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                    (false, false, true) => {
                        // find num on the right
                        if let Some(num) = calc_num_right(
                            input,
                            Pos {
                                row: (i + 1) as i32,
                                col: j + 1,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                    (false, true, _) => {
                        // find num on the left
                        // find num on the right
                        if let Some(num) = calc_num_right(
                            input,
                            Pos {
                                row: (i + 1) as i32,
                                col: j,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                    (true, false, false) => {
                        // find num on the left
                        if let Some(num) = calc_num_left(
                            input,
                            Pos {
                                row: (i + 1) as i32,
                                col: j - 1,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                    (_, true, false) => {
                        // find num on the left
                        if let Some(num) = calc_num_left(
                            input,
                            Pos {
                                row: (i + 1) as i32,
                                col: j,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                    (true, true, true) => {
                        // find leftmost digit
                        let mut k = j + 1;
                        while k >= 0 && row_down[k as usize].is_ascii_digit() {
                            k -= 1;
                        }

                        // find num on the right
                        if let Some(num) = calc_num_right(
                            input,
                            Pos {
                                row: (i + 1) as i32,
                                col: k + 1,
                            },
                        ) {
                            nums_around.push(num);
                        }
                    }
                }
            }

            if nums_around.len() == 2 {
                let ratio = nums_around[0] * nums_around[1];

                res += ratio;
            }

            j -= 1;
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

        let res = part2(&parsed);

        assert_eq!(res, 467835);
    }
}
