package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	data, err := fetchCaloriesList("input1.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	sortedElfCalories := getSortedElfCalories(data)

	elvesNum := len(sortedElfCalories)
	totalCals := 0
	fmt.Println("Top three elf calories:")
	for i := elvesNum - 1; i >= elvesNum-3; i-- {
		fmt.Printf("%d\n", sortedElfCalories[i])
		totalCals += sortedElfCalories[i]
	}
	fmt.Println("\nTotal calories of top three:")
	fmt.Printf("%d\n", totalCals)

}

// takes in a raw list of elf calories and returns a proceccesed sorted list with per elf calories
func getSortedElfCalories(input []int) []int {

	var elfCalories []int
	tempCals := 0
	for _, cals := range input {
		if cals != -1 {
			tempCals += cals
			continue
		}

		// We hit a blank line, add the calories up and reset
		elfCalories = append(elfCalories, tempCals)
		tempCals = 0

	}

	sort.Ints(elfCalories)
	return elfCalories
}

// parse the raw calories list and return an int array with seperators
func fetchCaloriesList(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			lines = append(lines, -1)
			continue
		}
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		lines = append(lines, line)
	}

	// append a terminator to the last group
	lines = append(lines, -1)

	return lines, scanner.Err()
}
