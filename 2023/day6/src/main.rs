mod part1;
mod part2;

use part1::part1;
use part2::part2;

struct Race {
    time: u64,
    distance: u64,
}

fn main() {
    let input = include_str!("../../inputs/day6.txt");

    let part1 = part1(&input);
    let part2 = part2(&input);

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn calc_ways_to_win(race: &Race) -> u64 {
    let &Race { time, distance } = race;

    let root = ((time.pow(2) - 4 * distance) as f64).sqrt();
    let sol_1 = (((time as f64) - root) / 2.0).floor() as u64;
    let sol_2 = (((time as f64) + root) / 2.0).ceil() as u64;

    sol_2 - sol_1 - 1
}
