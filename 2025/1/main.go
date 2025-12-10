package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input/1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dial := 50
	zero_count := 0
	zero_count2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println(str)
		var is_left, is_right bool

		str, is_left = strings.CutPrefix(str, "L")
		str, is_right = strings.CutPrefix(str, "R")

		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}

		zero_count2 += num / 100
		num %= 100

		if is_left {
			if dial <= num && dial > 0 {
				zero_count2++
				fmt.Println("Z2")
			}

			dial -= num
		} else if is_right {
			if dial + num >= 100 {
				zero_count2++
				fmt.Println("Z2")
			}

			dial += num
		} else {
			log.Fatal("unknown direction")
		}

		dial = ((dial % 100) + 100) % 100
		if dial == 0 {
			zero_count++;
			fmt.Println("Z1")
		}
	}

	fmt.Println(zero_count)
	fmt.Println(zero_count2)
}

