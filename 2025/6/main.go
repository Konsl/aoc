package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func get_op(op string) (op_func func(uint64, uint64) uint64, start uint64) {
	switch op {
	case "+":
		op_func = func(a, b uint64) uint64 { 
			return a + b
		}
		start = 0
	case "*":
		op_func = func(a, b uint64) uint64 { 
			return a * b
		}
		start = 1
	default:
		log.Fatal("invalid op", op)
	}

	return
}

func solve_part_1(grid [][]string) uint64 {
	sum := uint64(0)

	for i := range grid[0] {
		op := grid[len(grid) - 1][i]
		op_func, acc := get_op(op)

		for _, v := range grid {
			num, err := strconv.ParseUint(v[i], 10, 64)
			fmt.Println(num, err)
			if err != nil {
				continue
			}

			acc = op_func(acc, num)
		}

		sum += acc
	}

	return sum
}

func find_whitespace_column(lines []string) int {
outer_loop:
	for i := range lines[0] {
		for j := range lines {
			if lines[j][i] != ' ' {
				continue outer_loop
			}
		}

		return i
	}

	return len(lines[0])
}

func solve_part_2(lines []string) uint64 {
	whitespace := regexp.MustCompile(`\s+`)
	ops := whitespace.Split(strings.TrimSpace(lines[len(lines) - 1]), -1)
	lines = slices.Delete(lines, len(lines) - 1, len(lines))

	sum := uint64(0)

	for len(ops) > 0 {
		end_column := find_whitespace_column(lines)
		op := ops[0]
		op_func, acc := get_op(op)

		for i := range end_column {
			num_str := ""
			for _, line := range lines {
				if line[i] != ' ' {
					num_str += line[i:i+1]
				}
			}

			num, err := strconv.ParseUint(num_str, 10, 64)
			fmt.Println(num)
			if err != nil {
				continue
			}

			acc = op_func(acc, num)
		}
		fmt.Println("    ", acc)

		sum += acc

		ops = slices.Delete(ops, 0, 1)
		for i, v := range lines {
			lines[i] = v[min(end_column + 1, len(v)):]
		}
	}

	return sum
}

func main() {
	file, err := os.Open("../input/6")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]string, 0)
	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	whitespace := regexp.MustCompile(`\s+`)
	for scanner.Scan() {
		row := whitespace.Split(strings.TrimSpace(scanner.Text()), -1)
		grid = append(grid, row)
		lines = append(lines, scanner.Text())
	}

	fmt.Println(solve_part_1(grid))
	fmt.Println(solve_part_2(lines))
}

