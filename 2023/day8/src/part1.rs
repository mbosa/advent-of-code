use anyhow::Result;

use crate::parse_input;

pub fn part1(input: &str) -> Result<u32> {
    let input = parse_input(input)?;

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

    Ok(count)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)";

        let res = part1(input).unwrap();

        assert_eq!(res, 6);
    }
}
