use crate::{parse_input, Position, Velocity};

pub fn part2(input: &str) -> i64 {
    let input = parse_input(input);

    // https://www.reddit.com/r/adventofcode/comments/18pptor/comment/kepufsi/

    for vx in -300..300 {
        for vy in -300..300 {
            'loop_z: for vz in -300..300 {
                let mut hail0 = input[0].clone();
                let mut hail1 = input[1].clone();

                let new_velocity_0 = Velocity {
                    vx: hail0.velocity.vx - vx,
                    vy: hail0.velocity.vy - vy,
                    vz: hail0.velocity.vz - vz,
                };
                let new_velocity_1 = Velocity {
                    vx: hail1.velocity.vx - vx,
                    vy: hail1.velocity.vy - vy,
                    vz: hail1.velocity.vz - vz,
                };

                hail0.velocity = new_velocity_0;
                hail1.velocity = new_velocity_1;

                let mut t0: f64 = 0.0;
                let mut t1: f64 = 0.0;

                if hail0.velocity.vx == 0 {
                    t1 = (hail0.position.x - hail1.position.x) as f64 / hail1.velocity.vx as f64;
                } else if hail1.velocity.vx == 0 {
                    t0 = (hail1.position.x - hail0.position.x) as f64 / hail0.velocity.vx as f64;
                } else if hail0.velocity.vy == 0 {
                    t1 = (hail0.position.y - hail1.position.y) as f64 / hail1.velocity.vy as f64;
                } else if hail1.velocity.vy == 0 {
                    t0 = (hail1.position.y - hail0.position.y) as f64 / hail0.velocity.vy as f64;
                } else {
                    t1 = ((hail1.position.y - hail0.position.y) * hail0.velocity.vx
                        - (hail1.position.x - hail0.position.x) * hail0.velocity.vy)
                        as f64
                        / (hail1.velocity.vx * hail0.velocity.vy
                            - hail0.velocity.vx * hail1.velocity.vy)
                            as f64;
                }

                if t0 < 0.0 || t1 < 0.0 || t0 % 1.0 != 0.0 || t1 % 1.0 != 0.0 {
                    continue;
                }

                let intersection = if t0 != 0.0 {
                    Position {
                        x: hail0.position.x + hail0.velocity.vx * t0 as i64,
                        y: hail0.position.y + hail0.velocity.vy * t0 as i64,
                        z: hail0.position.z + hail0.velocity.vz * t0 as i64,
                    }
                } else {
                    Position {
                        x: hail1.position.x + hail1.velocity.vx * t1 as i64,
                        y: hail1.position.y + hail1.velocity.vy * t1 as i64,
                        z: hail1.position.z + hail1.velocity.vz * t1 as i64,
                    }
                };

                for h in input.iter() {
                    let mut h = h.clone();
                    let new_velocity = Velocity {
                        vx: h.velocity.vx - vx,
                        vy: h.velocity.vy - vy,
                        vz: h.velocity.vz - vz,
                    };
                    h.velocity = new_velocity;

                    let mut q = Vec::new();
                    if h.velocity.vx != 0 {
                        q.push((intersection.x - h.position.x) as f64 / h.velocity.vx as f64);
                    }
                    if h.velocity.vy != 0 {
                        q.push((intersection.y - h.position.y) as f64 / h.velocity.vy as f64);
                    }
                    if h.velocity.vz != 0 {
                        q.push((intersection.z - h.position.z) as f64 / h.velocity.vz as f64);
                    }

                    for &qq in q.iter() {
                        if qq != q[0] {
                            continue 'loop_z;
                        }
                    }
                }

                return intersection.x + intersection.y + intersection.z;
            }
        }
    }

    unreachable!()
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part2() {
        let input = "19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3";

        let res = part2(input);

        assert_eq!(res, 47);
    }
}
