mod part1;
mod part2;

use part1::part1;
use part2::part2;

fn main() {
    let input = include_str!("../../inputs/day1.txt");

    let part1 = part1(input);
    let part2 = part2(input);

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}
