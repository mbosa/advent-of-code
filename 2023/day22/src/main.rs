mod part1;
mod part2;

use std::cmp::Ordering;
use std::collections::VecDeque;

use part1::part1;
use part2::part2;

#[derive(Debug, Clone, Copy, PartialEq, Eq, Ord, PartialOrd)]
struct Position {
    x: u32,
    y: u32,
    z: u32,
}

struct Bricks {
    bricks: Vec<Brick>,
}

#[derive(Debug, Clone, PartialEq, Eq, Ord, PartialOrd)]
struct Brick {
    start: Position,
    end: Position,
    supports: Vec<usize>,
    supported_by: Vec<usize>,
}

fn main() {
    let input = include_str!("../../inputs/day22.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Bricks {
    let bricks = input
        .lines()
        .map(|line| {
            let (start, end) = line.split_once('~').unwrap();

            let mut start = start.split(',').map(|n| n.parse::<u32>().unwrap());
            let start = Position {
                x: start.next().unwrap(),
                y: start.next().unwrap(),
                z: start.next().unwrap(),
            };

            let mut end = end.split(',').map(|n| n.parse::<u32>().unwrap());
            let end = Position {
                x: end.next().unwrap(),
                y: end.next().unwrap(),
                z: end.next().unwrap(),
            };

            Brick {
                start,
                end,
                supports: Vec::new(),
                supported_by: Vec::new(),
            }
        })
        .collect::<Vec<_>>();

    Bricks { bricks }
}

impl Bricks {
    fn len(&self) -> usize {
        self.bricks.len()
    }

    fn drop_bricks(&mut self) {
        let bricks = &mut self.bricks;

        bricks.sort_unstable_by(|a, b| {
            if a.start.z < b.start.z {
                Ordering::Less
            } else {
                Ordering::Greater
            }
        });

        for i in 0..bricks.len() {
            // max drop -> drop to z = 1
            let mut drop = bricks[i].start.z - 1;

            for j in (0..i).rev() {
                if (bricks[i].start.x <= bricks[j].end.x && bricks[i].end.x >= bricks[j].start.x)
                    && (bricks[i].start.y <= bricks[j].end.y
                        && bricks[i].end.y >= bricks[j].start.y)
                {
                    let new_drop = bricks[i].start.z - bricks[j].end.z - 1;

                    if new_drop == drop {
                        bricks[i].supported_by.push(j);
                    } else if new_drop < drop {
                        bricks[i].supported_by.clear();
                        bricks[i].supported_by.push(j);

                        drop = new_drop;
                    }
                }
            }

            bricks[i].start.z -= drop;
            bricks[i].end.z -= drop;

            for k in 0..bricks[i].supported_by.len() {
                let brick_i = bricks[i].supported_by[k];
                bricks[brick_i].supports.push(i);
            }
        }
    }

    /// Returns how many bricks would fall if you destroy the brick in position `i`
    fn count_chain_reaction(&self, i: usize) -> usize {
        let bricks = &self.bricks;
        let mut res = 0;

        let mut queue = VecDeque::new();
        let mut moved = vec![false; bricks.len()];

        queue.push_back(i);

        while let Some(brick_i) = queue.pop_front() {
            moved[brick_i] = true;

            for &brick_j in bricks[brick_i].supports.iter() {
                if bricks[brick_j]
                    .supported_by
                    .iter()
                    .all(|&brick_k| moved[brick_k])
                {
                    queue.push_back(brick_j);
                    res += 1;
                }
            }
        }

        res
    }
}
