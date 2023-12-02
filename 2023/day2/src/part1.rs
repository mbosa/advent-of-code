use crate::{Game, Rgb};

pub fn part1(input: &Vec<Game>) -> u32 {
    let max_rgb: Rgb = Rgb(12, 13, 14);

    input
        .into_iter()
        .filter(|game| {
            let max_cubes: Rgb = game.rounds.iter().fold(Rgb(0, 0, 0), |mut acc, round| {
                acc.0 = acc.0.max(round.0);
                acc.1 = acc.1.max(round.1);
                acc.2 = acc.2.max(round.2);
                acc
            });

            max_cubes.0 <= max_rgb.0 && max_cubes.1 <= max_rgb.1 && max_cubes.2 <= max_rgb.2
        })
        .map(|game| game.id)
        .sum::<u32>()
}

#[cfg(test)]
mod test {
    use crate::parse_input;

    use super::*;

    #[test]
    fn test_part1() {
        let input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";

        let parsed = parse_input(input).unwrap();

        let res = part1(&parsed);

        assert_eq!(res, 8);
    }
}
