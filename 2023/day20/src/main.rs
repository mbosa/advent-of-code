mod part1;
mod part2;

use std::collections::HashMap;

use part1::part1;
use part2::part2;

#[derive(Debug, Clone, Eq, PartialEq)]
enum Module<'a> {
    FlipFlop {
        id: &'a str,
        state: FlipFlopState,
        dst: Vec<&'a str>,
    },
    Conjunction {
        id: &'a str,
        prev_pulses: HashMap<&'a str, Pulse<'a>>,
        dst: Vec<&'a str>,
    },
    Broadcaster {
        id: &'a str,
        dst: Vec<&'a str>,
    },
}

#[derive(Debug, Clone, Eq, PartialEq, Hash)]
struct Pulse<'a> {
    to: &'a str,
    from: &'a str,
    pulse_type: PulseType,
}

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
enum PulseType {
    High,
    Low,
}

#[derive(Debug, Clone, Eq, PartialEq, Hash)]
enum FlipFlopState {
    On,
    Off,
}

#[derive(Debug, Clone, Eq, PartialEq)]
struct Network<'a> {
    modules: HashMap<&'a str, Module<'a>>,
}

impl<'a> Module<'a> {
    fn receive_pulse(&mut self, pulse: &Pulse) -> Option<Vec<Pulse<'a>>> {
        match self {
            Module::FlipFlop { id, state, dst } => {
                if pulse.pulse_type == PulseType::High {
                    return None;
                }

                match state {
                    FlipFlopState::Off => {
                        *state = FlipFlopState::On;

                        let pulses = dst
                            .iter()
                            .map(|d| Pulse {
                                from: id,
                                to: d,
                                pulse_type: PulseType::High,
                            })
                            .collect();

                        return Some(pulses);
                    }
                    FlipFlopState::On => {
                        *state = FlipFlopState::Off;

                        let pulses = dst
                            .iter()
                            .map(|d| Pulse {
                                from: id,
                                to: d,
                                pulse_type: PulseType::Low,
                            })
                            .collect();

                        return Some(pulses);
                    }
                }
            }
            Module::Conjunction {
                id,
                prev_pulses,
                dst,
            } => {
                let prev_pulse = prev_pulses.get_mut(pulse.from).unwrap();
                prev_pulse.pulse_type = pulse.pulse_type;

                if prev_pulses
                    .values()
                    .all(|pulse| pulse.pulse_type == PulseType::High)
                {
                    let pulses = dst
                        .iter()
                        .map(|d| Pulse {
                            from: id,
                            to: d,
                            pulse_type: PulseType::Low,
                        })
                        .collect();

                    return Some(pulses);
                } else {
                    let pulses = dst
                        .iter()
                        .map(|d| Pulse {
                            from: id,
                            to: d,
                            pulse_type: PulseType::High,
                        })
                        .collect();

                    return Some(pulses);
                }
            }
            Module::Broadcaster { id, dst } => {
                let pulses = dst
                    .iter()
                    .map(|d| Pulse {
                        from: id,
                        to: d,
                        pulse_type: pulse.pulse_type,
                    })
                    .collect();

                return Some(pulses);
            }
        }
    }
}

fn main() {
    let input = include_str!("../../inputs/day20.txt");

    let part1 = part1(input);
    println!("part1: {}", part1);

    let part2 = part2(input);
    println!("part2: {}", part2);
}

fn parse_input(input: &str) -> Network {
    let mut modules = HashMap::new();

    for l in input.lines() {
        let (src, dst) = l.split_once(" -> ").unwrap();
        let id = if src == "broadcaster" { src } else { &src[1..] };
        let dst = dst.split(", ").collect::<Vec<_>>();

        let module = match &src[0..1] {
            "%" => Module::FlipFlop {
                id,
                state: FlipFlopState::Off,
                dst,
            },
            "&" => Module::Conjunction {
                id,
                prev_pulses: HashMap::new(),
                dst,
            },
            "b" => Module::Broadcaster { id, dst },
            _ => unreachable!(),
        };

        modules.insert(id, module);
    }

    for (src_m, module) in modules.clone().iter() {
        match module {
            Module::FlipFlop {
                id: _,
                state: _,
                dst,
            } => {
                for dst_m in dst {
                    let dst_module = modules.get_mut(dst_m).unwrap();

                    if let Module::Conjunction {
                        id,
                        prev_pulses,
                        dst: _,
                    } = dst_module
                    {
                        prev_pulses.insert(
                            src_m,
                            Pulse {
                                from: id,
                                to: dst_m,
                                pulse_type: PulseType::Low,
                            },
                        );
                    }
                }
            }
            Module::Conjunction {
                id: _,
                prev_pulses: _,
                dst,
            } => {
                for dst_m in dst {
                    let dst_module = modules.get_mut(dst_m);
                    // .expect(&format!("{} not found", dst_m));

                    if let Some(dst_module) = dst_module {
                        if let Module::Conjunction {
                            id,
                            prev_pulses,
                            dst: _,
                        } = dst_module
                        {
                            prev_pulses.insert(
                                src_m,
                                Pulse {
                                    from: id,
                                    to: dst_m,
                                    pulse_type: PulseType::Low,
                                },
                            );
                        }
                    }
                }
            }
            Module::Broadcaster { id: _, dst } => {
                for dst_m in dst {
                    let dst_module = modules.get_mut(dst_m).unwrap();

                    if let Module::Conjunction {
                        id,
                        prev_pulses,
                        dst: _,
                    } = dst_module
                    {
                        prev_pulses.insert(
                            src_m,
                            Pulse {
                                from: id,
                                to: dst_m,
                                pulse_type: PulseType::Low,
                            },
                        );
                    }
                }
            }
        }
    }

    Network { modules }
}
