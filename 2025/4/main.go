package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Grid [][]string

func isAccessible(grid Grid, x, y int) bool {
	count := 0

	for ix := max(x - 1, 0); ix < min(x + 2, len(grid[0])); ix++ {
		for iy := max(y - 1, 0); iy < min(y + 2, len(grid)); iy++ {
			if ix == x && iy == y {
				continue
			}

			if grid[iy][ix] == "@" {
				count++
			}
		}
	}

	return count < 4
}

func countAccessible(grid Grid) int {
	count := 0

	for y := range grid {
		for x := range grid[y] {
			c := grid[y][x]

			if c == "@" && isAccessible(grid, x, y) {
				count++
				grid[y][x] = "-"
			}
		}
	}

	return count
}

func main() {
	file, err := os.Open("../input/4")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		grid = append(grid, strings.Split(str, ""))
	}

	total := 0
	new_rolls := countAccessible(grid)
	for new_rolls > 0 {
		total += new_rolls
		new_rolls = countAccessible(grid)
	}
	total += new_rolls

	fmt.Println(total)
}

