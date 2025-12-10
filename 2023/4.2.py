import re

def digits(s: str):
    return re.sub(r"\D", "", s)

def numbers(s: str):
    return [int(x) for x in re.findall(r"\d+", s)]

def parse_card(card: str):
    parts1 = card.split(":")
    card_id = int(digits(parts1[0]))

    parts2 = parts1[1].split("|")
    nb_win = numbers(parts2[0])
    nb_have = numbers(parts2[1])

    return card_id, nb_win, nb_have

def get_card_wins(card: str):
    card_id, nb_win, nb_have = parse_card(card)
    return len([n for n in nb_win if n in nb_have])

wins = None
with open("input-4.txt") as f:
    wins = [get_card_wins(line) for line in f]

counts = [1] * len(wins)
for i, count in enumerate(counts):
    for j in range(i + 1, min(i + 1 + wins[i], len(counts))):
        counts[j] += count

print(sum(counts))