mod part1;
mod part2;

use std::time::Instant;

use part1::part1;
use part2::part2;

#[derive(Debug, Clone, Eq, PartialEq, Hash)]
struct Graph {
    size: usize,
    edges: Vec<Vec<Edge>>,
}

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct Edge {
    to: usize,
    position: Position,
    weight: u64,
}

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct Position {
    row: usize,
    col: usize,
}

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct StackItem {
    position: Position,
    distance: u64,
}

struct Grid<'a> {
    grid: Vec<&'a [u8]>,
}
impl<'a> Grid<'a> {
    fn neighbors(&self, position: Position) -> Vec<Position> {
        [[-1, 0], [1, 0], [0, -1], [0, 1]]
            .into_iter()
            .flat_map(|[dx, dy]| {
                let new_row = position.row as i32 + dx;
                let new_col = position.col as i32 + dy;

                if new_row < 0 || new_row >= self.rows() as i32 {
                    return None;
                }
                if new_col < 0 || new_col >= self.cols() as i32 {
                    return None;
                }

                let p = Position {
                    row: new_row as usize,
                    col: new_col as usize,
                };
                Some(p)
            })
            .filter(|position| self.grid[position.row][position.col] != b'#')
            .collect::<Vec<_>>()
    }
    fn rows(&self) -> usize {
        self.grid.len()
    }
    fn cols(&self) -> usize {
        self.grid[0].len()
    }
}

#[derive(Debug, Clone, Eq, PartialEq)]
struct Item {
    steps: usize,
    vertex: usize,
    prev_vertices: u64,
}

impl Item {
    fn is_vertex_seen(&self, vertex: usize) -> bool {
        self.prev_vertices & (1 << vertex) != 0
    }
    fn add_vertex_to_seen(&mut self, vertex: usize) {
        self.prev_vertices |= 1 << vertex;
    }
}

fn main() {
    let input = include_str!("../../inputs/day23.txt");

    let start = Instant::now();
    let part1 = part1(input);
    println!("part1 time: {:?}", start.elapsed());
    println!("part1: {}", part1);

    let start = Instant::now();
    let part2 = part2(input);
    println!("part2 time: {:?}", start.elapsed());
    println!("part2: {}", part2);
}

fn parse_input(input: &str, directed: bool) -> Graph {
    let grid = input
        .lines()
        .map(|line| line.as_bytes())
        .collect::<Vec<_>>();
    let grid = Grid { grid };

    let mut visited = vec![vec![false; grid.cols()]; grid.rows()];
    let mut vertices = vec![vec![usize::MAX; grid.cols()]; grid.rows()];

    let start_position = Position { row: 0, col: 1 };
    let end_position = Position {
        row: grid.rows() - 1,
        col: grid.cols() - 2,
    };

    let mut size = 0;
    let mut edges: Vec<Vec<Edge>> = Vec::new();

    let mut stack = Vec::new();
    let mut stack_vertices = Vec::new();

    stack.push(start_position);

    while let Some(pos) = stack.pop() {
        if visited[pos.row][pos.col] {
            continue;
        } else {
            visited[pos.row][pos.col] = true;
        }

        let neighbors = grid.neighbors(pos);
        if neighbors.len() > 2 || pos == start_position || pos == end_position {
            let vertex_id = size;
            size += 1;
            vertices[pos.row][pos.col] = vertex_id;
            edges.push(Vec::new());

            stack_vertices.push(pos)
        }
        for neighbor in neighbors {
            stack.push(neighbor);
        }
    }

    while let Some(pos) = stack_vertices.pop() {
        let from_id = vertices[pos.row][pos.col];

        let mut visited = vec![vec![false; grid.cols()]; grid.rows()];

        let mut stack = Vec::new();
        let start_item = StackItem {
            position: pos,
            distance: 0,
        };
        stack.push(start_item);

        while let Some(item) = stack.pop() {
            if visited[item.position.row][item.position.col] {
                continue;
            } else {
                visited[item.position.row][item.position.col] = true;
            }

            let neighbors = grid.neighbors(item.position);
            let vertex_id = vertices[item.position.row][item.position.col];
            if (neighbors.len() > 2
                || item.position == start_position
                || item.position == end_position)
                && vertex_id != from_id
            {
                let edge = Edge {
                    to: vertex_id,
                    position: item.position,
                    weight: item.distance,
                };

                edges[from_id].push(edge);
            } else if directed {
                match grid.grid[item.position.row][item.position.col] {
                    b'^' => {
                        let next_item = StackItem {
                            position: Position {
                                row: item.position.row - 1,
                                col: item.position.col,
                            },
                            distance: item.distance + 1,
                        };
                        stack.push(next_item);
                    }
                    b'>' => {
                        let next_item = StackItem {
                            position: Position {
                                row: item.position.row,
                                col: item.position.col + 1,
                            },
                            distance: item.distance + 1,
                        };
                        stack.push(next_item);
                    }
                    b'v' => {
                        let next_item = StackItem {
                            position: Position {
                                row: item.position.row + 1,
                                col: item.position.col,
                            },
                            distance: item.distance + 1,
                        };
                        stack.push(next_item);
                    }
                    b'<' => {
                        let next_item = StackItem {
                            position: Position {
                                row: item.position.row,
                                col: item.position.col - 1,
                            },
                            distance: item.distance + 1,
                        };
                        stack.push(next_item);
                    }
                    _ => {
                        for neighbor in neighbors {
                            let next_item = StackItem {
                                position: neighbor,
                                distance: item.distance + 1,
                            };
                            stack.push(next_item);
                        }
                    }
                }
            } else {
                for neighbor in neighbors {
                    let next_item = StackItem {
                        position: neighbor,
                        distance: item.distance + 1,
                    };
                    stack.push(next_item);
                }
            }
        }
    }

    let graph = Graph { size, edges };
    // dbg!(&graph);
    graph
}
