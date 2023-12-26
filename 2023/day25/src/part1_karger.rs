use rand::Rng;
use std::collections::{HashMap, HashSet};

type Edge<'a> = (&'a str, &'a str);

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
struct Cluster<'a> {
    id: &'a str,
    count: u64,
}

#[derive(Debug, Clone, PartialEq, Eq)]
struct Graph<'a> {
    nodes_count: usize,
    edges: Vec<Edge<'a>>,
    clusters: HashMap<&'a str, u64>,
}

impl Graph<'_> {
    fn remove_connection_random(&mut self) {
        let r = rand::thread_rng().gen_range(0..self.edges.len());
        let edge_to_remove = self.edges[r];

        for i in 0..self.edges.len() {
            let e = self.edges[i];

            if e.0 == edge_to_remove.1 {
                self.edges[i].0 = edge_to_remove.0;
            }
            if e.1 == edge_to_remove.1 {
                self.edges[i].1 = edge_to_remove.0;
            }
        }

        let deleted_cluster_count = self.clusters.remove(edge_to_remove.1).unwrap();
        self.clusters
            .entry(edge_to_remove.0)
            .and_modify(|v| *v += deleted_cluster_count);

        // remove edge with itself
        let mut i_to_remove = Vec::new();
        for i in 0..self.edges.len() {
            let e = self.edges[i];
            if e.0 == e.1 {
                i_to_remove.push(i)
            }
        }
        for &i in i_to_remove.iter().rev() {
            self.edges.remove(i);
        }
        self.nodes_count -= 1;
    }
}

fn parse_input(input: &str) -> Graph {
    let mut edges = Vec::new();
    let mut nodes_set = HashSet::new();
    let mut clusters = HashMap::new();

    for line in input.lines() {
        let (src, dsts) = line.split_once(": ").unwrap();
        nodes_set.insert(src);
        clusters.insert(src, 1);

        for dst in dsts.split_whitespace() {
            nodes_set.insert(dst);
            clusters.insert(dst, 1);
            edges.push((src, dst));
        }
    }

    Graph {
        nodes_count: nodes_set.len(),
        edges,
        clusters,
    }
}

/// Calculate the result using Karger's algorithm (https://en.wikipedia.org/wiki/Karger%27s_algorithm)
/// * Keep contracting a random edge until only 2 nodes remain.
/// * Repeat the process from scratch until the 2 remaining nodes have only 3 edges.
///
/// The algorithm is non-deterministic, it takes a variable amount of time to complete
pub fn part1(input: &str) -> u64 {
    let input = parse_input(input);

    loop {
        let mut graph = input.clone();
        while graph.nodes_count > 2 {
            graph.remove_connection_random();
        }
        if graph.edges.len() == 3 {
            return graph.clusters.iter().map(|(_, v)| v).product();
        }
    }
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
