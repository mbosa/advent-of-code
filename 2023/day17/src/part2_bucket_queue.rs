use crate::{find_min_path_bucket_queue, parse_input};

pub fn part2(input: &str) -> usize {
    let input = parse_input(input);

    find_min_path_bucket_queue(input, 4, 10)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2_1() {
        let input = "2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533";

        let res = part2(input);

        assert_eq!(res, 94);
    }
    #[test]
    fn test_part2_2() {
        let input = "111111111111
999999999991
999999999991
999999999991
999999999991";

        let res = part2(input);

        assert_eq!(res, 71);
    }
}
