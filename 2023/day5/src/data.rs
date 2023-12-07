use std::str::FromStr;

use anyhow::Error;

#[derive(Debug)]
pub struct MapItem {
    pub dst_start: i64,
    pub src_start: i64,
    pub range: i64,
}

#[derive(Debug)]
pub struct Input {
    pub seeds: Vec<i64>,
    pub maps: Vec<Vec<MapItem>>,
}

impl FromStr for MapItem {
    type Err = Error;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut s_split = s.split(" ");

        let dst_start = s_split
            .next()
            .ok_or(Error::msg("Error: missing dst_start"))?
            .parse::<i64>()?;

        let src_start = s_split
            .next()
            .ok_or(Error::msg("Error: missing src_start"))?
            .parse::<i64>()?;

        let range = s_split
            .next()
            .ok_or(Error::msg("Error: missing range"))?
            .parse::<i64>()?;

        let map_item = MapItem {
            dst_start,
            src_start,
            range,
        };

        Ok(map_item)
    }
}
