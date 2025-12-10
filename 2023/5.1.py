import re


def numbers(s: str):
    return [int(x) for x in re.findall(r"\d+", s)]

def make_ranges(nb: list[int]):
    return (range(nb[1], nb[1] + nb[2]), range(nb[0], nb[0] + nb[2]))

seeds = None
table = []
row = []

with open("input-5.txt") as f:
    for i, line in enumerate(f):
        nb = numbers(line)

        if i == 0:
            seeds = nb
            continue

        if len(nb) == 0 and len(row) > 0:
            table.append(row)
            row = []
            continue

        if len(nb) == 0:
            continue
        
        ranges = make_ranges(nb)
        row.append(ranges)

if len(row) > 0:
    table.append(row)

def transform(n: int, transforms: list[tuple[range, range]]):
    for ranges in transforms:
        if n in ranges[0]:
            return ranges[1][0] + (n - ranges[0][0])
    return n

for transforms in table:
    for i, seed in enumerate(seeds):
        seeds[i] = transform(seed, transforms)

print(min(seeds))
