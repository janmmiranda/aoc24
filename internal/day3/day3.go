package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	fmt.Println("Day 3")
	input, err := readInput("internal/day3/input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v", err)
	}
	findMuls(input)
}

func findMuls(data string) {
	matches := applyRegexMul(data)
	fmt.Println(matches)
	inputs := applyRegexInner(matches)
	var result int

	for _, input := range inputs {
		res := input[0] * input[1]
		result += res
	}

	fmt.Printf("total: %v\n", result)
}

func readInput(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(content), nil
}

func applyRegexMul(data string) []string {
	// re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(data, -1)

	do := true
	var result []string
	for i, match := range matches {
		fmt.Printf("%v: %v\n", i, match)
		if match == "don't()" {
			do = false
		} else if match == "do()" {
			do = true
			continue
		}

		if do {
			result = append(result, match)
		}
	}

	return result
}

func applyRegexInner(dataList []string) [][]int {
	re := regexp.MustCompile(`\d{1,3},\d{1,3}`)
	var results [][]int
	for _, val := range dataList {
		var result []int
		matches := re.FindAllString(val, -1)
		// fmt.Printf("inner matches len: %v \n", len(matches))
		// fmt.Printf("found inner matches: %v\n", matches[0])
		for _, match := range strings.Split(string(matches[0]), ",") {
			num, err := strconv.Atoi(strings.TrimSpace(match))
			if err != nil {
				// fmt.Printf("error occured: %v\n", err)
				return nil
			}
			// fmt.Printf("%v converted to %v\n", match, num)
			result = append(result, num)
		}
		results = append(results, result)
	}

	return results
}
