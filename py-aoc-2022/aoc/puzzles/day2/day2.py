from enum import auto, StrEnum
from typing import Dict, List
from dataclasses import dataclass


class Shape(StrEnum):
    ROCK = auto()
    PAPER = auto()
    SCISSORS = auto()


class Result(StrEnum):
    WIN = auto()
    DRAW = auto()
    LOSS = auto()


shape_points: Dict[Shape, int] = {Shape.ROCK: 1, Shape.PAPER: 2, Shape.SCISSORS: 3}

result_points: Dict[Result, int] = {Result.WIN: 6, Result.DRAW: 3, Result.LOSS: 0}

map_shapes: Dict[str, Shape] = {"A": Shape.ROCK, "B": Shape.PAPER, "C": Shape.SCISSORS}

map_my_shapes: Dict[str, Shape] = {
    "X": Shape.ROCK,
    "Y": Shape.PAPER,
    "Z": Shape.SCISSORS,
}

result_map: Dict[Shape, Dict[Shape, Result]] = {
    Shape.ROCK: {
        Shape.ROCK: Result.DRAW,
        Shape.PAPER: Result.LOSS,
        Shape.SCISSORS: Result.WIN,
    },
    Shape.PAPER: {
        Shape.ROCK: Result.WIN,
        Shape.PAPER: Result.DRAW,
        Shape.SCISSORS: Result.LOSS,
    },
    Shape.SCISSORS: {
        Shape.ROCK: Result.LOSS,
        Shape.PAPER: Result.WIN,
        Shape.SCISSORS: Result.DRAW,
    },
}


@dataclass
class Round:
    opponentShape: Shape
    myShape: Shape


def round_result(round: Round) -> Result:
    return result_map[round.myShape][round.opponentShape]


def round_score(result: Result, shape: Shape) -> int:
    return shape_points[shape] + result_points[result]


def read_file() -> List[Round]:
    rounds: List[Round] = []
    with open("./input.txt", "r") as f:
        lines = f.readlines()

        for l in lines:
            items = l.split(" ")
            rounds.append(Round(map_shapes[items[0]], map_my_shapes[items[1].rstrip()]))

    return rounds


def step_1():
    rounds = read_file()
    round_scores: List[int] = []

    for round in rounds:
        result = round_result(round)
        score = round_score(result=result, shape=round.myShape)
        round_scores.append(score)

    print(f"Total score: {sum(round_scores)}")


def main():
    step_1()


if __name__ == "__main__":
    main()
