use crate::Input;

pub fn part1(input: &Input) -> u32 {
    let mut count = 0;
    let mut position = "AAA";

    for char in input.instructions.chars().cycle() {
        let node = input.nodes.get(position).unwrap();

        position = match char {
            'L' => node[0],
            'R' => node[1],
            _ => unreachable!(),
        };
        count += 1;

        if position == "ZZZ" {
            break;
        }
    }

    count
}

#[cfg(test)]
mod test {
    use crate::parse_input;

    use super::*;

    #[test]
    fn test_part1() {
        let input = "LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)";

        let parsed = parse_input(&input).unwrap();

        let res = part1(&parsed);

        assert_eq!(res, 6);
    }
}
