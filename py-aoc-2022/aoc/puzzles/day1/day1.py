from typing import List


def main():
    elves = read_file()

    print(f"Elf with the most calories has {step1(elves)}")
    print(f"Top 3 elf calories {step2(elves)}")


def step1(elves) -> int:
    summed = [sum(elf) for elf in elves]
    summed.sort()
    return summed[-1]


def step2(elves) -> int:
    summed = [sum(elf) for elf in elves]
    summed.sort()
    return sum(summed[-3:])


def read_file() -> List[List[int]]:
    elves = []
    with open("./input.txt", "r") as file:
        lines = file.readlines()
        elfMeals = []
        for line in lines:
            if line == "\n":
                elves.append(elfMeals)
                elfMeals = []
                continue

            elfMeals.append(int(line))

    return elves


if __name__ == "__main__":
    main()
