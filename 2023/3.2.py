import re

def get_gear_indices(s: str, offset: int):
    return [m.start() + offset for m in re.finditer(r'\*', s)]

def has_symbol(s: str):
    return re.search(r'[^0-9.]', s) is not None

lines = []
with open('input-3.txt') as f:
    lines = f.readlines()
lines = [x.strip() for x in lines]

gears = {}

def add_gear_number(line: int, col: int, number: int):
    if (line, col) not in gears:
        gears[(line, col)] = []
    gears[(line, col)].append(number)

for i, line in enumerate(lines):
    for number_match in re.finditer(r'\d+', line):
        range_start = max(number_match.start() - 1, 0)
        range_end = min(number_match.end() + 1, len(line))
        
        if i > 1:
            for ind in get_gear_indices(lines[i - 1][range_start:range_end], range_start):
                add_gear_number(i - 1, ind, int(number_match.group()))
        
        for ind in get_gear_indices(line[range_start:range_end], range_start):
            add_gear_number(i, ind, int(number_match.group()))
        
        if i < len(lines) - 1:
            for ind in get_gear_indices(lines[i + 1][range_start:range_end], range_start):
                add_gear_number(i + 1, ind, int(number_match.group()))

the_sum = 0
for line, col in gears:
    if len(gears[(line, col)]) == 2:
        the_sum += gears[(line, col)][0] * gears[(line, col)][1]
print(the_sum)