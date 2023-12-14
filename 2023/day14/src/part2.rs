use std::collections::HashMap;

use crate::{calc_north_load, parse_input, tilt_east, tilt_north, tilt_south, tilt_west};

pub fn part2(input: &str) -> usize {
    let mut input = parse_input(input);

    let target_cycles = 1000000000;
    let mut map = HashMap::new();
    let mut cycle_count = 0;

    while cycle_count < target_cycles {
        map.insert(input.clone(), cycle_count);

        cycle_count += 1;
        tilt_north(&mut input);
        tilt_west(&mut input);
        tilt_south(&mut input);
        tilt_east(&mut input);

        if map.contains_key(&input) {
            break;
        }
    }

    if let Some(loop_start) = map.get(&input) {
        let loop_length = cycle_count - loop_start;
        let last_cycles_after_loop = (target_cycles - loop_start) % loop_length;

        for _ in 0..last_cycles_after_loop {
            tilt_north(&mut input);
            tilt_west(&mut input);
            tilt_south(&mut input);
            tilt_east(&mut input);
        }
    }

    calc_north_load(&input)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = "O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....";
        let res = part2(input);

        assert_eq!(res, 64);
    }
}
