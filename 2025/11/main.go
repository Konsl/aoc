package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func count_paths_inner(graph map[string][]string, path []string, goal string, cache map[string]uint64) uint64 {
	current_node := path[len(path) - 1]
	cache_val, cache_ok := cache[current_node]
	if cache_ok {
		return cache_val
	}

	count := uint64(0)
	for _, next_node := range graph[current_node] {
		if next_node == goal {
			fmt.Println("passed", path, next_node)
			count++
			continue
		}

		if slices.Contains(path, next_node) {
			log.Fatal("cycle detected", path, next_node)
		}

		next_path := append(path, next_node)
		count += count_paths_inner(graph, next_path, goal, cache)
	}

	cache[current_node] = count
	return count
}

func count_paths(graph map[string][]string, start string, goal string) uint64 {
	return count_paths_inner(graph, []string{start}, goal, make(map[string]uint64))
}

func main() {
	file, err := os.Open("../input/11")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	graph := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name, dev_con_s, _ := strings.Cut(scanner.Text(), ":")
		dev_con := strings.Split(strings.TrimSpace(dev_con_s), " ")
		graph[name] = dev_con
	}

	fmt.Println(graph)
	part_1 := count_paths(graph, "you", "out")
	way_1 := count_paths(graph, "svr", "fft") * count_paths(graph, "fft", "dac") * count_paths(graph, "dac", "out")
	way_2 := count_paths(graph, "svr", "dac") * count_paths(graph, "dac", "fft") * count_paths(graph, "fft", "out")
	part_2 := way_1 + way_2

	fmt.Println(part_1)
	fmt.Println(part_2)
}

