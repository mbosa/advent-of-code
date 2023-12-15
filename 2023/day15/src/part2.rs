use crate::calc_hash;

#[derive(Debug, Clone, Copy)]
struct Lens<'a> {
    label: &'a str,
    focal_length: usize,
}

pub fn part2(input: &str) -> usize {
    let mut boxes: Vec<Vec<Lens<'_>>> = vec![Vec::new(); 256];

    for step in input.split(",") {
        let sign_idx = step.bytes().position(|b| b == b'=' || b == b'-').unwrap();

        let label = &step[..sign_idx];
        let box_idx = calc_hash(label);
        let sign = step.as_bytes()[sign_idx];

        match sign {
            b'=' => {
                let focal_length = step[sign_idx + 1..].parse::<usize>().unwrap();

                let lens = Lens {
                    label,
                    focal_length,
                };

                match boxes[box_idx].iter().position(|lens| lens.label == label) {
                    Some(i) => boxes[box_idx][i] = lens,
                    None => boxes[box_idx].push(lens),
                }
            }
            b'-' => {
                if let Some(i) = boxes[box_idx].iter().position(|lens| lens.label == label) {
                    boxes[box_idx].remove(i);
                }
            }
            _ => unreachable!(),
        }
    }

    boxes
        .iter()
        .enumerate()
        .flat_map(|(i, b)| {
            b.iter()
                .enumerate()
                .map(move |(j, lens)| (i + 1) * (j + 1) * lens.focal_length)
        })
        .sum()
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7";

        let res = part2(input);

        assert_eq!(res, 145);
    }
}
