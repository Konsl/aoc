package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	"strings"
	"strconv"
)

func is_silly_pattern(number uint64) bool {
	n_str := strconv.FormatUint(number, 10)

	if len(n_str) % 2 != 0 {
		return false
	}

	half_len := len(n_str) / 2
	half1 := n_str[:half_len]
	half2 := n_str[half_len:]

	return half1 == half2
}

func is_silly_pattern2(number uint64) bool {
	n_str := strconv.FormatUint(number, 10)

the_loop:
	for i := 1; i < len(n_str); i++ {
		if len(n_str) % i != 0 {
			continue
		}

		part := n_str[:i]
		for j := 1; j < len(n_str) / i; j++ {
			if n_str[j * i:(j + 1) * i] != part {
				continue the_loop
			}
		}

		return true
	}

	return false
}

func solve_part(matcher func(uint64) bool, line string) {
	var sum uint64 = 0
	
	for r := range strings.SplitSeq(line, ",") {
		r := strings.Split(r, "-")

		start, err := strconv.ParseUint(r[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		end, err := strconv.ParseUint(r[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		for i := start; i <= end; i++ {
			if matcher(i) {
				sum += i
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	file, err := os.Open("../input/2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("no line")
	}

	line := scanner.Text()
	solve_part(is_silly_pattern, line)
	solve_part(is_silly_pattern2, line)
}

