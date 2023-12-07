from aoc.puzzles.day_7 import (
    Hand,
    check_four_of_a_kind,
    check_pair,
    check_two_pair,
    count_cards,
    day_7_pt_1,
    parse_input,
)
from aoc.utils.util import read_input


def test_parse_input():
    input = read_input("../inputs/day7-ex.txt")
    hands = parse_input(input)

    assert hands == [
        Hand(cards="32T3K", bid=765),
        Hand(cards="T55J5", bid=684),
        Hand(cards="KK677", bid=28),
        Hand(cards="KTJJT", bid=220),
        Hand(cards="QQQJA", bid=483),
    ]


def test_part_1():
    input = read_input("../inputs/day7-ex.txt")
    result = day_7_pt_1(input)
    assert result == 6440


def test_part_2():
    input = read_input("../inputs/day7-ex.txt")
    result = day_7_pt_1(input)
    assert result == 5905


def test_pair():
    card_map = {"A": 1, "K": 1, "2": 1, "J": 1, "4": 1}
    result = check_pair(card_map)
    assert result is True


def test_two_pair():
    card_map = {"K": 2, "6": 1, "7": 1, "J": 1}
    result = check_two_pair(card_map)
    assert result is True


def test_four_of_a_kind():
    card_map = {"J": 1, "4": 3, "2": 1}
    result = check_four_of_a_kind(card_map)
    assert result is True


def test_count_cards():
    hand = Hand(cards="AK2J4", bid=736)
    result = count_cards(hand.cards)
    print(result)
    assert result == {"A": 1, "K": 1, "2": 1, "J": 1, "4": 1}
