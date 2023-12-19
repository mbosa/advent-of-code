use crate::{parse_input, Category, Rule, Symbol};

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct StackItem<'a> {
    workflow_id: &'a str,
    x_range: Range,
    m_range: Range,
    a_range: Range,
    s_range: Range,
}

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct Range(u64, u64);
impl Range {
    /// Split Range at `value`.
    /// `value` is excluded from `left` and included in `right`
    fn split(&self, value: u64) -> Option<[Self; 2]> {
        if self.1 - self.0 < 2 {
            return None;
        }

        if value < self.0 || value >= self.1 {
            return None;
        }

        let left = Range(self.0, value);
        let right = Range(value, self.1);

        Some([left, right])
    }
    fn is_empty(&self) -> bool {
        self.0 == self.1
    }
}

pub fn part2(input: &str) -> u64 {
    let input = parse_input(input);

    let mut res = 0;
    let mut stack: Vec<StackItem> = Vec::new();
    let start_item = StackItem {
        workflow_id: "in",
        x_range: Range(1, 4001),
        m_range: Range(1, 4001),
        a_range: Range(1, 4001),
        s_range: Range(1, 4001),
    };
    stack.push(start_item);

    while let Some(item) = stack.pop() {
        if item.workflow_id == "R" {
            continue;
        }
        if item.workflow_id == "A" {
            let x = item.x_range.1 - item.x_range.0;
            let m = item.m_range.1 - item.m_range.0;
            let a = item.a_range.1 - item.a_range.0;
            let s = item.s_range.1 - item.s_range.0;

            res += x * m * a * s;

            continue;
        }

        let workflow = input
            .workflows
            .iter()
            .find(|w| w.id == item.workflow_id)
            .unwrap();

        for rule in workflow.rules.iter() {
            match rule {
                Rule::DefaultRule { dst } => {
                    let new_item = StackItem {
                        workflow_id: dst,
                        x_range: item.x_range,
                        m_range: item.m_range,
                        a_range: item.a_range,
                        s_range: item.s_range,
                    };

                    stack.push(new_item);

                    break;
                }
                Rule::ConditionRule {
                    category,
                    symbol,
                    value,
                    dst,
                } => {
                    let category_range = match category {
                        Category::X => item.x_range,
                        Category::M => item.m_range,
                        Category::A => item.a_range,
                        Category::S => item.s_range,
                    };

                    let ranges = match symbol {
                        Symbol::Gt => category_range.split(*value + 1),
                        Symbol::Lt => category_range
                            .split(*value)
                            .map(|[left, right]| [right, left]),
                    };

                    if ranges.is_none() {
                        continue;
                    }

                    let [range_to_src_workflow, range_to_dst_workflow] = ranges.unwrap();

                    if range_to_dst_workflow.is_empty() {
                        continue;
                    }

                    let mut item_to_src = StackItem {
                        workflow_id: item.workflow_id,
                        x_range: item.x_range,
                        m_range: item.m_range,
                        a_range: item.a_range,
                        s_range: item.s_range,
                    };
                    match category {
                        Category::X => item_to_src.x_range = range_to_src_workflow,
                        Category::M => item_to_src.m_range = range_to_src_workflow,
                        Category::A => item_to_src.a_range = range_to_src_workflow,
                        Category::S => item_to_src.s_range = range_to_src_workflow,
                    }

                    let mut item_to_dst = StackItem {
                        workflow_id: dst,
                        x_range: item.x_range,
                        m_range: item.m_range,
                        a_range: item.a_range,
                        s_range: item.s_range,
                    };
                    match category {
                        Category::X => item_to_dst.x_range = range_to_dst_workflow,
                        Category::M => item_to_dst.m_range = range_to_dst_workflow,
                        Category::A => item_to_dst.a_range = range_to_dst_workflow,
                        Category::S => item_to_dst.s_range = range_to_dst_workflow,
                    }
                    stack.push(item_to_src);
                    stack.push(item_to_dst);
                    break;
                }
            }
        }
    }

    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
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

        let res = part2(input);

        assert_eq!(res, 167409079868000);
    }
}
