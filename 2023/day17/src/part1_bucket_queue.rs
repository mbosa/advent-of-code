use crate::{find_min_path_bucket_queue, parse_input};

pub fn part1(input: &str) -> usize {
    let input = parse_input(input);

    find_min_path_bucket_queue(input, 1, 3)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
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

        let res = part1(input);

        assert_eq!(res, 102);
    }
}
