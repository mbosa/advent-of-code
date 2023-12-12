use std::{collections::HashMap, time::Instant};

use crate::{count, parse_input, Cache};

pub fn part2(input: &str) -> usize {
    let now = Instant::now();

    let input = parse_input(input);

    let mut cache: Cache = HashMap::new();

    let mut res = 0;

    for line in input.iter() {
        let mut new_s = Vec::new();
        let mut new_n = Vec::new();

        for _ in 0..5 {
            new_s.push(line.s.clone());
            new_n.push(line.n.clone());
        }

        let new_s = new_s.join(&'?');
        let new_n = new_n.iter().cloned().flatten().collect::<Vec<_>>();

        res += count(new_s, new_n, &mut cache);
    }

    println!("time part2: {:?}", now.elapsed());

    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = "???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1";

        let res = part2(input);

        assert_eq!(res, 525152);
    }
}
