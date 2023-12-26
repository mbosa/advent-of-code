mod part1;
mod part1_karger;

use part1::part1;
use part1_karger::part1 as part1_karger;

fn main() {
    let input = include_str!("../../inputs/day25.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part1_karger = part1_karger(input);
    println!("part1_karger: {}", part1_karger);
}
