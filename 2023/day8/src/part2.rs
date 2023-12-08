use crate::Input;

pub fn part2(input: &Input) -> u64 {
    let mut results: Vec<u64> = Vec::new();

    let starting_positions = input
        .nodes
        .keys()
        .filter(|key| key.ends_with("A"))
        .map(|&e| e)
        .collect::<Vec<_>>();

    for starting_position in starting_positions {
        let mut count = 0;
        let mut position = starting_position;

        for char in input.instructions.chars().cycle() {
            let node = input.nodes.get(position).unwrap();

            position = match char {
                'L' => node[0],
                'R' => node[1],
                _ => unreachable!(),
            };
            count += 1;

            if position.ends_with("Z") {
                break;
            }
        }
        results.push(count);
    }

    results.into_iter().reduce(|acc, el| lcm(acc, el)).unwrap()
}

fn gcd(a: u64, b: u64) -> u64 {
    let mut a = a;
    let mut b = b;
    while a != b {
        if a > b {
            a = a - b;
        } else {
            b = b - a;
        }
    }

    a
}

fn lcm(a: u64, b: u64) -> u64 {
    (a * b) / gcd(a, b)
}

#[cfg(test)]
mod test {
    use crate::parse_input;

    use super::*;

    #[test]
    fn test_part2() {
        let input = "LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)";

        let parsed = parse_input(&input).unwrap();

        let res = part2(&parsed);

        assert_eq!(res, 6);
    }
}
