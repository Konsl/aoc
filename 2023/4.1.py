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

the_sum = 0
with open("input-4.txt") as f:
    for line in f:
        card_id, nb_win, nb_have = parse_card(line)
        
        n_winning_nb = [n for n in nb_win if n in nb_have]

        if len(n_winning_nb) > 0:
            the_sum += 0b1 << (len(n_winning_nb) - 1)

print(the_sum)