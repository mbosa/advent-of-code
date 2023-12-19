use crate::{parse_input, Category, Part, Rule, Symbol, Workflow};

pub fn part1(input: &str) -> u64 {
    let input = parse_input(input);

    let mut res = 0;

    for part in input.parts {
        let mut workflow = "in";

        while workflow != "A" && workflow != "R" {
            let w = input.workflows.iter().find(|w| w.id == workflow).unwrap();

            workflow = calc_workflow(&part, &w);
        }

        if workflow == "A" {
            let sum = part.x + part.m + part.s + part.a;
            res += sum;
        }
    }

    res
}

fn calc_workflow<'a>(part: &'a Part, workflow: &'a Workflow) -> &'a str {
    for rule in workflow.rules.iter() {
        match rule {
            Rule::DefaultRule { dst } => return dst,
            Rule::ConditionRule {
                category,
                symbol,
                value,
                dst,
            } => {
                let cmp_val = match category {
                    Category::X => part.x,
                    Category::M => part.m,
                    Category::A => part.a,
                    Category::S => part.s,
                };
                match symbol {
                    Symbol::Gt => {
                        if cmp_val > *value {
                            return dst;
                        }
                    }
                    Symbol::Lt => {
                        if cmp_val < *value {
                            return dst;
                        }
                    }
                }
            }
        }
    }
    unreachable!()
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}";

        let res = part1(input);

        assert_eq!(res, 19114);
    }
}
