from aoc.puzzles.day_6 import calc_race_distance, part_1, part_2
from aoc.utils.util import read_input


def test_part_1():
    input = read_input("../inputs/day6-ex.txt")
    result = part_1(input)
    assert result == 288


def test_part_2():
    input = read_input("../inputs/day6-ex.txt")
    result = part_2(input)
    assert result == 71503


def test_race_distance():
    race_1 = calc_race_distance(0, 7)
    race_2 = calc_race_distance(1, 7)
    race_3 = calc_race_distance(2, 7)
    race_4 = calc_race_distance(3, 7)
    race_5 = calc_race_distance(4, 7)
    race_6 = calc_race_distance(5, 7)
    race_7 = calc_race_distance(6, 7)
    race_8 = calc_race_distance(7, 7)

    assert race_1 == 0
    assert race_2 == 6
    assert race_3 == 10
    assert race_4 == 12
    assert race_5 == 12
    assert race_6 == 10
    assert race_7 == 6
    assert race_8 == 0
