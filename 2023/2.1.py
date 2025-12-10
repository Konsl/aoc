import re

r = 12
g = 13
b = 14


def digits(s: str):
    return re.sub(r"\D", "", s)


def parse_grab(s: str):
    parts = [part.strip() for part in s.split(",")]
    the_r = 0
    the_g = 0
    the_b = 0
    for part in parts:
        if part.endswith("red"):
            the_r = int(digits(part))
        elif part.endswith("green"):
            the_g = int(digits(part))
        elif part.endswith("blue"):
            the_b = int(digits(part))
    return the_r, the_g, the_b


the_sum = 0
with open("input-2.txt") as f:
    for line in f:
        parts1 = line.split(":")

        the_id = int(digits(parts1[0]))
        grabs = [parse_grab(grab) for grab in parts1[1].split(";")]

        if all(the_r <= r and the_g <= g and the_b <= b for the_r, the_g, the_b in grabs):
                the_sum += the_id

print(the_sum)
