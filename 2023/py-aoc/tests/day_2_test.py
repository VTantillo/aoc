# from aoc.puzzles.day_2 import CubeGame, CubeNumbers, calc_game_power, day_2
# from aoc.utils.util import read_input
#
#
# def test_part_2():
#     input = read_input("../inputs/day2-ex.txt")
#     result = day_2(input)
#     assert result == 2286
#
#
# def test_calc_game_power():
#     game = CubeGame(
#         id=1,
#         games=[
#             CubeNumbers(red_cubes=4, green_cubes=0, blue_cubes=3),
#             CubeNumbers(red_cubes=1, green_cubes=2, blue_cubes=6),
#             CubeNumbers(red_cubes=0, green_cubes=2, blue_cubes=0),
#         ],
#     )
#
#     power = calc_game_power(game)
#     assert power == 48
