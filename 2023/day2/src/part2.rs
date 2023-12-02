use crate::{Game, Rgb};

pub fn part2(input: &Vec<Game>) -> u32 {
    input
        .into_iter()
        .map(|game| {
            let max_cubes: Rgb = game.rounds.iter().fold(Rgb(0, 0, 0), |mut acc, round| {
                acc.0 = acc.0.max(round.0);
                acc.1 = acc.1.max(round.1);
                acc.2 = acc.2.max(round.2);
                acc
            });

            max_cubes
        })
        .map(|cubes| cubes.0 * cubes.1 * cubes.2)
        .sum::<u32>()
}
#[cfg(test)]
mod test {
    use crate::parse_input;

    use super::*;

    #[test]
    fn test_part2() {
        let input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";

        let parsed = parse_input(input).unwrap();

        let res = part2(&parsed);

        assert_eq!(res, 2286);
    }
}
