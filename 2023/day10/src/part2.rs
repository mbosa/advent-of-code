use std::collections::HashSet;

use crate::{get_path_positions_and_mutate_start, parse_input};

enum RelativePosition {
    In,
    Out,
}

pub fn part2(input: &str) -> u32 {
    let mut input = parse_input(input);

    let positions: HashSet<[usize; 2]> = get_path_positions_and_mutate_start(&mut input);

    let mut res = 0;

    let mut relative_position = RelativePosition::Out;

    for (i, row) in input.iter().enumerate() {
        for (j, &el) in row.iter().enumerate() {
            if positions.contains(&[i, j]) {
                if let '|' | 'L' | 'J' = el {
                    relative_position = match relative_position {
                        RelativePosition::In => RelativePosition::Out,
                        RelativePosition::Out => RelativePosition::In,
                    };
                }
            } else if let RelativePosition::In = relative_position {
                res += 1;
            }
        }
    }

    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2_1() {
        let input = "...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........";

        let res = part2(input);

        assert_eq!(res, 4);
    }

    #[test]
    fn test_part2_2() {
        let input = "..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........";

        let res = part2(input);

        assert_eq!(res, 4);
    }

    #[test]
    fn test_part2_3() {
        let input = ".F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...";

        let res = part2(input);

        assert_eq!(res, 8);
    }

    #[test]
    fn test_part2_4() {
        let input = "FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L";

        let res = part2(input);

        assert_eq!(res, 10);
    }
}
