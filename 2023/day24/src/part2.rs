use crate::parse_input;

/*
    The hailstones and the rock move in linear motion, general formula: P1 = P0 - Vt

    r = rock, 0,1,2... = hailstones
    For each hailstone: after some time t the rock and the hailstone will have the same position (they collide)
    Pr + Vr*t0 = P0 + V0*t0 -> Pr - P0 + (Vr - V0)*t0 = P1 + (Vr - V1)*t1 = P2 + (Vr - V2)*t2 = 0
    For each rock-hailstone pair, this gives 3 non-linear equations (X, Y, Z) in 7 unknowns: xr, yr, zr, vxr, vyr, vzr, t (t is different for each hailstone)
    rock-hailstone0
    xr - x0 + (vxr - vx0)*t0 = 0
    yr - y0 + (vyr - vy0)*t0 = 0
    zr - z0 + (vzr - vz0)*t0 = 0

    Step 1: remove t

    t0 can be obtained from one of the three equations and replaced in another equation
    t0 = (x0 - xr)/(vxr - vx0)
    t0 = (y0 - yr)/(vyr - vy0)
    t0 = (z0 - zr)/(vzr - vz0)

    For example, you can obtain t0 from Y and replace it in X, from Z and replce it in Y, from X and replace it in Z.
    This gives 3 non-linear equations in 6 unknowns
    xr - x0 + (vxr - vx0)*(y0 - yr)/(vyr - vy0) = 0 -> (xr - x0)*(vyr - vy0) - (vxr - vx0)*(yr - y0) = 0
    yr - y0 + (vyr - vy0)*(z0 - zr)/(vzr - vz0) = 0 -> (yr - y0)*(vzr - vz0) - (vyr - vy0)*(zr - z0) = 0
    zr - z0 + (vzr - vz0)*(x0 - xr)/(vxr - vx0) = 0 -> (zr - z0)*(vxr - vx0) - (vzr - vz0)*(xr - x0) = 0

    Step 2: make the non-linear equations linear

    Because all the hailstones will have similar equations equal to 0, the equations R-0 are equal to the equations R-1
    (xr - x0)*(vyr - vy0) - (vxr - vx0)*(yr - y0) = (xr - x1)*(vyr - vy1) - (vxr - vx1)*(yr - y1)
    (yr - y0)*(vzr - vz0) - (vyr - vy0)*(zr - z0) = (yr - y1)*(vzr - vz1) - (vyr - vy1)*(zr - z1)
    (zr - z0)*(vxr - vx0) - (vzr - vz0)*(xr - x0) = (zr - z1)*(vxr - vx1) - (vzr - vz1)*(xr - x1)

    The second degree terms cancel out, giving 3 linear equations in 6 unknowns, obtained from rock-hailstone0-hailstone1
    xr*(vy1 - vy0) + yr*(vx0 - vx1) + vxr*(y0 - y1) + vyr*(x1 - x0) = y0*vx0 - x0*vy0 + x1*vy1 - y1*vx1
    yr*(vz1 - vz0) + zr*(vy0 - vy1) + vyr*(z0 - z1) + vzr*(y1 - y0) = z0*vy0 - y0*vz0 + y1*vz1 - z1*vy1
    xr*(vz0 - vz1) + zr*(vx1 - vx0) + vxr*(z1 - z0) + vzr*(x0 - x1) = x0*vz0 - z0*vx0 + z1*vx1 - x1*vz1

    The same thing can be done with a different pair of hailstones, like rock-hailstone1-hailstone2, to get 3 more equations in the same unknowns
    xr*(vy2 - vy1) + yr*(vx1 - vx2) + vxr*(y1 - y2) + vyr*(x2 - x1) = y1*vx1 - x1*vy1 + x2*vy2 - y2*vx2
    yr*(vz2 - vz1) + zr*(vy1 - vy2) + vyr*(z1 - z2) + vzr*(y2 - y1) = y1*vz1 - z1*vy1 + z2*vy2 - y2*vz2
    xr*(vz1 - vz2) + zr*(vx2 - vx1) + vxr*(z2 - z1) + vzr*(x1 - x2) = x1*vz1 - z1*vx1 + z2*vx2 - x2*vz2

    This is now a system of 6 linear equations in 6 unknowns, that can be solved, for example, with gaussian elimination (https://en.wikipedia.org/wiki/Gaussian_elimination)
*/

pub fn part2(input: &str) -> i64 {
    let input = parse_input(input);

    let hail0 = input[0];
    let hail1 = input[1];
    let hail2 = input[2];

    let (x0, y0, z0, vx0, vy0, vz0) = (
        hail0.position.x as f64,
        hail0.position.y as f64,
        hail0.position.z as f64,
        hail0.velocity.vx as f64,
        hail0.velocity.vy as f64,
        hail0.velocity.vz as f64,
    );
    let (x1, y1, z1, vx1, vy1, vz1) = (
        hail1.position.x as f64,
        hail1.position.y as f64,
        hail1.position.z as f64,
        hail1.velocity.vx as f64,
        hail1.velocity.vy as f64,
        hail1.velocity.vz as f64,
    );
    let (x2, y2, z2, vx2, vy2, vz2) = (
        hail2.position.x as f64,
        hail2.position.y as f64,
        hail2.position.z as f64,
        hail2.velocity.vx as f64,
        hail2.velocity.vy as f64,
        hail2.velocity.vz as f64,
    );

    // 6 equations in 6 unknowns: xr, yr, zr, vxr, vyr, vzr
    let matrix = [
        [
            vy1 - vy0,
            vx0 - vx1,
            0.0,
            y0 - y1,
            x1 - x0,
            0.0,
            x1 * vy1 + y0 * vx0 - x0 * vy0 - y1 * vx1,
        ],
        [
            0.0,
            vz1 - vz0,
            vy0 - vy1,
            0.0,
            z0 - z1,
            y1 - y0,
            y1 * vz1 + z0 * vy0 - y0 * vz0 - z1 * vy1,
        ],
        [
            vz0 - vz1,
            0.0,
            vx1 - vx0,
            z1 - z0,
            0.0,
            x0 - x1,
            z1 * vx1 + x0 * vz0 - z0 * vx0 - x1 * vz1,
        ],
        [
            vy2 - vy1,
            vx1 - vx2,
            0.0,
            y1 - y2,
            x2 - x1,
            0.0,
            x2 * vy2 + y1 * vx1 - x1 * vy1 - y2 * vx2,
        ],
        [
            0.0,
            vz2 - vz1,
            vy1 - vy2,
            0.0,
            z1 - z2,
            y2 - y1,
            y2 * vz2 + z1 * vy1 - y1 * vz1 - z2 * vy2,
        ],
        [
            vz1 - vz2,
            0.0,
            vx2 - vx1,
            z2 - z1,
            0.0,
            x1 - x2,
            z2 * vx2 + x1 * vz1 - z1 * vx1 - x2 * vz2,
        ],
    ];

    let [xr, yr, zr, _vxr, _vyr, _vzr] = solve_linear_equations(&matrix);

    // the numbers may not be integers due to the precision of f64, so round them
    xr.round() as i64 + yr.round() as i64 + zr.round() as i64
}

fn solve_linear_equations<const N: usize, const M: usize>(matrix: &[[f64; M]; N]) -> [f64; N] {
    let mut matrix = *matrix;

    // gaussian elimination
    for k in 0..N {
        // Choose the largest absolute value as the k-th pivot
        let (i_max, _) = (k..N).fold((0, f64::MIN), |(i_max, max), i| {
            let a = matrix[i][k].abs();
            if a > max {
                (i, a)
            } else {
                (i_max, max)
            }
        });
        matrix.swap(k, i_max);

        // for all rows below pivot
        for i in k + 1..N {
            let factor = matrix[i][k] / matrix[k][k];
            // fill with zeros the lower part of pivot column
            matrix[i][k] = 0.0;
            // for all remaining elements in current row
            for j in k + 1..M {
                matrix[i][j] = matrix[i][j] - matrix[k][j] * factor;
            }
        }
    }

    // back substitution to find results
    let mut results = [0.0; N];
    for i in (0..N).rev() {
        results[i] = matrix[i][N];
        for j in (i + 1)..N {
            results[i] -= matrix[i][j] * results[j];
        }
        results[i] /= matrix[i][i];
    }

    results
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
