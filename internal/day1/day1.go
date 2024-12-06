package day1

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func Day1() {
	fmt.Println("Day 1")
	data, err := readJSONFile()
	if err != nil {
		fmt.Printf("Error with opening/reading JSON file: %w", err)
	}
	findDistance(data[0], data[1])
	findSimilarity(data[0], data[1])
}

func readJSONFile() ([][]int, error) {
	file, err := os.Open("internal/day1/input.json")
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

func findDistance(list1 []int, list2 []int) {
	// fmt.Printf("list1 has length of %v and list2 has length of %v \n", len(list1), len(list2))
	sort.Ints(list1)
	sort.Ints(list2)

	var difference []int
	for i := 0; i < len(list1); i++ {
		difference = append(difference, absInt(list1[i]-list2[i]))
	}

	var total int
	for _, value := range difference {
		total += value
	}

	fmt.Printf("The total difference of list1 and list2 is: %v \n", total)
}

func findSimilarity(list1 []int, list2 []int) {
	var similarities []int
	for _, val1 := range list1 {
		var same int = 0
		for _, val2 := range list2 {
			if val1 == val2 {
				same += 1
			}
		}
		similarities = append(similarities, same*val1)
	}

	var total int
	for _, val := range similarities {
		total += val
	}

	fmt.Printf("The similarity score is: %v \n", total)

}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
