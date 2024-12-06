package day2

import (
	"encoding/json"
	"fmt"
	"os"
)

var dampener = false

func Day2() {
	fmt.Println("Day 2")
	data, err := readJSONFile()
	if err != nil {
		fmt.Printf("Error with opening/reading JSON file: %w", err)
	}
	checkSafness(data) //part1 no dampener
	dampener = true
	checkSafness(data) //part 2 has dampener
}

func readJSONFile() ([][]int, error) {
	file, err := os.Open("internal/day2/input.json")
	// file, err := os.Open("internal/day2/inputTest.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var data [][]int
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return data, nil
}

func checkSafness(list [][]int) {
	var safeReports int

	for _, curList := range list {
		if validate(curList) || validateWithRemoval(curList) {
			safeReports += 1
			// fmt.Println("safe")
		} else {
			// fmt.Println("not safe")
			continue
		}
	}

	fmt.Printf("The total amount of safe reports is: %v \n", safeReports)
}

func validateWithRemoval(data []int) bool {
	if !dampener {
		// fmt.Println("has dampener")
		return false
	}
	for i := 0; i < len(data); i++ {
		data1 := append([]int{}, data[:i]...)
		data1 = append(data1, data[i+1:]...)
		if validate(data1) {
			return true
		}
	}
	return false
}

func validate(data []int) bool {
	desc := checkDecreasing(data)
	incr := checkIncreasing(data)
	maxVal := checkMaxDiff(data)
	minVal := checkMinDiff(data)
	// fmt.Printf("desc: %v, incr: %v, maxVal: %v, minVal: %v \n", desc, incr, maxVal, minVal)
	// fmt.Printf("data: %v \n", data)
	return (desc || incr) && maxVal && minVal
}

func checkDecreasing(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			return false
		}
	}
	return true
}

func checkIncreasing(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i-1] < list[i] {
			return false
		}
	}
	return true
}

func checkMaxDiff(list []int) bool {
	for i := 1; i < len(list); i++ {
		if absValue(list[i]-list[i-1]) > 3 {
			return false
		}
	}
	return true
}

func checkMinDiff(list []int) bool {
	for i := 1; i < len(list); i++ {
		if absValue(list[i]-list[i-1]) < 1 {
			return false
		}
	}
	return true
}

func absValue(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
