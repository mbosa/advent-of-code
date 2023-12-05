pub fn part1(input: &Vec<&str>) -> u32 {
    todo!();
}

#[cfg(test)]
mod test {
    use crate::parse_input;

    use super::*;

    #[test]
    fn test_part1() {
        let input = "";

        let parsed = parse_input(&input);

        let res = part1(&parsed);

        assert_eq!(res, 1);
    }
}
