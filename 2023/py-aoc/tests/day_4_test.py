# from aoc.puzzles.day_4 import (
#     Scratchoff,
#     check_scratchoff,
#     part_1,
#     make_scratchoff,
#     part_2,
# )
# from aoc.utils.util import read_input
#
#
# def test_pt_1():
#     input = read_input("../inputs/day4-ex.txt")
#     result = part_1(input)
#     assert result == 13
#
#
# def test_pt_2():
#     input = read_input("../inputs/day4-ex.txt")
#     result = part_2(input)
#     assert result == 30
#
#
# def test_make_scratch_off():
#     input = "Card 1: 41 48 83 86 17 | 83 86 6 31 17 9 48 53"
#     scratchoff = make_scratchoff(input)
#
#     assert scratchoff.num == 1
#     assert scratchoff.winning_nums == [41, 48, 83, 86, 17]
#     assert scratchoff.card_nums == [83, 86, 6, 31, 17, 9, 48, 53]
#
#
# def test_check_scratchoff():
#     sc_1 = Scratchoff(
#         num=1,
#         winning_nums=[41, 48, 83, 86, 17],
#         card_nums=[83, 86, 6, 31, 17, 9, 48, 53],
#     )
#     sc_2 = Scratchoff(
#         num=2,
#         winning_nums=[13, 32, 20, 16, 61],
#         card_nums=[61, 30, 68, 82, 17, 32, 24, 19],
#     )
#     sc_3 = Scratchoff(
#         num=3,
#         winning_nums=[1, 21, 53, 59, 44],
#         card_nums=[69, 82, 63, 72, 16, 21, 14, 1],
#     )
#
#     score_1 = check_scratchoff(sc_1)
#     score_2 = check_scratchoff(sc_2)
#     score_3 = check_scratchoff(sc_3)
#
#     assert score_1 == 8
#     assert score_2 == 2
#     assert score_3 == 2
