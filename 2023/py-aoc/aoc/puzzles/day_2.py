from pydantic import BaseModel
from rich import print


class Day2Params(BaseModel):
    red_cubes: int
    green_cubes: int
    blue_cubes: int


class CubeNumbers(BaseModel):
    red_cubes: int
    green_cubes: int
    blue_cubes: int


class CubeGame(BaseModel):
    id: int
    games: list[CubeNumbers]


class CubeGameTotals(BaseModel):
    id: int
    totals: CubeNumbers


def day_2(input: list[str]) -> int:
    games = parse_input(input)
    answer = part_2(games)
    return answer


def part_2(games: list[CubeGame]) -> int:
    powers: list[int] = []
    for g in games:
        answer = calc_game_power(g)
        powers.append(answer)

    sum = 0
    for p in powers:
        sum += p
    return sum


def calc_game_power(game: CubeGame) -> int:
    min_red = 0
    min_green = 0
    min_blue = 0
    for g in game.games:
        if g.red_cubes > min_red:
            min_red = g.red_cubes

        if g.green_cubes > min_green:
            min_green = g.green_cubes

        if g.blue_cubes > min_blue:
            min_blue = g.blue_cubes

    return min_red * min_green * min_blue


def part_1(games: list[CubeGame], params: Day2Params) -> int:
    totals: list[CubeGameTotals] = []

    for game in games:
        total_game = CubeGameTotals(
            id=game.id, totals=CubeNumbers(red_cubes=0, green_cubes=0, blue_cubes=0)
        )

        for g in game.games:
            total_game.totals.red_cubes = (
                g.red_cubes
                if g.red_cubes > total_game.totals.red_cubes
                else total_game.totals.red_cubes
            )
            total_game.totals.blue_cubes = (
                g.blue_cubes
                if g.blue_cubes > total_game.totals.blue_cubes
                else total_game.totals.blue_cubes
            )
            total_game.totals.green_cubes = (
                g.green_cubes
                if g.green_cubes > total_game.totals.green_cubes
                else total_game.totals.green_cubes
            )
        totals.append(total_game)

    answer_totals: list[CubeGameTotals] = []
    game_ids: list[int] = []

    for t in totals:
        if (
            t.totals.blue_cubes <= params.blue_cubes
            and t.totals.red_cubes <= params.red_cubes
            and t.totals.green_cubes <= params.green_cubes
        ):
            answer_totals.append(t)
            game_ids.append(t.id)

    print(answer_totals)

    sum = 0
    for id in game_ids:
        sum += id

    return sum


def parse_input(input: list[str]) -> list[CubeGame]:
    games: list[CubeGame] = []

    for line in input:
        line_split = line.split(":")

        game_id = int(line_split[0].split(" ")[-1])

        game = CubeGame(id=game_id, games=[])

        for turn in line_split[1].split(";"):
            numbers = CubeNumbers(red_cubes=0, green_cubes=0, blue_cubes=0)

            for t in turn.split(","):
                cube_color = t.strip().split(" ")

                if cube_color[-1] == "blue":
                    numbers.blue_cubes = int(cube_color[0])
                elif cube_color[-1] == "red":
                    numbers.red_cubes = int(cube_color[0])
                elif cube_color[-1] == "green":
                    numbers.green_cubes = int(cube_color[0])

            game.games.append(numbers)

        games.append(game)

    print(games)
    return games
