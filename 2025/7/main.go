package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func solve_part_1(grid [][]string) uint64 {
	beams := make(map[int]bool)
	splits := 0

	for _, line := range grid {
		for i, field := range line {
			switch field {
			case "S":
				beams[i] = true
			case "^":
				v, ok := beams[i]
				if v && ok {
					beams[i] = false
					beams[i + 1] = true
					beams[i - 1] = true

					splits++
				}
			}
		}

		fmt.Println(beams)
	}

	return uint64(splits)
}

func solve_part_2(grid [][]string) uint64 {
	beams := make(map[int]int)

	for _, line := range grid {
		for i, field := range line {
			switch field {
			case "S":
				beams[i]++
			case "^":
				v, ok := beams[i]
				if v > 0 && ok {
					beams[i + 1] += beams[i]
					beams[i - 1] += beams[i]
					beams[i] = 0
				}
			}
		}
	}

	sum := 0
	for _, c := range beams {
		sum += c
	}

	return uint64(sum)
}

func main() {
	file, err := os.Open("../input/7")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, strings.Split(scanner.Text(), ""))
	}

	fmt.Println(solve_part_1(data))
	fmt.Println(solve_part_2(data))
}

