from pydantic import BaseModel
from rich import print


class Race(BaseModel):
    time: int
    distance: int


def part_1(input: list[str]) -> int:
    races = parse_input(input)

    possibilities: list[int] = []
    for race in races:
        solutions = 0
        for t in range(race.time):
            distance = calc_race_distance(t, race.time)
            if distance > race.distance:
                solutions += 1
        possibilities.append(solutions)

    answer = 0
    for p in possibilities:
        if answer == 0:
            answer += p
        else:
            answer *= p

    return answer


def part_2(input: list[str]) -> int:
    race = pase_input_2(input)

    solutions = 0
    for t in range(race.time):
        distance = calc_race_distance(t, race.time)
        if distance > race.distance:
            solutions += 1

    return solutions


def calc_race_distance(speed: int, time: int) -> int:
    remaining_time = time - speed

    if remaining_time < 0:
        return 0

    return remaining_time * speed


def parse_input(input: list[str]) -> list[Race]:
    races: list[Race] = []
    for line_num, line in enumerate(input):
        line_split = line.split(" ")
        if line_num == 0:
            for i in range(1, len(line_split)):
                print(line_split[i])
                race = Race(time=int(line_split[i]), distance=-1)
                races.append(race)
        else:
            for i in range(1, len(line_split)):
                races[i - 1].distance = int(line_split[i])

    print(races)
    return races


def pase_input_2(input: list[str]) -> Race:
    race = Race(time=-1, distance=-1)
    for line_num, line in enumerate(input):
        line_split = line.split(":")
        if line_num == 0:
            race.time = int(line_split[1].replace(" ", ""))
        else:
            race.distance = int(line_split[1].replace(" ", ""))

    return race
