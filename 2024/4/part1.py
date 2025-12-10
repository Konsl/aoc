import numpy as np

def count_in_2d(grid):
    return sum(["".join(l).count("XMAS") for l in grid])

def get_diags(a):
    return [a[::-1,:].diagonal(i) for i in range(-a.shape[0]+1,a.shape[1])]

def count_straight_diag(grid):
    cnt = count_in_2d(grid)
    cnt += count_in_2d(get_diags(np.array(grid)))
    return cnt

def main():
    with open("input") as f:
        the_input = f.readlines()
    the_input = [x.strip() for x in the_input]
    the_input = np.array([list(x) for x in the_input])

    cnt = count_straight_diag(the_input)
    the_input = np.rot90(the_input)
    cnt += count_straight_diag(the_input)
    the_input = np.rot90(the_input)
    cnt += count_straight_diag(the_input)
    the_input = np.rot90(the_input)
    cnt += count_straight_diag(the_input)

    print(cnt)

if __name__ == "__main__":
    main()