mod part1;
mod part2;
mod part2_bruteforce;

use part1::part1;
use part2::part2;
use part2_bruteforce::part2 as part2_bruteforce;

#[derive(Debug, Copy, Clone)]
struct Position {
    x: i64,
    y: i64,
    z: i64,
}

#[derive(Debug, Copy, Clone)]
struct Velocity {
    vx: i64,
    vy: i64,
    vz: i64,
}

#[derive(Debug, Copy, Clone)]
struct Hail {
    position: Position,
    velocity: Velocity,
}

type Input = Vec<Hail>;

fn main() {
    let input = include_str!("../../inputs/day24.txt");

    let part1 = part1(input, 200000000000000.0, 400000000000000.0);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);

    let part2_bruteforce = part2_bruteforce(input);
    println!("part2_bruteforce: {}", part2_bruteforce);
}

fn parse_input(input: &str) -> Input {
    input
        .lines()
        .map(|line| {
            let (p, v) = line.split_once(" @ ").unwrap();
            let mut ps = p.split(", ").map(|n| n.parse::<i64>().unwrap());
            let mut vs = v.split(", ").map(|n| n.trim().parse::<i64>().unwrap());

            let position = Position {
                x: ps.next().unwrap(),
                y: ps.next().unwrap(),
                z: ps.next().unwrap(),
            };
            let velocity = Velocity {
                vx: vs.next().unwrap(),
                vy: vs.next().unwrap(),
                vz: vs.next().unwrap(),
            };
            Hail { position, velocity }
        })
        .collect::<Vec<_>>()
}
