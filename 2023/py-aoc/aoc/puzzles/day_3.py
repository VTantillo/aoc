def day_3(input: list[str]) -> int:
    symbols = set([])
    for line in input:
        for char in line:
            if (
                not (char >= "0" and char <= "9")
                and not char == "."
                and not char == "\n"
            ):
                symbols.add(char)

    part_numbers = []

    for l_num, line in enumerate(input):
        is_in_number = False
        line_numbers = []
        for c_num, char in enumerate(line):
            if char == "." or char in symbols:
                is_in_number = False
            elif is_digit(char) and not is_in_number:
                is_in_number = True
                ratio = find_gear_ratios(symbols, input, l_num, c_num)
                line_numbers.append(ratio)
        part_numbers.extend(line_numbers)

    sum = 0
    for n in part_numbers:
        sum += n

    return sum


def is_digit(char: str) -> bool:
    return char >= "0" and char <= "9"


def find_part_number(
    symbols: set[str], input: list[str], line_num: int, char_num: int
) -> int:
    in_number = True
    num_end = char_num

    while in_number:
        if (
            input[line_num][num_end] == "."
            or input[line_num][num_end] in symbols
            or num_end == len(input[line_num]) - 1
        ):
            in_number = False
        else:
            num_end += 1

    num = int(input[line_num][char_num:num_end])

    num_end -= 1

    first_row = line_num == 0
    last_row = line_num == len(input) - 1
    first_col = char_num == 0
    last_col = num_end == len(input[line_num]) - 1

    for i in range(char_num, num_end + 1):
        if not first_row and input[line_num - 1][i] in symbols:
            return num
        if not last_row and input[line_num + 1][i] in symbols:
            return num
        if i == char_num:
            if (
                not first_row
                and not first_col
                and input[line_num - 1][i - 1] in symbols
            ):
                return num
            if not last_row and not first_col and input[line_num + 1][i - 1] in symbols:
                return num
            if not first_col and input[line_num][i - 1] in symbols:
                return num
        if i == num_end:
            if not first_row and not last_col and input[line_num - 1][i + 1] in symbols:
                return num
            if not last_row and not last_col and input[line_num + 1][i + 1] in symbols:
                return num
            if not last_col and input[line_num][i + 1] in symbols:
                return num

    return -1


def find_gear_ratios(
    symbols: set[str], input: list[str], line_num: int, char_num: int
) -> int:
    num, num_start, num_end = find_number(symbols, input, line_num, char_num)

    first_row = line_num == 0
    last_row = line_num == len(input) - 1
    first_col = num_start == 0
    last_col = num_end == len(input[line_num]) - 1

    for i in range(char_num, num_end + 1):
        if not first_row and input[line_num - 1][i] == "*":
            return find_other_ratio(symbols, input, line_num - 1, i, num)
        if not last_row and input[line_num + 1][i] == "*":
            return find_other_ratio(symbols, input, line_num + 1, i, num)
        if i == char_num:
            if not first_row and not first_col and input[line_num - 1][i - 1] == "*":
                return find_other_ratio(symbols, input, line_num - 1, i - 1, num)
            if not last_row and not first_col and input[line_num + 1][i - 1] == "*":
                return find_other_ratio(symbols, input, line_num + 1, i - 1, num)
            if not first_col and input[line_num][i - 1] == "*":
                return find_other_ratio(symbols, input, line_num, i - 1, num)
        if i == num_end:
            if not first_row and not last_col and input[line_num - 1][i + 1] == "*":
                return find_other_ratio(symbols, input, line_num - 1, i + 1, num)
            if not last_row and not last_col and input[line_num + 1][i + 1] == "*":
                return find_other_ratio(symbols, input, line_num + 1, i + 1, num)
            if not last_col and input[line_num][i + 1] == "*":
                return find_other_ratio(symbols, input, line_num, i + 1, num)

    return 0


def find_other_ratio(
    symbols: set[str], input: list[str], line_num: int, char_num: int, num: int
) -> int:
    first_row = line_num == 0
    last_row = line_num == len(input) - 1
    first_col = char_num == 0
    last_col = char_num == len(input[line_num]) - 1

    second_num = (0, 0, 0)

    if not first_row and is_digit(input[line_num - 1][char_num]):
        second_num = find_number(symbols, input, line_num - 1, char_num)
        if second_num[0] == num:
            second_num = (0, 0, 0)
    if not last_row and is_digit(input[line_num + 1][char_num]):
        second_num = find_number(symbols, input, line_num + 1, char_num)
        if second_num[0] == num:
            second_num = (0, 0, 0)
    if not first_row and not first_col and is_digit(input[line_num - 1][char_num - 1]):
        second_num = find_number(symbols, input, line_num - 1, char_num - 1)
        if second_num[0] == num:
            second_num = (0, 0, 0)
    if not last_row and not first_col and is_digit(input[line_num + 1][char_num - 1]):
        second_num = find_number(symbols, input, line_num + 1, char_num - 1)
        if second_num[0] == num:
            second_num = (0, 0, 0)
    if not first_col and is_digit(input[line_num][char_num - 1]):
        second_num = find_number(symbols, input, line_num, char_num - 1)
        if second_num[0] == num:
            second_num = (0, 0, 0)
    if not first_row and not last_col and is_digit(input[line_num - 1][char_num + 1]):
        second_num = find_number(symbols, input, line_num - 1, char_num + 1)
        if second_num[0] == num:
            second_num = (0, 0, 0)
    if not last_row and not last_col and is_digit(input[line_num + 1][char_num + 1]):
        second_num = find_number(symbols, input, line_num + 1, char_num + 1)
        if second_num[0] == num:
            second_num = (0, 0, 0)
    if not last_col and is_digit(input[line_num][char_num + 1]):
        second_num = find_number(symbols, input, line_num, char_num + 1)
        if second_num[0] == num:
            second_num = (0, 0, 0)

    print(f"num: {num}, second: {second_num[0]}, ratio: {num*second_num[0]}")

    return num * second_num[0]


def find_number(
    symbols: set[str], input: list[str], line_num: int, char_num: int
) -> tuple[int, int, int]:
    in_number = True
    num_start = char_num
    num_end = char_num

    while in_number:
        if (
            input[line_num][num_start] == "."
            or input[line_num][num_start] in symbols
            or num_start == 0
        ):
            in_number = False
        else:
            num_start -= 1

    in_number = True
    while in_number:
        if (
            input[line_num][num_end] == "."
            or input[line_num][num_end] in symbols
            or num_end == len(input[line_num]) - 1
        ):
            in_number = False
        else:
            num_end += 1

    if input[line_num][num_start] == "." or input[line_num][num_start] in symbols:
        num_start += 1
    num = int(input[line_num][num_start:num_end])
    num_end -= 1

    return (num, num_start, num_end)
