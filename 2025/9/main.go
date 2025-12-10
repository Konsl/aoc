package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func get_rect(t1, t2 []int) (c1, c2 []int, area int) {
	lx := t1[0]
	hx := t2[0]
	if lx > hx {
		lx, hx = hx, lx
	}

	ly := t1[1]
	hy := t2[1]
	if ly > hy {
		ly, hy = hy, ly
	}

	c1 = []int{lx, ly}
	c2 = []int{hx, hy}
	area = (hx - lx + 1) * (hy - ly + 1)
	return
}

func solve_part_1(tiles [][]int) uint64 {
	max_area := 0

	for t1 := range tiles {
		for t2 := range tiles {
			_, _, area := get_rect(tiles[t1], tiles[t2])
			if area > max_area {
				max_area = area
			}
		}
	}

	return uint64(max_area)
}

func is_tile_rg(t []int, tiles [][]int) bool {
	count := 0

	prev := tiles[len(tiles) - 1]
	last_dir := 0
	for _, cur := range tiles {
		if t[0] == cur[0] && t[1] == cur[1] {
			return true
		}

		x_range := t[0] >= min(cur[0], prev[0]) && t[0] <= max(cur[0], prev[0])
		y_range := t[1] >= min(cur[1], prev[1]) && t[1] <= max(cur[1], prev[1])
		//in_x_range := t[0] > min(cur[0], prev[0]) && t[0] < max(cur[0], prev[0])
		in_y_range := t[1] > min(cur[1], prev[1]) && t[1] < max(cur[1], prev[1])
		is_horizontal := cur[1] == prev[1]

		if x_range && y_range {
			return true
		}

		if in_y_range && cur[0] > t[0] {
			count++
		}

		if !is_horizontal && y_range && !in_y_range && cur[0] > t[0] {
			var dir int
			if prev[1] - cur[1] > 0 {
				dir = -1
			} else {
				dir = 1
			}

			switch last_dir {
			case 0:
				last_dir = dir
			case dir:
				last_dir = 0
				count++
			default:
				last_dir = 0
			}
		}

		prev = cur
	}

	return count % 2 == 1
}

func is_rect_invalid(tiles [][]int, c1, c2 []int) bool {
	for _, c3 := range tiles {
		if c3[0] > c1[0] && c3[0] < c2[0] && c3[1] > c1[1] && c3[1] < c2[1] {
			return true
		}
	}

	for x := c1[0]; x <= c2[0]; x++ {
		if !is_tile_rg([]int{x, c1[1]}, tiles) {
			return true
		}
		if !is_tile_rg([]int{x, c2[1]}, tiles) {
			return true
		}
	}
	for y := c1[1]; y <= c2[1]; y++ {
		if !is_tile_rg([]int{c1[0], y}, tiles) {
			return true
		}
		if !is_tile_rg([]int{c2[0], y}, tiles) {
			return true
		}
	}
	return false
}

func solve_part_2(tiles [][]int) uint64 {
	max_area := 0

	for t1 := range tiles {
	tile_loop:
		for t2 := range tiles {
			c1, c2, area := get_rect(tiles[t1], tiles[t2])

			fmt.Println("testing", c1, c1, area)
			if is_rect_invalid(tiles, c1, c2) {
				continue tile_loop
			}

			if area > max_area {
				fmt.Println("found", area, c1, c2)
				max_area = area
			}
		}
	}

	return uint64(max_area)
}

func main() {
	file, err := os.Open("../input/9")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tiles := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		tiles = append(tiles, []int{x, y})
	}

	for y := range 9 {
		for x := range 14 {
			if is_tile_rg([]int{x, y}, tiles) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println(solve_part_1(tiles))
	fmt.Println(solve_part_2(tiles))
}


