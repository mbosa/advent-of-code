use std::{collections::HashMap, time::Instant};

use crate::{count, parse_input, Cache};

pub fn part1(input: &str) -> usize {
    let now = Instant::now();

    let input = parse_input(input);

    let mut cache: Cache = HashMap::new();

    let mut res = 0;

    for line in input.iter() {
        res += count(line.s.clone(), line.n.clone(), &mut cache);
    }

    println!("time part1: {:?}", now.elapsed());

    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1";

        let res = part1(input);

        assert_eq!(res, 21);
    }
}
