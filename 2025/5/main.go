package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"slices"
)

func collide(x []uint64, y []uint64) bool {
	return !(x[0] > y[1] || y[0] > x[1])
}

func test(r [][]uint64) bool {
	for ix, x := range r {
		for iy, y := range r {
			if iy != ix && collide(x, y) {
				return false
			}
		}
	}

	return true
}

func count_ranges(r [][]uint64) uint64 {
	for !test(r) {
outer_l:
		for ix, x := range r {
			for iy, y := range r {
				if iy != ix && collide(x, y) {
					new_r := make([]uint64, 2)
					new_r[0] = min(x[0], y[0])
					new_r[1] = max(x[1], y[1])
					r[ix] = new_r
					r = slices.Delete(r, iy, iy+1)
					// fmt.Println("merged", x, y, new_r)

					break outer_l
				}
			}
		}
	}

	cnt := uint64(0)
	for _, x := range r {
		cnt += x[1] + 1 - x[0]
	}

	return cnt
}

func main() {
	file, err := os.Open("../input/5")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fresh := make([][]uint64, 0)
	in_fresh := true
	good_count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()

		if str == "" {
			in_fresh = false
			continue
		}
		if in_fresh {
			parts := strings.Split(str, "-")
			start, _ := strconv.ParseUint(parts[0], 10, 64)
			end, _ := strconv.ParseUint(parts[1], 10, 64)
			range_ := make([]uint64, 2)
			range_[0] = start
			range_[1] = end

			fresh = append(fresh, range_)
		} else {
			n, _ := strconv.ParseUint(str, 10, 64)
			for _, r := range fresh {
				if n >= r[0] && n <= r[1] {
					// fmt.Println("good", n)
					good_count++
					break
				}
			}
		}
	}

	fmt.Println(good_count)
	fmt.Println(test(fresh))
	fmt.Println(count_ranges(fresh))
}

