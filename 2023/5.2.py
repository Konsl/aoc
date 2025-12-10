import re


def numbers(s: str):
    return [int(x) for x in re.findall(r"\d+", s)]

def make_ranges(nb: list[int]):
    return ((nb[1], nb[1] + nb[2]), (nb[0], nb[0] + nb[2]))

seeds = []
table = []
row = []

with open("input-5.txt") as f:
    for i, line in enumerate(f):
        nb = numbers(line)

        if i == 0:
            seeds = nb
            continue

        if len(nb) == 0 and len(row) > 0:
            row.sort(key=lambda x: x[0][0])
            table.append(row)
            row = []
            continue

        if len(nb) == 0:
            continue
        
        ranges = make_ranges(nb)
        row.append(ranges)

if len(row) > 0:
    row.sort(key=lambda x: x[0][0])
    table.append(row)

parsed_seeds = []
for i in range(len(seeds) // 2):
    parsed_seeds.append((seeds[i * 2], seeds[i * 2] + seeds[i * 2 + 1]))
seeds = parsed_seeds

def split_ranges(inner: tuple[int, int], outer: tuple[int, int]) -> tuple[tuple[int, int], tuple[int, int], tuple[int, int]]:
    if inner[0] >= outer[1]:
        return (outer, None, None)
    if inner[1] <= outer[0]:
        return (None, None, outer)
    
    lower_bound = max(inner[0], outer[0])
    upper_bound = min(inner[1], outer[1])

    if lower_bound == outer[0]:
        if upper_bound == outer[1]:
            return (None, (lower_bound, upper_bound), None)
        return (None, (lower_bound, upper_bound), (upper_bound, outer[1]))
    if upper_bound == outer[1]:
        return ((outer[0], lower_bound), (lower_bound, upper_bound), None)
    return ((outer[0], lower_bound), (lower_bound, upper_bound), (upper_bound, outer[1]))

def map_range_part(s: tuple[int, int], d: tuple[int, int], r: tuple[int, int]):
    if r == None:
        return None
    return (r[0] - s[0] + d[0], r[1] - s[0] + d[0])

def transform_range(r: tuple[int, int], transforms: list[tuple[tuple[int, int], tuple[int, int]]]):
    output = []
    for ranges in transforms:
        splits = split_ranges(ranges[0], r)
        splits_transformed = (splits[0], map_range_part(ranges[0], ranges[1], splits[1]), splits[2])
        output.append(splits_transformed[0])
        output.append(splits_transformed[1])
        r = splits_transformed[2]

        if r == None:
            return [entry for entry in output if entry != None]
    return [entry for entry in output if entry != None]

for transforms in table:
    new_seeds = []
    for seed in seeds:
        for transformed in transform_range((seed[0], seed[1]), transforms):
            new_seeds.append(transformed)
    seeds = new_seeds

print(seeds)
