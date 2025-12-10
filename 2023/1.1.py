import re

the_sum = 0
the_pattern = re.compile(r'[^\d]+')

with open('input-1.txt') as f:
    for line in f:
        line = the_pattern.sub('', line)
        the_sum += int(line[0] + line[-1])

print(the_sum)