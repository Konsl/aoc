import z3
import re


# this is part 2
def run_machine(desc):
    btn_p = re.compile(r"\(([^)]+)\)")
    goal_p = re.compile(r"\{([^}]+)\}")
    goal = [int(g) for g in goal_p.findall(desc)[0].split(",")]
    buttons = [[int(b) for b in b_str.split(",")] for b_str in btn_p.findall(desc)]
    print(buttons, goal)

    btn_vars = [z3.Int(f"btn_{i}") for i in range(len(buttons))]
    btns_in_range = z3.And(*[v >= 0 for v in btn_vars])

    btn_equations = []
    for i, g in enumerate(goal):
        sum_term = z3.IntVal(0)
        for j, b in enumerate(buttons):
            if i in b:
                sum_term += btn_vars[j]

        btn_equations.append(sum_term == z3.IntVal(g))
    all_equations = z3.And(*btn_equations)

    lower_bound = 0
    upper_bound = sum(goal)

    while upper_bound != lower_bound:
        print(f"range: {lower_bound} - {upper_bound}")
        test_val = (upper_bound + lower_bound) // 2
        max_check = sum(btn_vars) <= z3.IntVal(test_val)

        s = z3.Solver()
        s.add(btns_in_range)
        s.add(all_equations)
        s.add(max_check)

        if s.check() != z3.sat:
            print(f"check {test_val} failed")
            lower_bound = test_val + 1
        else:
            print(f"check {test_val} succeeded")
            upper_bound = test_val

    print(f"test result is {lower_bound} - {upper_bound}")
    return lower_bound


def main():
    with open("../input/10", "rt", encoding="utf8") as f:
        the_sum = 0
        for line in f.readlines():
            the_sum += run_machine(line.strip())
        print("the sum:", the_sum)


if __name__ == "__main__":
    main()

