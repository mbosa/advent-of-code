use crate::calc_hash;

pub fn part1(input: &str) -> usize {
    input.split(",").map(calc_hash).sum()
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7";

        let res = part1(input);

        assert_eq!(res, 1320);
    }
}
