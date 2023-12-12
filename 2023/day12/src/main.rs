mod part1;
mod part2;

use std::collections::HashMap;

use part1::part1;
use part2::part2;

type Input = Vec<Line>;
type Cache = HashMap<(Vec<char>, Vec<usize>), usize>;

#[derive(Debug)]
struct Line {
    s: Vec<char>,
    n: Vec<usize>,
}

fn main() {
    let input = include_str!("../../inputs/day12.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Input {
    input
        .lines()
        .map(|line| {
            let (s, n) = line.split_once(" ").unwrap();
            let s = s.chars().collect::<Vec<_>>();
            let n = n
                .split(",")
                .map(|n| n.parse::<usize>().unwrap())
                .collect::<Vec<_>>();

            Line { s, n }
        })
        .collect::<Vec<_>>()
}

fn count(s: Vec<char>, n: Vec<usize>, cache: &mut Cache) -> usize {
    if let Some(v) = cache.get(&(s.clone(), n.clone())) {
        println!("cache hit");
        return *v;
    }

    if s.len() == 0 {
        if n.len() == 0 {
            cache.insert((s, n), 1);
            return 1;
        } else {
            cache.insert((s, n), 0);
            return 0;
        }
    }

    if n.len() == 0 {
        if s.iter().any(|&c| c == '#') {
            cache.insert((s, n), 0);

            return 0;
        } else {
            cache.insert((s, n), 1);

            return 1;
        }
    }

    if s.len() < n.iter().sum::<usize>() + n.len() - 1 {
        cache.insert((s, n), 0);

        return 0;
    }

    if s[0] == '.' {
        let res = count(s[1..].to_vec(), n.clone(), cache);
        cache.insert((s, n), res);

        return res;
    }

    if s[0] == '#' {
        if s[..n[0]].iter().any(|&c| c == '.') {
            cache.insert((s, n), 0);

            return 0;
        }
        if s.len() > n[0] && s[n[0]] == '#' {
            cache.insert((s, n), 0);

            return 0;
        }

        let next_start = usize::min(s.len(), n[0] + 1);

        let res = count(s[next_start..].to_vec(), n[1..].to_vec(), cache);
        cache.insert((s, n), res);
        return res;
    }

    let mut next_s_dot = s.clone();
    next_s_dot[0] = '.';
    let mut next_s_pound = s.clone();
    next_s_pound[0] = '#';

    return count(next_s_dot, n.clone(), cache) + count(next_s_pound, n, cache);
}
