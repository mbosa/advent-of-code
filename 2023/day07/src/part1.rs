use std::{cmp::Ordering, collections::HashMap};

use crate::{
    data::{HandType, Outcome},
    parse_input,
};

pub fn part1(input: &str) -> u32 {
    let input = parse_input(input);

    let mut input = input
        .iter()
        .map(|&el| el.into())
        .collect::<Vec<HandPart1>>();

    input.sort_unstable_by(|a, b| match a.play(b) {
        Outcome::Win => Ordering::Greater,
        Outcome::Loss => Ordering::Less,
    });

    input
        .iter()
        .enumerate()
        .map(|(i, hand)| hand.bid * (i + 1) as u32)
        .sum()
}

struct HandPart1<'a> {
    cards: &'a str,
    bid: u32,
}

impl<'a> HandPart1<'a> {
    fn get_card_value(card: char) -> u32 {
        if card.is_ascii_digit() {
            return card.to_digit(10).unwrap();
        }

        match card {
            'T' => 10,
            'J' => 11,
            'Q' => 12,
            'K' => 13,
            'A' => 14,
            _ => unreachable!(),
        }
    }

    fn calc_type(&self) -> HandType {
        let hm = self.cards.chars().fold(HashMap::new(), |mut acc, el| {
            let v = acc.get(&el).copied().unwrap_or(0);
            acc.insert(el, v + 1);
            acc
        });

        let mut max = 0;
        let mut second_max = 0;

        for (_, count) in hm {
            if count > max {
                second_max = max;
                max = count;
            } else if count > second_max {
                second_max = count;
            }
        }

        match (max, second_max) {
            (5, _) => HandType::FiveOfAKind,
            (4, _) => HandType::FourOfAKind,
            (3, 2) => HandType::FullHouse,
            (3, _) => HandType::ThreeOfAKind,
            (2, 2) => HandType::TwoPair,
            (2, _) => HandType::TwoOfAKind,
            _ => HandType::HighCard,
        }
    }

    fn play(&self, other: &Self) -> Outcome {
        let self_type = self.calc_type();
        let other_type = other.calc_type();

        if self_type == other_type {
            for (a, b) in self.cards.chars().zip(other.cards.chars()) {
                if a == b {
                    continue;
                }

                if HandPart1::get_card_value(a) > HandPart1::get_card_value(b) {
                    return Outcome::Win;
                }
                return Outcome::Loss;
            }

            unreachable!()
        } else if self_type > other_type {
            return Outcome::Win;
        } else {
            return Outcome::Loss;
        }
    }
}

impl<'a> From<(&'a str, u32)> for HandPart1<'a> {
    fn from((cards, bid): (&'a str, u32)) -> Self {
        HandPart1 { cards, bid }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483";

        let res = part1(input);

        assert_eq!(res, 6440);
    }
}
