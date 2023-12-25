use crate::parse_input;

/*
    The hailstones move in linear motion, general formula: P1 = P0 - Vt
    The trajectories intercect when
    * x is the same for hail0 and hail1
    * y is the same for hail0 and hail1

    The hails don't need to collide, so hail0 and hail1 can cross the intersection point at different times
    x0 + vx0*t0 = x1 + vx1*t1
    y0 + vy0*t0 = y1 + vy1*t1
*/

pub fn part1(input: &str, min_val: f64, max_val: f64) -> u64 {
    let input = parse_input(input);

    let mut res = 0;

    for i in 0..input.len() - 1 {
        let hail0 = input[i];
        for j in i + 1..input.len() {
            let hail1 = input[j];

            let (x0, y0, vx0, vy0) = (
                hail0.position.x as f64,
                hail0.position.y as f64,
                hail0.velocity.vx as f64,
                hail0.velocity.vy as f64,
            );
            let (x1, y1, vx1, vy1) = (
                hail1.position.x as f64,
                hail1.position.y as f64,
                hail1.velocity.vx as f64,
                hail1.velocity.vy as f64,
            );

            let determinant = vx1 * vy0 - vx0 * vy1;

            // hail0 and hail1 trajectories are parallel
            if determinant == 0.0 {
                continue;
            }

            // time when hail0 path crosses hail1 path
            let t0 = ((y1 - y0) * vx1 - (x1 - x0) * vy1) / determinant;
            // time when hail1 path crosses hail0 path
            let t1 = ((y1 - y0) * vx0 - (x1 - x0) * vy0) / determinant;

            // one of the hails intersects in the past
            if t0 < 0.0 || t1 < 0.0 {
                continue;
            }

            // intersection position
            let x = x0 + vx0 * t0;
            let y = y0 + vy0 * t0;

            if x >= min_val && x <= max_val && y >= min_val && y <= max_val {
                res += 1;
            }
        }
    }

    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3";

        let res = part1(input, 7.0, 27.0);

        assert_eq!(res, 2);
    }
}
