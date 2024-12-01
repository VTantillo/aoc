def read_input(input: str):
    with open(input) as file:
        lines = file.readlines()
        stripped = [" ".join(line.split()) for line in lines]
        return stripped
