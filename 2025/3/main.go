package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"slices"
	"strconv"
)

func getMax(bank string, n int) string {
	slice := strings.Split(bank[:len(bank) - n + 1], "")
	d := slices.Max(slice)
	i := slices.Index(slice, d)

	if n <= 1 {
		return d
	} else {
		return d + getMax(bank[i + 1:], n - 1)
	}
}

func main() {
	file, err := os.Open("../input/3")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total1 uint64 = 0
	var total2 uint64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()

		max1 := getMax(str, 2)
		max2 := getMax(str, 12)

		fmt.Println(str)
		fmt.Println(max1, max2)

		n1, _ := strconv.ParseUint(max1, 10, 64)
		n2, _ := strconv.ParseUint(max2, 10, 64)

		total1 += n1
		total2 += n2
	}

	fmt.Println(total1, total2)
}

