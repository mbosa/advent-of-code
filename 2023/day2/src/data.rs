use anyhow::Error;
use lazy_static::lazy_static;
use regex::Regex;
use std::str::FromStr;

lazy_static! {
    static ref ROUND_RE: Regex =
        Regex::new(r"(?P<reds>\d+) red|(?P<greens>\d+) green|(?P<blues>\d+) blue")
            .expect("ROUND_RE regex cannot be compiled");
    static ref GAME_ID_RE: Regex =
        Regex::new(r"Game (?P<id>\d+)").expect("GAME_ID_RE regex cannot be compiled");
}

#[derive(Debug)]
pub struct Rgb(pub u32, pub u32, pub u32);

#[derive(Debug)]
pub struct Game {
    pub id: u32,
    pub rounds: Vec<Rgb>,
}

impl FromStr for Game {
    type Err = Error;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut parts = s.split(":");
        let game_id = parts.next().ok_or(Error::msg("Error BadFormat"))?;
        let rounds_str = parts.next().ok_or(Error::msg("Error BadFormat"))?;

        let id = GAME_ID_RE
            .captures(game_id)
            .and_then(|cap| cap.name("id"))
            .ok_or(Error::msg("Error BadFormat"))?
            .as_str()
            .parse::<u32>()?;

        let mut rounds: Vec<Rgb> = Vec::new();

        for r in rounds_str.split(";") {
            let mut rgb = Rgb(0, 0, 0);

            for cap in ROUND_RE.captures_iter(r) {
                if let Some(r) = cap.name("reds") {
                    rgb.0 = r.as_str().parse::<u32>()?;
                }
                if let Some(r) = cap.name("greens") {
                    rgb.1 = r.as_str().parse::<u32>()?;
                }
                if let Some(r) = cap.name("blues") {
                    rgb.2 = r.as_str().parse::<u32>()?;
                }
            }

            rounds.push(rgb);
        }

        Ok(Game { id, rounds })
    }
}
