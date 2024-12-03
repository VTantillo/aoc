import re

from rich import print

from utils import read_input


def part_1(mem: list[str]) -> int:
    all_muls: list[str] = []
    for line in mem:
        muls = re.findall("mul\(\d+,\d+\)", line)  # pyright: ignore[reportInvalidStringEscapeSequence]
        for mul in muls:  # pyright: ignore[reportAny]
            match mul:
                case str():
                    all_muls.append(mul)
                case _:  # pyright: ignore[reportAny]
                    pass

    total = 0
    for mul in all_muls:
        vals = mul[4:-1]
        nums = vals.split(",")

        total = total + (int(nums[0]) * int(nums[1]))

    return total


def part_2(mem: list[str]) -> int:
    print(mem)

    all_matches: list[str] = []

    for line in mem:
        matches = re.findall("mul\(\d+,\d+\)|don't\(\)|do\(\)", line)  # pyright: ignore[reportInvalidStringEscapeSequence]
        for m in matches:  # pyright: ignore[reportAny]
            match m:
                case str():
                    all_matches.append(m)
                case _:  # pyright: ignore[reportAny]
                    print("idk bro", m)
                    pass

    print(all_matches)

    is_mul_enabled = True

    for m in all_matches:
        match m:
            case str(x) if "mul" in x:
                print(f"Handle mul {is_mul_enabled}")
                if not is_mul_enabled:
                    continue
                pass
            case str(x) if "do()" in x:
                print("Mul is enabled")
                is_mul_enabled = True
            case str(x) if "don't()" in x:
                print("Mul is disabled")
                is_mul_enabled = False
            case _:
                print(f"Idk, what that instruction is {m}")

    total = 0

    return total


def main() -> tuple[int, int]:
    puzzle_input = read_input("./puzzles/2024/03/input.txt")

    pt_1 = part_1(puzzle_input)
    pt_2 = part_2(puzzle_input)

    return pt_1, pt_2


if __name__ == "__main__":
    _ = main()
