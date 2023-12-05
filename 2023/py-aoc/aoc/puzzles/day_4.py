from pydantic import BaseModel
import math


class Scratchoff(BaseModel):
    num: int
    winning_nums: list[int]
    card_nums: list[int]
    matches: int = 0


def part_1(input: list[str]) -> int:
    scratch_offs = [make_scratchoff(x) for x in input]
    scores = [check_scratchoff(c) for c in scratch_offs]

    sum = 0
    for s in scores:
        sum += s

    return sum


def part_2(input: list[str]) -> int:
    scratch_offs = [make_scratchoff(x) for x in input]

    for c in scratch_offs:
        matches = 0
        for n in c.card_nums:
            if n in c.winning_nums:
                matches += 1
        c.matches = matches

    lookup = [x for x in scratch_offs]

    for i, c in enumerate(scratch_offs):
        for j in reversed(range(c.matches)):
            scratch_offs.insert(i + 1, lookup[c.num - 1 + j + 1])

    return len(scratch_offs)


def check_scratchoff(sc: Scratchoff) -> int:
    matches = 0
    for n in sc.card_nums:
        if n in sc.winning_nums:
            matches += 1

    return int(math.pow(2, matches - 1))


def make_scratchoff(line: str) -> Scratchoff:
    split_line = line.split(":")
    numbers = split_line[1].split("|")

    winning = []
    for win in numbers[0].strip().split(" "):
        if win != "":
            winning.append(int(win))

    card = []
    for c in numbers[1].strip().split(" "):
        if c != "":
            card.append(int(c))

    card_num = 0
    for x in split_line[0].strip().split(" "):
        if x != "" and x.isdigit():
            card_num = int(x)

    return Scratchoff(
        num=card_num,
        winning_nums=winning,
        card_nums=card,
    )
