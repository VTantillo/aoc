from aoc_24 import day_02

example = [
    [7, 6, 4, 2, 1],
    [1, 2, 7, 8, 9],
    [9, 7, 6, 2, 1],
    [1, 3, 2, 4, 5],
    [8, 6, 4, 4, 1],
    [1, 3, 6, 7, 9],
]


def test_report_1():
    check = day_02.check_report(example[0])
    assert check[0]


def test_report_2():
    check = day_02.check_report(example[1])
    assert not check[0]


def test_report_3():
    check = day_02.check_report(example[2])
    assert not check[0]


def test_report_4():
    check = day_02.check_report(example[3])
    assert not check[0]


def test_report_5():
    check = day_02.check_report(example[4])
    assert not check[0]


def test_report_6():
    check = day_02.check_report(example[5])
    assert check[0]


def test_part_1():
    pt1_answer = day_02.part_1(example)
    assert pt1_answer == 2


def test_part_2():
    pt2_answer = day_02.part_2(example)
    assert pt2_answer == 4
