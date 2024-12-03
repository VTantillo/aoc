from rich import print

from utils import read_input


def parse_input(raw: list[str]) -> list[list[int]]:
    reports: list[list[int]] = []

    for line in raw:
        report: list[int] = [int(k) for k in line.split(" ")]
        reports.append(report)

    return reports


def check_report(report: list[int]) -> tuple[bool, int]:
    diffs: list[int] = []
    for i in range(1, len(report)):
        diffs.append(report[i - 1] - report[i])

    is_increasing = report[-1] - report[0] > 0

    bads: list[int] = []
    for i, diff in enumerate(diffs):
        match is_increasing:
            case True:
                if diff < 0:
                    bads.append(i)
            case False:
                if diff > 0:
                    bads.append(i)

    print(bads)

    if report[0] > report[1] and report[1] > report[2]:
        report_is_increasing = False
    elif report[0] < report[1] and report[1] < report[2]:
        report_is_increasing = True
    else:
        print("Not increasing or decreasing")
        return False, 0

    for i in range(1, len(report)):
        difference = report[i - 1] - report[i]
        is_increasing = True if difference < 0 else False

        if is_increasing != report_is_increasing:
            print(
                f"Not consistent. diff {difference}, is_increasing {is_increasing}, report_is_increasing {report_is_increasing}"
            )
            return False, i

        if not (abs(difference) >= 1 and abs(difference) <= 3):
            print(f"Not smooth. diff {abs(difference)}")
            return False, i

    return True, -1


def part_1(reports: list[list[int]]) -> int:
    num_safe = 0
    for report in reports:
        is_report_safe = check_report(report)
        print(f"{report}, {is_report_safe}")
        if is_report_safe[0]:
            num_safe = num_safe + 1

    return num_safe


def part_2(reports: list[list[int]]) -> int:
    print(reports)
    return -1


def main() -> tuple[int, int]:
    puzzle_input = read_input("./puzzles/2024/02/ex.txt")
    reports = parse_input(puzzle_input)
    print(reports)

    pt_1 = part_1(reports)
    pt_2 = part_2(reports)

    return pt_1, pt_2


if __name__ == "__main__":
    _ = main()
