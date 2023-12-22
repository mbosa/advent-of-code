use std::collections::VecDeque;

use crate::{parse_input, Pulse, PulseType};

pub fn part1(input: &str) -> usize {
    let mut network = parse_input(input);

    let mut low_pulse_count = 0;
    let mut high_pulse_count = 0;
    let mut q = VecDeque::new();

    for _ in 0..1000 {
        let start_pulse = Pulse {
            to: "broadcaster",
            from: "button",
            pulse_type: PulseType::Low,
        };

        q.push_back(start_pulse);

        while let Some(p) = q.pop_front() {
            match p.pulse_type {
                PulseType::Low => low_pulse_count += 1,
                PulseType::High => high_pulse_count += 1,
            }

            if let Some(dst_module) = network.modules.get_mut(p.to) {
                let output = dst_module.receive_pulse(&p);

                if let Some(output) = output {
                    q.extend(output.into_iter());
                }
            }
        }
    }

    low_pulse_count * high_pulse_count
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output";

        let res = part1(input);

        assert_eq!(res, 11687500);
    }
}
