import re

def has_symbol(s: str):
    return re.search(r'[^0-9.]', s) is not None

lines = None
with open('input-3.txt') as f:
    lines = f.readlines()
lines = [x.strip() for x in lines]

the_sum = 0
for i, line in enumerate(lines):
    for number_match in re.finditer(r'\d+', line):
        range_start = max(number_match.start() - 1, 0)
        range_end = min(number_match.end() + 1, len(line))
        symbol_found = False
        if i > 1 and has_symbol(lines[i - 1][range_start:range_end]):
            symbol_found = True
        if has_symbol(line[range_start:range_end]):
            symbol_found = True
        if i < len(lines) - 1 and has_symbol(lines[i + 1][range_start:range_end]):
            symbol_found = True

        if symbol_found:
            the_sum += int(number_match.group())

print(the_sum)