use crate::data::{Input, MapItem};

pub fn part1(input: &Input) -> i64 {
    input
        .seeds
        .iter()
        .map(|&seed| input.maps.iter().fold(seed, |acc, maps| mapper(acc, maps)))
        .min()
        .unwrap()
}

fn mapper(item: i64, map: &Vec<MapItem>) -> i64 {
    for &MapItem {
        dst_start,
        src_start,
        range,
    } in map
    {
        if item < src_start {
            continue;
        }

        if item >= src_start + range {
            continue;
        }

        return item - src_start + dst_start;
    }

    item
}

#[cfg(test)]
mod test {
    use crate::parse_input;

    use super::*;

    #[test]
    fn test_part1() {
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

        let res = part1(&parsed);

        assert_eq!(res, 35);
    }
}
