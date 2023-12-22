use crate::{flood_bfs, parse_input, Position};

/*
    This only works on the real input, not on the test input.

    Idea taken from reddit.

    The input has some special characteristics that make it solvable:
    * The start position is exactly at the center.
    * There are patterns of empty spaces: horizontal and vertical lines that cross the start, and a diamond pattern.
    * The number of steps makes you reach exactly the start of one of the grids: grid length = 131, start at 65, 26501365 % 131 = 65.

    The logic is as follows:
    * Calculate the result for 3 distances: from the start, to the boundary 0, 1, and 2 grids away.
    * Use Lagrange interpolation to derive a quadratic equation that describes the result for any number of grids away.
    * Use the equation to calculate the result for the target steps, which we know are exactly (26501365 - 65) / 131 grids away.
*/
pub fn part2(input: &str, target_steps: usize) -> usize {
    let input = parse_input(input);

    let size = input.len() as isize;

    let mut start: Position = Position { row: 0, col: 0 };
    for i in 0..input.len() {
        for j in 0..input[0].len() {
            if input[i][j] == 'S' {
                start = Position {
                    row: i as isize,
                    col: j as isize,
                };
            }
        }
    }

    let val0 = flood_bfs(&input, start, (start.row + size * 0) as usize) as isize;
    let val1 = flood_bfs(&input, start, (start.row + size * 1) as usize) as isize;
    let val2 = flood_bfs(&input, start, (start.row + size * 2) as usize) as isize;

    let target_grids = (target_steps - 65) / 131;
    let res = lagrange_interpolation(&[val0, val1, val2], target_grids as isize);

    res as usize
}

fn lagrange_interpolation(values: &[isize], target: isize) -> isize {
    let mut res = 0;

    for i in 0..values.len() {
        let mut term = values[i];
        for j in 0..values.len() {
            let i = i as isize;
            let j = j as isize;
            if i == j {
                continue;
            }
            term *= (target - j) / (i - j);
        }

        res += term
    }

    res
}
