use std::collections::{HashMap, HashSet, VecDeque};

type Graph<'a> = HashMap<&'a str, Vec<&'a str>>;

struct Item<'a> {
    id: &'a str,
    path: Vec<&'a str>,
}

fn parse_input_(input: &str) -> Graph {
    let mut graph = HashMap::new();

    for line in input.lines() {
        let (src, dsts) = line.split_once(": ").unwrap();
        let node_src = graph.entry(src).or_insert(Vec::new());

        for dst in dsts.split_whitespace() {
            node_src.push(dst);
        }

        for dst in dsts.split_whitespace() {
            let node_dst = graph.entry(dst).or_insert(Vec::new());
            node_dst.push(src);
        }
    }

    graph
}

pub fn part1(input: &str) -> usize {
    let graph = parse_input_(input);

    let ordered_nodes = get_ordered_nodes(&graph);
    let removed = calc_removed_edges(&graph, &ordered_nodes);

    let groups = count_groups(&graph, &removed);

    groups.iter().product()
}

/// Return all the nodes ordered closest to farthest from an arbitrary start node */
fn get_ordered_nodes<'a>(graph: &'a Graph) -> Vec<&'a str> {
    let mut q: VecDeque<&str> = VecDeque::new();
    let mut seen: HashSet<&str> = HashSet::new();
    let mut ordered_nodes = Vec::new();

    let start = graph.keys().next().unwrap();
    q.push_back(start);

    while let Some(node) = q.pop_front() {
        if !seen.insert(node) {
            continue;
        }
        ordered_nodes.push(node);

        let connections = graph.get(node).unwrap();
        q.extend(connections);
    }

    ordered_nodes
}

/// Calculate the 3 edges to remove using Girvan–Newman algorithm (https://en.wikipedia.org/wiki/Girvan%E2%80%93Newman_algorithm)
/// * calculate the shortest path between pairs of nodes and keep track of the frequency of each edge
/// * remove the edge with the highest frequency
/// * repeat until all 3 edges are found
///
/// The pairs are made using `ordered_nodes` so that the nodes of the pair are far from each other
fn calc_removed_edges<'a>(graph: &'a Graph, ordered_nodes: &'a Vec<&'a str>) -> [[&'a str; 2]; 3] {
    let mut removed = [[""; 2]; 3];

    for i in 0..3 {
        let mut freq: HashMap<[&str; 2], usize> = HashMap::new();

        'loop_pairs: for j in 0..usize::min(graph.len(), 100) {
            let start = ordered_nodes[j];
            let end = ordered_nodes[ordered_nodes.len() - 1 - j];

            let mut q = VecDeque::new();
            let mut seen = HashSet::new();
            let start_item = Item {
                id: start,
                path: Vec::new(),
            };
            q.push_back(start_item);

            while let Some(mut item) = q.pop_front() {
                if !seen.insert(item.id) {
                    continue;
                }
                item.path.push(item.id);

                if item.id == end {
                    for i in 1..item.path.len() {
                        let a = item.path[i - 1];
                        let b = item.path[i];
                        let key = if a < b { [a, b] } else { [b, a] };
                        *freq.entry(key).or_insert(0) += 1;
                    }

                    continue 'loop_pairs;
                }

                let connections = graph.get(item.id).unwrap();
                for &c in connections {
                    let key = if item.id < c {
                        [item.id, c]
                    } else {
                        [c, item.id]
                    };
                    if removed.contains(&key) {
                        continue;
                    }

                    let item = Item {
                        id: c,
                        path: item.path.clone(),
                    };
                    q.push_back(item);
                }
            }
        }

        let max = freq.iter().max_by_key(|(_, &n)| n).unwrap();

        removed[i] = *max.0;
    }

    removed
}

fn count_groups<'a>(graph: &'a Graph, removed: &[[&str; 2]; 3]) -> [usize; 2] {
    let mut q: VecDeque<&str> = VecDeque::new();
    let mut seen: HashSet<&str> = HashSet::new();
    let &start = graph.keys().next().unwrap();

    q.push_back(start);

    while let Some(node) = q.pop_front() {
        if !seen.insert(node) {
            continue;
        }

        let connections = graph.get(node).unwrap();
        for &c in connections {
            let key = if node < c { [node, c] } else { [c, node] };
            if removed.contains(&key) {
                continue;
            }

            q.push_back(c);
        }
    }

    [seen.len(), graph.len() - seen.len()]
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr";

        let res = part1(input);

        assert_eq!(res, 54);
    }
}
