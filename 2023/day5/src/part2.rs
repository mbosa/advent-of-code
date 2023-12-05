use std::{
    sync::{Arc, Mutex},
    thread,
};

use crate::{data::Input, mapper};

pub fn part2(input: &Input) -> u64 {
    let min: Arc<Mutex<u64>> = Arc::new(Mutex::new(u64::MAX));

    thread::scope(|scope| {
        for chunk in input.seeds.chunks_exact(2) {
            let from = chunk[0];
            let to = chunk[0] + chunk[1];

            println!("start chunk {}", from);

            let shared_min = min.clone();

            scope.spawn(move || {
                println!("spawn thread for chunk {}", from);

                let chunk_min = (from..to)
                    .map(|seed| {
                        let soil = mapper(seed, &input.seed_to_soil);
                        let fertilizer = mapper(soil, &input.soil_to_fertilizer);
                        let water = mapper(fertilizer, &input.fertilizer_to_water);
                        let light = mapper(water, &input.water_to_light);
                        let temperature = mapper(light, &input.light_to_temperature);
                        let humidity = mapper(temperature, &input.temperature_to_humidity);
                        let location = mapper(humidity, &input.humidity_to_location);
                        location
                    })
                    .min()
                    .unwrap();

                let mut guard = shared_min.lock().unwrap();

                *guard = std::cmp::min(*guard, chunk_min);
                println!("finish chunk {}", from);
            });
        }
    });

    let x = *min.lock().unwrap();
    x
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
