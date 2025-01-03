use std::collections::HashSet;

mod part1;
mod part2;

use part1::part1;
use part2::part2;

struct Scratchcard {
    id: u32,
    winning_nums: HashSet<u32>,
    user_nums: HashSet<u32>,
}

fn main() {
    let input = include_str!("../../inputs/day4.txt");

    let part1 = part1(input);
    let part2 = part2(input);

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Vec<Scratchcard> {
    input
        .lines()
        .map(|line| line.split_once(": ").unwrap())
        .map(|(id, s)| (id.split(" ").last().unwrap(), s.split_once("|").unwrap()))
        .map(|(id, (winning_nums, user_nums))| {
            let id = id.parse::<u32>().unwrap();

            let winning_nums = winning_nums
                .split(" ")
                .filter(|n| !n.is_empty())
                .map(|n| n.parse::<u32>().unwrap())
                .collect::<HashSet<u32>>();

            let user_nums = user_nums
                .split(" ")
                .filter(|n| !n.is_empty())
                .map(|n| n.parse::<u32>().unwrap())
                .collect::<HashSet<u32>>();

            Scratchcard {
                id,
                winning_nums,
                user_nums,
            }
        })
        .collect::<Vec<Scratchcard>>()
}

fn calc_card_result(card: &Scratchcard) -> u32 {
    card.winning_nums.intersection(&card.user_nums).count() as u32
}
