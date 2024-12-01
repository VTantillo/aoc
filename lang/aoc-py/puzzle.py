from rich import print


def read_input(input: str):
    with open(input) as file:
        lines = file.readlines()
        stripped = [" ".join(line.split()) for line in lines]
        return stripped


def part_1(left_ids: list[int], right_ids: list[int]) -> int:
    left_ids.sort()
    right_ids.sort()

    total: int = 0
    for left, right in zip(left_ids, right_ids):
        total = total + abs(left - right)

    return total


def main():
    puzzle_input = read_input("./puzzles/2024/01/input.txt")

    left_ids: list[int] = []
    right_ids: list[int] = []

    for line in puzzle_input:
        split_line = line.split(" ")
        left_ids.append(int(split_line[0]))
        right_ids.append(int(split_line[1]))

    pt_1 = part_1(left_ids, right_ids)

    left_map: dict[int, list[int]] = {l_id: [] for l_id in left_ids}

    for r_id in right_ids:
        if left_map.get(r_id) is not None:
            left_map[r_id].append(r_id)

    total = 0
    for k, v in left_map.items():
        total = total + (k * len(v))

    print(f"Part 1 was {pt_1}")
    print(f"Part 2 answer: {total}")


if __name__ == "__main__":
    main()
