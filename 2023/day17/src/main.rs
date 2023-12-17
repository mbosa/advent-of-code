mod part1;
mod part1_bucket_queue;
mod part2;
mod part2_bucket_queue;

use std::collections::{BinaryHeap, HashSet};

use part1::part1;
use part1_bucket_queue::part1 as part1_bucket_queue;
use part2::part2;
use part2_bucket_queue::part2 as part2_bucket_queue;

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
enum Direction {
    Up,
    Down,
    Right,
    Left,
}

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct Position {
    row: usize,
    col: usize,
}

#[derive(Debug, Clone, Eq, PartialEq, Hash)]
struct QueueItem {
    cost: usize,
    position: Position,
    direction: Direction,
    steps_same_direction: usize,
}

impl Ord for QueueItem {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        if self.cost < other.cost {
            return std::cmp::Ordering::Greater;
        } else {
            return std::cmp::Ordering::Less;
        }
    }
}

impl PartialOrd for QueueItem {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}

#[derive(Debug, Clone, Eq, PartialEq)]
struct BucketQueue<T> {
    buckets: Vec<Vec<T>>,
    first: Option<usize>,
}

impl<T> BucketQueue<T> {
    fn new() -> Self {
        BucketQueue {
            buckets: Vec::new(),
            first: None,
        }
    }

    fn push(&mut self, cost: usize, el: T) {
        if cost >= self.buckets.len() {
            self.buckets.resize_with(cost + 1, || Vec::new())
        }
        self.buckets[cost].push(el);

        match self.first {
            Some(f) => self.first = Some(usize::min(f, cost)),
            None => self.first = Some(cost),
        }
    }

    fn pop(&mut self) -> Option<T> {
        let Some(first) = self.first else {
            return None;
        };

        let ret = self.buckets[first].pop();

        if self.buckets[first].is_empty() {
            self.first = self.buckets.iter().position(|bucket| !bucket.is_empty());
        }

        ret
    }
}

fn main() {
    let input = include_str!("../../inputs/day17.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part1_bucket_queue = part1_bucket_queue(input);
    println!("part1_bucket_queue: {}", part1_bucket_queue);

    let part2 = part2(input);
    println!("part2: {}", part2);

    let part2_bucket_queue = part2_bucket_queue(input);
    println!("part2_bucket_queue: {}", part2_bucket_queue);
}

fn parse_input(input: &str) -> Vec<Vec<usize>> {
    input
        .lines()
        .map(|line| {
            line.chars()
                .map(|el| el.to_digit(10).unwrap() as usize)
                .collect()
        })
        .collect()
}

fn find_min_path(
    input: Vec<Vec<usize>>,
    min_steps_same_direction: usize,
    max_steps_same_direction: usize,
) -> usize {
    let mut seen = HashSet::new();
    let mut priority_queue: BinaryHeap<QueueItem> = BinaryHeap::new();

    let start_right = QueueItem {
        cost: 0,
        position: Position { row: 0, col: 0 },
        direction: Direction::Right,
        steps_same_direction: 0,
    };
    let start_down = QueueItem {
        cost: 0,
        position: Position { row: 0, col: 0 },
        direction: Direction::Down,
        steps_same_direction: 0,
    };

    priority_queue.push(start_right);
    priority_queue.push(start_down);

    while let Some(item) = priority_queue.pop() {
        let already_seen = seen.insert((
            item.position.row,
            item.position.col,
            item.direction,
            item.steps_same_direction,
        ));
        if !already_seen {
            continue;
        }

        if item.position.row == input.len() - 1 && item.position.col == input[0].len() - 1 {
            if item.steps_same_direction < min_steps_same_direction {
                continue;
            } else {
                return item.cost;
            }
        }

        if item.steps_same_direction < min_steps_same_direction {
            let [dx, dy] = match item.direction {
                Direction::Up => [-1, 0],
                Direction::Down => [1, 0],
                Direction::Right => [0, 1],
                Direction::Left => [0, -1],
            };

            // check bounds
            if (item.position.row == 0 && item.direction == Direction::Up)
                || (item.position.row == input.len() - 1 && item.direction == Direction::Down)
                || (item.position.col == 0 && item.direction == Direction::Left)
                || (item.position.col == input[0].len() - 1 && item.direction == Direction::Right)
            {
                continue;
            }

            let new_row = (item.position.row as i32 + dx) as usize;
            let new_col = (item.position.col as i32 + dy) as usize;

            let new_item = QueueItem {
                cost: item.cost + input[new_row][new_col],
                position: Position {
                    row: new_row,
                    col: new_col,
                },
                direction: item.direction,
                steps_same_direction: item.steps_same_direction + 1,
            };
            priority_queue.push(new_item);

            continue;
        }

        for (direction, [dx, dy]) in [
            (Direction::Up, [-1, 0]),
            (Direction::Down, [1, 0]),
            (Direction::Left, [0, -1]),
            (Direction::Right, [0, 1]),
        ] {
            // cannot reverse direction
            match [item.direction, direction] {
                [Direction::Up, Direction::Down]
                | [Direction::Down, Direction::Up]
                | [Direction::Left, Direction::Right]
                | [Direction::Right, Direction::Left] => continue,
                _ => {}
            }

            // check bounds
            if (item.position.row == 0 && direction == Direction::Up)
                || (item.position.row == input.len() - 1 && direction == Direction::Down)
                || (item.position.col == 0 && direction == Direction::Left)
                || (item.position.col == input[0].len() - 1 && direction == Direction::Right)
            {
                continue;
            }

            // max steps in the same direction
            if item.direction == direction && item.steps_same_direction == max_steps_same_direction
            {
                continue;
            }

            let new_row = (item.position.row as i32 + dx) as usize;
            let new_col = (item.position.col as i32 + dy) as usize;

            let new_item = QueueItem {
                cost: item.cost + input[new_row][new_col],
                position: Position {
                    row: new_row,
                    col: new_col,
                },
                direction,
                steps_same_direction: if item.direction == direction {
                    item.steps_same_direction + 1
                } else {
                    1
                },
            };
            priority_queue.push(new_item);
        }
    }

    panic!("Solution not found")
}

fn find_min_path_bucket_queue(
    input: Vec<Vec<usize>>,
    min_steps_same_direction: usize,
    max_steps_same_direction: usize,
) -> usize {
    let mut seen = HashSet::new();
    let mut priority_queue = BucketQueue::new();

    let start_right = QueueItem {
        cost: 0,
        position: Position { row: 0, col: 0 },
        direction: Direction::Right,
        steps_same_direction: 0,
    };
    let start_down = QueueItem {
        cost: 0,
        position: Position { row: 0, col: 0 },
        direction: Direction::Down,
        steps_same_direction: 0,
    };

    priority_queue.push(start_right.cost, start_right);
    priority_queue.push(start_down.cost, start_down);

    while let Some(item) = priority_queue.pop() {
        let already_seen = seen.insert((
            item.position.row,
            item.position.col,
            item.direction,
            item.steps_same_direction,
        ));
        if !already_seen {
            continue;
        }

        if item.position.row == input.len() - 1 && item.position.col == input[0].len() - 1 {
            if item.steps_same_direction < min_steps_same_direction {
                continue;
            } else {
                return item.cost;
            }
        }

        if item.steps_same_direction < min_steps_same_direction {
            let [dx, dy] = match item.direction {
                Direction::Up => [-1, 0],
                Direction::Down => [1, 0],
                Direction::Right => [0, 1],
                Direction::Left => [0, -1],
            };

            // check bounds
            if (item.position.row == 0 && item.direction == Direction::Up)
                || (item.position.row == input.len() - 1 && item.direction == Direction::Down)
                || (item.position.col == 0 && item.direction == Direction::Left)
                || (item.position.col == input[0].len() - 1 && item.direction == Direction::Right)
            {
                continue;
            }

            let new_row = (item.position.row as i32 + dx) as usize;
            let new_col = (item.position.col as i32 + dy) as usize;

            let new_item = QueueItem {
                cost: item.cost + input[new_row][new_col],
                position: Position {
                    row: new_row,
                    col: new_col,
                },
                direction: item.direction,
                steps_same_direction: item.steps_same_direction + 1,
            };
            priority_queue.push(new_item.cost, new_item);

            continue;
        }

        for (direction, [dx, dy]) in [
            (Direction::Up, [-1, 0]),
            (Direction::Down, [1, 0]),
            (Direction::Left, [0, -1]),
            (Direction::Right, [0, 1]),
        ] {
            // cannot reverse direction
            match [item.direction, direction] {
                [Direction::Up, Direction::Down]
                | [Direction::Down, Direction::Up]
                | [Direction::Left, Direction::Right]
                | [Direction::Right, Direction::Left] => continue,
                _ => {}
            }

            // check bounds
            if (item.position.row == 0 && direction == Direction::Up)
                || (item.position.row == input.len() - 1 && direction == Direction::Down)
                || (item.position.col == 0 && direction == Direction::Left)
                || (item.position.col == input[0].len() - 1 && direction == Direction::Right)
            {
                continue;
            }

            // max steps in the same direction
            if item.direction == direction && item.steps_same_direction == max_steps_same_direction
            {
                continue;
            }

            let new_row = (item.position.row as i32 + dx) as usize;
            let new_col = (item.position.col as i32 + dy) as usize;

            let new_item = QueueItem {
                cost: item.cost + input[new_row][new_col],
                position: Position {
                    row: new_row,
                    col: new_col,
                },
                direction,
                steps_same_direction: if item.direction == direction {
                    item.steps_same_direction + 1
                } else {
                    1
                },
            };
            priority_queue.push(new_item.cost, new_item);
        }
    }

    panic!("Solution not found")
}
