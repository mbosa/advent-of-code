fn is_special_char(c: char) -> bool {
    c != '.' && !c.is_ascii_digit()
}

pub fn contains_special_char(a: &[char]) -> bool {
    a.iter().any(|&c| is_special_char(c))
}

pub fn slice_char_to_u32(arr: &[char]) -> u32 {
    arr.iter()
        .fold(0, |acc, c| acc * 10 + c.to_digit(10).unwrap())
}
