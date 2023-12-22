use crate::parse_input;

pub fn part1(input: &str) -> usize {
    let mut bricks = parse_input(input);
    bricks.drop_bricks();

    (0..bricks.len())
        .filter(|&i| bricks.count_chain_reaction(i) == 0)
        .count()
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1_1() {
        let input = "1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9";

        let res = part1(input);

        assert_eq!(res, 5);
    }
    #[test]
    fn test_part1_2() {
        let input = "0,0,1~0,1,1
1,1,1~1,1,1
0,0,2~0,0,2
0,1,2~1,1,2";

        let res = part1(input);

        assert_eq!(res, 3);
    }
    #[test]
    fn test_part1_3() {
        let input = "0,0,1~1,0,1
0,1,1~0,1,2
0,0,5~0,0,5
0,0,4~0,1,4";

        let res = part1(input);

        assert_eq!(res, 2);
    }
}
