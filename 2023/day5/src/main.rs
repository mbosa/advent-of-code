use std::{env, fs};

mod data;
mod part1;
mod part2;

use anyhow::{Error, Result};
use data::{Input, MapItem};
use part1::part1;
use part2::part2;

fn main() {
    let input_path = env::current_dir().unwrap().join("inputs/day5.txt");

    let input = fs::read_to_string(input_path).unwrap();

    let parsed = parse_input(&input).unwrap();

    let part1 = part1(&parsed);
    let part2 = part2(&parsed);

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Result<Input> {
    let mut i_iter = input.split("\n\n");
    let (_, seeds) = i_iter.next().unwrap().split_once(": ").unwrap();
    let seeds = seeds
        .split(" ")
        .map(|s| s.parse::<u64>().unwrap())
        .collect::<Vec<u64>>();

    let seed_to_soil = i_iter
        .next()
        .ok_or(Error::msg("seed-to-soil map missing"))?
        .lines()
        .skip(1)
        .map(|s| s.parse::<MapItem>())
        .collect::<Result<Vec<MapItem>>>()?;

    let soil_to_fertilizer = i_iter
        .next()
        .ok_or(Error::msg("soil_to_fertilizer map missing"))?
        .lines()
        .skip(1)
        .map(|s| s.parse::<MapItem>())
        .collect::<Result<Vec<MapItem>>>()?;

    let fertilizer_to_water = i_iter
        .next()
        .ok_or(Error::msg("fertilizer_to_water map missing"))?
        .lines()
        .skip(1)
        .map(|s| s.parse::<MapItem>())
        .collect::<Result<Vec<MapItem>>>()?;

    let water_to_light = i_iter
        .next()
        .ok_or(Error::msg("water_to_light map missing"))?
        .lines()
        .skip(1)
        .map(|s| s.parse::<MapItem>())
        .collect::<Result<Vec<MapItem>>>()?;

    let light_to_temperature = i_iter
        .next()
        .ok_or(Error::msg("light_to_temperature map missing"))?
        .lines()
        .skip(1)
        .map(|s| s.parse::<MapItem>())
        .collect::<Result<Vec<MapItem>>>()?;

    let temperature_to_humidity = i_iter
        .next()
        .ok_or(Error::msg("temperature_to_humidity map missing"))?
        .lines()
        .skip(1)
        .map(|s| s.parse::<MapItem>())
        .collect::<Result<Vec<MapItem>>>()?;

    let humidity_to_location = i_iter
        .next()
        .ok_or(Error::msg("humidity_to_location map missing"))?
        .lines()
        .skip(1)
        .map(|s| s.parse::<MapItem>())
        .collect::<Result<Vec<MapItem>>>()?;

    let i = Input {
        seeds,
        seed_to_soil,
        soil_to_fertilizer,
        fertilizer_to_water,
        water_to_light,
        light_to_temperature,
        temperature_to_humidity,
        humidity_to_location,
    };

    Ok(i)
}

fn mapper(item: u64, map: &Vec<MapItem>) -> u64 {
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
