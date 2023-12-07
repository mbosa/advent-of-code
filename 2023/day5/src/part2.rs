use std::cmp;

use crate::data::{Input, MapItem};

pub fn part2(input: &Input) -> i64 {
    let mut outcome: Vec<[i64; 2]> = Vec::new();
    let mut stack = input
        .seeds
        .chunks_exact(2)
        .map(|chunk| [chunk[0], chunk[0] + chunk[1]])
        .collect::<Vec<_>>();

    for maps in &input.maps {
        while let Some(range) = stack.pop() {
            let o = mapper(range, maps, &mut stack);
            outcome.push(o);
        }

        stack = outcome.clone();
        outcome = Vec::new();
    }

    stack.iter().map(|&[from, _to]| from).min().unwrap()
}

fn transform_range([from, to]: [i64; 2], delta: i64) -> [i64; 2] {
    let new_from = from + delta;
    let new_to = to + delta;

    [new_from, new_to]
}

fn mapper(range: [i64; 2], map: &Vec<MapItem>, stack: &mut Vec<[i64; 2]>) -> [i64; 2] {
    for item in map {
        let delta = item.dst_start - item.src_start;
        let range_map = [item.src_start, item.src_start + item.range];

        if range[1] <= range_map[0] || range[0] >= range_map[1] {
            // no intersection
            continue;
        }

        if range[0] >= range_map[0] && range[1] <= range_map[1] {
            // range inside range_map
            let mapped_range: [i64; 2] = transform_range(range, delta);

            return mapped_range;
        } else if range_map[0] > range[0] && range_map[1] < range[1] {
            // range_map inside range
            let mapped_range = transform_range(range_map, delta);
            let left_range = [range[0], range_map[0]];
            let right_range = [range_map[1], range[1]];

            stack.push(left_range);
            stack.push(right_range);

            return mapped_range;
        } else {
            // intersection
            let from = cmp::max(range[0], range_map[0]);
            let to = cmp::min(range[1], range_map[1]);

            let mapped_range = transform_range([from, to], delta);

            let remainder_range = if range[0] < range_map[0] {
                [range[0], range_map[0]]
            } else {
                [range_map[1], range[1]]
            };
            stack.push(remainder_range);

            return mapped_range;
        }
    }
    // no intersection
    range
}

#[cfg(test)]
mod test {
    use crate::parse_input;

    use super::*;

    #[test]
    fn test_part2() {
        let input = "seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4";

        let parsed = parse_input(&input).unwrap();

        let res = part2(&parsed);

        assert_eq!(res, 46);
    }
}
