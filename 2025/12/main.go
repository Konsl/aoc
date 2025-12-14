package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Present struct {
	tiles [][]bool
	tileCount int
}

type TestCase struct {
	width, height int
	presents []int
}

func main() {
	file, err := os.Open("../input/12")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	presents := make([]Present, 0)
	testCases := make([]TestCase, 0)

	scanner := bufio.NewScanner(file)
	testCasePattern := regexp.MustCompile(`(\d+)x(\d+): ([0-9 ]+)`)
	presentLinePattern := regexp.MustCompile(`[#.]{3}`)

	currentPresent := make([]string, 0)

	for scanner.Scan() {
		testCaseMatch := testCasePattern.FindStringSubmatch(scanner.Text())
		if testCaseMatch != nil {
			width, _ := strconv.Atoi(testCaseMatch[1])
			height, _ := strconv.Atoi(testCaseMatch[2])
			presents := []int{}

			for p := range strings.SplitSeq(testCaseMatch[3], " ") {
				present, _ := strconv.Atoi(p)
				presents = append(presents, present)
			}

			testCases = append(testCases, TestCase{ width, height, presents })
		}

		if presentLinePattern.MatchString(scanner.Text()) {
			currentPresent = append(currentPresent, strings.TrimSpace(scanner.Text()))
		}

		if len(strings.TrimSpace(scanner.Text())) == 0 && len(currentPresent) >= 3 {
			lines := make([][]bool, 3)
			count := 0
			for i := range 3 {
				line := make([]bool, 3)
				for j := range 3 {
					line[j] = currentPresent[i][j] == '#'
					if line[j] {
						count++
					}
				}
				lines[i] = line
			}

			presents = append(presents, Present{ lines, count })

			currentPresent = make([]string, 0)
		}
	}

	fmt.Println(presents)
	fmt.Println(testCases)
	fmt.Println("total count", len(testCases))

	immediateFailCount := 0
	immediatePassCount := 0
	for _, tc := range testCases {
		sum := 0
		pSum := 0
		for i, p := range presents {
			sum += p.tileCount * tc.presents[i]
			pSum += tc.presents[i]
		}
		if sum > tc.width * tc.height {
			immediateFailCount++
		}
		if pSum <= (tc.width / 3) * (tc.height / 3) {
			immediatePassCount++
		}
	}

	fmt.Println("immediate fail count", immediateFailCount)
	fmt.Println("immediate pass count", immediatePassCount)

	if immediateFailCount + immediatePassCount == len(testCases) {
		fmt.Println("this is everything")
	} else {
		fmt.Println("idk how to do the rest")
	}
}

