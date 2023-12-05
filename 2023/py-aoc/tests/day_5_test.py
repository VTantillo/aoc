# seeds
# seed-to-soil
# soil-to-fertilizer
# fertilizer-to-water
# water-to-light
# light-to-temperature
# temperature-to-humidity
# humidity-to-location

from aoc.puzzles.day_5 import (
    RangeMap,
    get_dst,
    get_range_map,
    get_seeds,
)
from aoc.puzzles.day_5_2 import take_2
from aoc.utils.util import read_input


def test_part_1():
    input = read_input("../inputs/day5-ex.txt")
    result = take_2(input)
    assert result == 46


def test_get_seeds():
    input = read_input("../inputs/day5-ex.txt")

    result = get_seeds(input)

    assert result == [79, 14, 55, 13]


def test_get_dst():
    input = read_input("../inputs/day5-ex.txt")

    seed_to_soil = get_range_map(input, "seed-to-soil")
    seed_to_soil.sort(key=lambda x: x.src_start)

    soil_1 = get_dst(79, seed_to_soil)
    soil_2 = get_dst(14, seed_to_soil)
    soil_3 = get_dst(55, seed_to_soil)
    soil_4 = get_dst(13, seed_to_soil)

    assert soil_1 == 81
    assert soil_2 == 14
    assert soil_3 == 57
    assert soil_4 == 13

    soil_5 = get_dst(97, seed_to_soil)
    soil_6 = get_dst(98, seed_to_soil)

    assert soil_5 == 99
    assert soil_6 == 50

    fertilizer_to_water = get_range_map(input, "fertilizer-to-water")
    fertilizer_to_water.sort(key=lambda x: x.src_start)

    water_1 = get_dst(81, fertilizer_to_water)
    water_2 = get_dst(53, fertilizer_to_water)
    water_3 = get_dst(57, fertilizer_to_water)
    water_4 = get_dst(52, fertilizer_to_water)

    assert water_1 == 81
    assert water_2 == 49
    assert water_3 == 53
    assert water_4 == 41


def test_get_range_map():
    input = read_input("../inputs/day5-ex.txt")

    seed_to_soil = get_range_map(input, "seed-to-soil")

    assert seed_to_soil == [
        RangeMap(dst_start=50, src_start=98, range_len=2),
        RangeMap(dst_start=52, src_start=50, range_len=48),
    ]
    water_to_light = get_range_map(input, "water-to-light")

    assert water_to_light == [
        RangeMap(dst_start=88, src_start=18, range_len=7),
        RangeMap(dst_start=18, src_start=25, range_len=70),
    ]
