mod part1;
mod part2;

use part1::part1;
use part2::part2;

#[derive(Debug, Hash)]
struct Input<'a> {
    workflows: Vec<Workflow<'a>>,
    parts: Vec<Part>,
}

#[derive(Debug, Hash)]
struct Workflow<'a> {
    id: &'a str,
    rules: Vec<Rule<'a>>,
}

#[derive(Debug, Hash)]
enum Rule<'a> {
    ConditionRule {
        category: Category,
        symbol: Symbol,
        value: u64,
        dst: &'a str,
    },
    DefaultRule {
        dst: &'a str,
    },
}

#[derive(Debug, Hash)]
enum Category {
    X,
    M,
    A,
    S,
}

#[derive(Debug, Hash)]
enum Symbol {
    Gt,
    Lt,
}

#[derive(Debug, Hash)]
struct ConditionRule<'a> {
    category: char,
    symbol: char,
    value: u64,
    dst: &'a str,
}

#[derive(Debug, Hash)]
struct DefaultRule<'a> {
    dst: &'a str,
}

#[derive(Debug, Hash)]
struct Part {
    x: u64,
    m: u64,
    a: u64,
    s: u64,
}

fn main() {
    let input = include_str!("../../inputs/day19.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Input {
    let (workflows, parts) = input.split_once("\n\n").unwrap();

    let workflows = workflows
        .lines()
        .map(|line| {
            let (id, rest) = line.split_once("{").unwrap();

            let rules = rest[..rest.len() - 1]
                .split(",")
                .map(parse_rule)
                .collect::<Vec<_>>();

            Workflow { id, rules }
        })
        .collect::<Vec<_>>();

    let parts = parts
        .lines()
        .map(|line| {
            let mut s = line[1..line.len() - 1].split(",");

            let (_x, x_val) = s.next().unwrap().split_once("=").unwrap();
            let (_m, m_val) = s.next().unwrap().split_once("=").unwrap();
            let (_a, a_val) = s.next().unwrap().split_once("=").unwrap();
            let (_s, s_val) = s.next().unwrap().split_once("=").unwrap();

            Part {
                x: x_val.parse().unwrap(),
                m: m_val.parse().unwrap(),
                a: a_val.parse().unwrap(),
                s: s_val.parse().unwrap(),
            }
        })
        .collect::<Vec<_>>();

    Input { workflows, parts }
}

fn parse_rule(rule: &str) -> Rule {
    if !rule.contains(':') {
        return Rule::DefaultRule { dst: rule };
    }

    let mut c = rule.chars();
    let category = c
        .next()
        .map(|c| match c {
            'x' => Category::X,
            'm' => Category::M,
            'a' => Category::A,
            's' => Category::S,
            _ => unreachable!(),
        })
        .unwrap();
    let symbol = c
        .next()
        .map(|s| match s {
            '>' => Symbol::Gt,
            '<' => Symbol::Lt,
            _ => unreachable!(),
        })
        .unwrap();
    let rest = &rule[2..];
    let (value, dst) = rest
        .split_once(":")
        .map(|(v, d)| (v.parse::<u64>().unwrap(), d))
        .unwrap();
    let dst = dst;

    Rule::ConditionRule {
        category,
        symbol,
        value,
        dst,
    }
}
