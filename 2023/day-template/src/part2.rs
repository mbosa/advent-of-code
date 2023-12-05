pub fn part2(input: &Vec<&str>) -> u32 {
    todo!();
}

#[cfg(test)]
mod test {
    use crate::parse_input;

    use super::*;

    #[test]
    fn test_part2() {
        let input = "";

        let parsed = parse_input(&input);

        let res = part2(&parsed);

        assert_eq!(res, 1);
    }
}
