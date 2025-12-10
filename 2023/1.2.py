import re

the_sum = 0
the_pattern = re.compile(r'[^\d]+')
letter_pattern = re.compile(r'(one|two|three|four|five|six|seven|eight|nine)')

replacements = {
    'one': '1',
    'two': '2',
    'three': '3',
    'four': '4',
    'five': '5',
    'six': '6',
    'seven': '7',
    'eight': '8',
    'nine': '9'
}

def remove_letters(line):
    line = letter_pattern.sub(lambda letters: replacements[letters.group(0)], line)
    line = the_pattern.sub('', line)
    return line

with open('input-1.txt') as f:
    for line in f:
        line = remove_letters(line)
        the_sum += int(line[0] + line[-1])

print(the_sum)