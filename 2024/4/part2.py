import numpy as np

def check_x_mas(grid, x, y):
    if grid[y][x] == "A" and grid[y + 1][x + 1] == "S" and grid[y - 1][x - 1] == "M" and grid[y + 1][x - 1] == "M" and grid[y - 1][x + 1] == "S":
        return True
    return False

def main():
    with open("input") as f:
        the_input = f.readlines()
    
    cnt = 0

    for y in range(1, len(the_input) - 1):
        for x in range(1, len(the_input[y]) - 1):
            if check_x_mas(the_input, x, y):
                cnt += 1

    print(cnt)

if __name__ == "__main__":
    main()