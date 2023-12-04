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
                part_num = find_part_number(symbols, input, l_num, c_num)
                if part_num != -1:
                    line_numbers.append(part_num)
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
