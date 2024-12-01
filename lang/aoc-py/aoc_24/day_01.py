from utils import read_input


def part_1(left_ids: list[int], right_ids: list[int]) -> int:
    left_ids.sort()
    right_ids.sort()

    total: int = 0
    for left, right in zip(left_ids, right_ids):
        total = total + abs(left - right)

    return total


def part_2(left_ids: list[int], right_ids: list[int]) -> int:
    left_map: dict[int, list[int]] = {l_id: [] for l_id in left_ids}

    for r_id in right_ids:
        if left_map.get(r_id) is not None:
            left_map[r_id].append(r_id)

    total = 0
    for k, v in left_map.items():
        total = total + (k * len(v))

    return total


def main() -> tuple[int, int]:
    puzzle_input = read_input("./puzzles/2024/01/input.txt")

    left_ids: list[int] = []
    right_ids: list[int] = []

    for line in puzzle_input:
        split_line = line.split(" ")
        left_ids.append(int(split_line[0]))
        right_ids.append(int(split_line[1]))

    pt_1 = part_1(left_ids, right_ids)
    pt_2 = part_2(left_ids, right_ids)

    return pt_1, pt_2


if __name__ == "__main__":
    _ = main()
