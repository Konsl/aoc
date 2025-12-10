package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func euc_dist2(a, b []uint64) uint64 {
	d0 := int64(a[0]) - int64(b[0])
	d1 := int64(a[1]) - int64(b[1])
	d2 := int64(a[2]) - int64(b[2])

	return uint64(d0 * d0 + d1 * d1 + d2 * d2)
}

func find_nearest_nodes(circuits [][][]uint64, minDist uint64) (c1, e1, c2, e2 int, dist uint64) {
	c1 = 0
	e1 = 0
	c2 = 0
	e2 = 0

	dist = math.MaxUint64

	for ic1 := range circuits {
		for ie1 := range circuits[ic1] {
			for ic2 := range circuits {
				for ie2 := range circuits[ic2] {
					if ic2 == ic1 && ie1 == ie2 {
						continue
					}

					d := euc_dist2(circuits[ic1][ie1], circuits[ic2][ie2])
					if d > minDist && d < dist {
						dist = d
						c1 = ic1
						e1 = ie1
						c2 = ic2
						e2 = ie2
					}
				}
			}
		}
	}

	return
}

func solve_part_1(nodes [][]uint64) uint64 {
	circuits := make([][][]uint64, 0)
	for _, n := range nodes {
		c := make([][]uint64, 1)
		c[0] = n
		circuits = append(circuits, c)
	}

	last_dist := uint64(0)
	for range 1000 {
		c1, e1, c2, e2, dist := find_nearest_nodes(circuits, last_dist)
		last_dist = dist
		if c1 == c2 {
			fmt.Println("skipping", circuits[c1][e1], circuits[c2][e2], dist)
			continue
		}
		fmt.Println("merging", circuits[c1][e1], circuits[c2][e2], dist)


		for _, v := range circuits[c2] {
			circuits[c1] = append(circuits[c1], v)
		}

		circuits = slices.Delete(circuits, c2, c2 + 1)
	}

	slices.SortFunc(circuits, func(a, b [][]uint64) int {
		return len(b) - len(a)
	})

	return uint64(len(circuits[0]) * len(circuits[1]) * len(circuits[2]))
}

func solve_part_2(nodes [][]uint64) uint64 {
	circuits := make([][][]uint64, 0)
	for _, n := range nodes {
		c := make([][]uint64, 1)
		c[0] = n
		circuits = append(circuits, c)
	}

	last_dist := uint64(0)
	for {
		c1, e1, c2, e2, dist := find_nearest_nodes(circuits, last_dist)
		last_dist = dist
		if c1 == c2 {
			fmt.Println("skipping", circuits[c1][e1], circuits[c2][e2], dist)
			continue
		}
		fmt.Println("merging", circuits[c1][e1], circuits[c2][e2], dist, "len", len(circuits) - 1)

		e2 += len(circuits[c1])
		for _, v := range circuits[c2] {
			circuits[c1] = append(circuits[c1], v)
		}

		circuits = slices.Delete(circuits, c2, c2 + 1)

		if len(circuits) == 1 {
			a := circuits[c1][e1]
			b := circuits[c1][e2]
			return uint64(a[0] * b[0])
		}
	}
}

func main() {
	file, err := os.Open("../input/8")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nodes := make([][]uint64, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		node := make([]uint64, 3)
		for i := range node {
			node[i], _ = strconv.ParseUint(parts[i], 10, 64)
		}

		nodes = append(nodes, node)
	}

	// fmt.Println(solve_part_1(nodes))
	fmt.Println(solve_part_2(nodes))
}

