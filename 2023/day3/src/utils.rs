pub struct Pos {
    pub row: i32,
    pub col: i32,
}

fn is_special_char(c: char) -> bool {
    if c == '.' {
        return false;
    }
    if c.is_ascii_digit() {
        return false;
    }

    return true;
}

pub fn contains_special_char(a: &[char]) -> bool {
    a.iter().any(|&c| is_special_char(c))
}

pub fn calc_num_right(grid: &Vec<Vec<char>>, pos: Pos) -> Option<i32> {
    if !grid[pos.row as usize][pos.col as usize].is_ascii_digit() {
        return None;
    }

    let cols = grid[0].len() as i32;

    let mut num: i32 = 0;
    let row = &grid[pos.row as usize];
    let mut i = pos.col;

    while i < cols && row[i as usize].is_ascii_digit() {
        let digit = row[i as usize].to_digit(10).unwrap() as i32;

        num = num * 10 + digit;

        i += 1;
    }

    Some(num)
}

pub fn calc_num_left(grid: &Vec<Vec<char>>, pos: Pos) -> Option<i32> {
    if !grid[pos.row as usize][pos.col as usize].is_ascii_digit() {
        return None;
    }

    let mut num: i32 = 0;
    let row = &grid[pos.row as usize];
    let mut i = pos.col;

    while i >= 0 && row[i as usize].is_ascii_digit() {
        let digit = row[i as usize].to_digit(10).unwrap() as i32;

        num += digit * 10i32.pow((pos.col - i) as u32);

        i -= 1;
    }

    Some(num)
}
