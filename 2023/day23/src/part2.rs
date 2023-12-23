use crate::{parse_input, Item};

pub fn part2(input: &str) -> u64 {
    let graph = parse_input(input, false);

    let mut res = 0;

    let start_vertex = 0;
    let end_vertex = graph.edges.len() - 1;

    let mut stack = Vec::new();

    let start_item = Item {
        steps: 0,
        vertex: start_vertex,
        prev_vertices: 0,
    };
    stack.push(start_item);

    while let Some(item) = stack.pop() {
        if item.is_vertex_seen(item.vertex) {
            continue;
        }

        if item.vertex == end_vertex {
            if item.steps > res {
                res = item.steps;
            }
            continue;
        }

        for edge in &graph.edges[item.vertex] {
            let mut next_item = Item {
                steps: item.steps + edge.weight as usize,
                vertex: edge.to,
                prev_vertices: item.prev_vertices,
            };
            next_item.add_vertex_to_seen(item.vertex);
            stack.push(next_item);
        }
    }

    res as u64
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = "#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#";

        let res = part2(input);

        assert_eq!(res, 154);
    }
}
