use crate::{parse_input, Direction, Position};

pub fn part2(input: &str) -> i64 {
    let input = parse_input(input);

    let mut tot_area = 0;
    let mut boundary = 0;

    let mut prev: Position = Position { row: 0, col: 0 };

    for dig in input {
        let n = &dig.color[1..dig.color.len() - 1];
        let n = i64::from_str_radix(n, 16).unwrap();

        let dir = &dig.color[dig.color.len() - 1..];
        let dir = match dir {
            "0" => Direction::Right,
            "1" => Direction::Down,
            "2" => Direction::Left,
            "3" => Direction::Up,
            _ => unreachable!(),
        };

        let curr = match dir {
            Direction::Up => Position {
                row: prev.row - n,
                col: prev.col,
            },
            Direction::Down => Position {
                row: prev.row + n,
                col: prev.col,
            },
            Direction::Right => Position {
                row: prev.row,
                col: prev.col + n,
            },
            Direction::Left => Position {
                row: prev.row,
                col: prev.col - n,
            },
        };

        // shoelace formula
        tot_area += prev.row * curr.col - prev.col * curr.row;
        boundary += n;

        prev = curr;
    }

    tot_area = tot_area.abs() / 2;
    // Pick theorem
    // total_area = inside_area + boundary/2 - 1 -> inside_area = total_area + 1 - boundary/2
    let inside_area = tot_area + 1 - boundary / 2;

    inside_area + boundary
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = "R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)";

        let res = part2(input);

        assert_eq!(res, 952408144115);
    }
}
