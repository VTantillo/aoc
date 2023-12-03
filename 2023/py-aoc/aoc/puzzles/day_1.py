def day_1(input: list[str]) -> int:
    total = 0
    for line in input:
        stripped = line.strip()
        print(stripped)
        digits = ""
        letters = ""

        for i in range(len(stripped)):
            if stripped[i] >= "0" and stripped[i] <= "9":
                digits += stripped[i]
                break

            letters += stripped[i]
            digit = digit_word(letters)

            if digit is not None:
                digits += stripped[i]
                break

        letters = ""
        for i in reversed(range(len(stripped))):
            print(i)

    return int(total)


def digit_word(input: str) -> str | None:
    match input:
        case str(x) if "one" in x:
            return "1"
        case str(x) if "two" in x:
            return "2"
        case str(x) if "three" in x:
            return "3"
        case str(x) if "four" in x:
            return "4"
        case str(x) if "five" in x:
            return "5"
        case str(x) if "six" in x:
            return "6"
        case str(x) if "seven" in x:
            return "7"
        case str(x) if "eight" in x:
            return "8"
        case str(x) if "nine" in x:
            return "9"
        case _:
            return None
