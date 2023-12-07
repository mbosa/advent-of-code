use std::cmp::Ordering;

#[derive(Debug, Clone, Copy, Eq, PartialEq)]
pub enum HandType {
    FiveOfAKind,
    FourOfAKind,
    FullHouse,
    ThreeOfAKind,
    TwoPair,
    TwoOfAKind,
    HighCard,
}

impl Ord for HandType {
    fn cmp(&self, other: &Self) -> Ordering {
        if self == other {
            return Ordering::Equal;
        }

        match (self, other) {
            (Self::FiveOfAKind, _) => Ordering::Greater,
            (Self::FourOfAKind, Self::FiveOfAKind) => Ordering::Less,
            (Self::FourOfAKind, _) => Ordering::Greater,
            (Self::FullHouse, Self::FiveOfAKind) => Ordering::Less,
            (Self::FullHouse, Self::FourOfAKind) => Ordering::Less,
            (Self::FullHouse, _) => Ordering::Greater,
            (Self::ThreeOfAKind, Self::FiveOfAKind) => Ordering::Less,
            (Self::ThreeOfAKind, Self::FourOfAKind) => Ordering::Less,
            (Self::ThreeOfAKind, Self::FullHouse) => Ordering::Less,
            (Self::ThreeOfAKind, _) => Ordering::Greater,
            (Self::TwoPair, Self::TwoOfAKind) => Ordering::Greater,
            (Self::TwoPair, Self::HighCard) => Ordering::Greater,
            (Self::TwoPair, _) => Ordering::Less,
            (Self::TwoOfAKind, Self::HighCard) => Ordering::Greater,
            (Self::TwoOfAKind, _) => Ordering::Less,
            (Self::HighCard, _) => Ordering::Less,
        }
    }
}

impl PartialOrd for HandType {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

pub enum Outcome {
    Win,
    Loss,
}
