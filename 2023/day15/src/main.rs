mod part1;
mod part2;

use part1::part1;
use part2::part2;

fn main() {
    let input = include_str!("../../inputs/day15.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn calc_hash(s: &str) -> usize {
    s.bytes()
        .fold(0, |acc, c| u8::wrapping_add(acc, c).wrapping_mul(17)) as usize
}
