use std::str::FromStr;

use anyhow::Error;

#[derive(Debug)]
pub struct MapItem {
    pub dst_start: u64,
    pub src_start: u64,
    pub range: u64,
}

#[derive(Debug)]
pub struct Input {
    pub seeds: Vec<u64>,
    pub seed_to_soil: Vec<MapItem>,
    pub soil_to_fertilizer: Vec<MapItem>,
    pub fertilizer_to_water: Vec<MapItem>,
    pub water_to_light: Vec<MapItem>,
    pub light_to_temperature: Vec<MapItem>,
    pub temperature_to_humidity: Vec<MapItem>,
    pub humidity_to_location: Vec<MapItem>,
}

impl FromStr for MapItem {
    type Err = Error;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut s_split = s.split(" ");

        let dst_start = s_split
            .next()
            .ok_or(Error::msg("Error: missing dst_start"))?
            .parse::<u64>()?;

        let src_start = s_split
            .next()
            .ok_or(Error::msg("Error: missing src_start"))?
            .parse::<u64>()?;

        let range = s_split
            .next()
            .ok_or(Error::msg("Error: missing range"))?
            .parse::<u64>()?;

        let map_item = MapItem {
            dst_start,
            src_start,
            range,
        };

        Ok(map_item)
    }
}
