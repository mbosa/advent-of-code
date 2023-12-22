use std::collections::{HashMap, VecDeque};

use crate::{parse_input, Pulse, PulseType};

pub fn part2(input: &str) -> u64 {
    let mut network = parse_input(input);

    let mut q = VecDeque::new();

    let final_modules = ["cl", "rp", "lb", "nj"];
    let mut c = HashMap::new();

    let mut i = 0;
    'outer: loop {
        i += 1;

        let start_pulse = Pulse {
            to: "broadcaster",
            from: "button",
            pulse_type: PulseType::Low,
        };

        q.push_back(start_pulse);

        while let Some(p) = q.pop_front() {
            if final_modules.contains(&p.to) && p.pulse_type == PulseType::Low {
                c.entry(p.to).or_insert(i);
            }

            if c.len() == 4 {
                break 'outer;
            }

            if let Some(dst_module) = network.modules.get_mut(p.to) {
                let output = dst_module.receive_pulse(&p);

                if let Some(output) = output {
                    q.extend(output.into_iter());
                }
            }
        }
    }

    let res = c.into_values().reduce(|acc, el| lcm(acc, el)).unwrap();
    res
}

fn gcd(a: u64, b: u64) -> u64 {
    let mut a = a;
    let mut b = b;
    while a != b {
        if a > b {
            a = a - b;
        } else {
            b = b - a;
        }
    }

    a
}

fn lcm(a: u64, b: u64) -> u64 {
    (a * b) / gcd(a, b)
}
