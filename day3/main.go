package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Sack struct {
	left        []rune
	right       []rune
	uniqueItems map[rune]bool
}

// given an initialized sack, map each of the items of the left sack and compare aganst the right item
// return the common item in each side
func (s Sack) findDuplicateItem() rune {
	// initialze a map of the items of the left
	itemMap := make(map[rune]bool)
	for _, item := range s.left {
		itemMap[item] = true
	}

	var returnVal rune

	// compare the left side to the right side
	for _, item := range s.right {
		if itemMap[item] {
			returnVal = item
		}
	}

	return returnVal
}

func (s *Sack) initUniqueItemsMap() {
	s.uniqueItems = make(map[rune]bool)

	for _, item := range s.left {
		s.uniqueItems[item] = true
	}
	for _, item := range s.right {
		s.uniqueItems[item] = true
	}
}

// takes a group of three rucksacks (a group) and returns the duplicate item between all three
func findBadge(group []Sack) rune {
	itemMap := make(map[rune]int)
	for _, sack := range group {
		for key, _ := range sack.uniqueItems {
			itemMap[key]++
			if itemMap[key] == 3 {
				return key
			}
		}
	}
	return '!'
}

func main() {
	data, err := fetchInput("input.txt")
	if err != nil {
		log.Fatalf("fetchInput: %s", err)
	}

	total := 0

	var group []Sack
	groupCounter := 0
	badgeTotal := 0

	for _, s := range data {
		s.initUniqueItemsMap()
		total += convertRuneToInt(s.findDuplicateItem())

		group = append(group, s)
		groupCounter++
		if groupCounter == 3 {
			badgeTotal += convertRuneToInt(findBadge(group))
			group = []Sack{}
			groupCounter = 0
		}

	}

	fmt.Printf("Total duplicate values: %d\n", total)
	fmt.Printf("Total badge values: %d\n", badgeTotal)
}

// takes a rune and converts it to an int
func convertRuneToInt(r rune) int {
	// lowercase
	if r >= 97 && r <= 123 {
		return int(r) - 96

	}
	// uppercase
	if r >= 65 && r <= 91 {
		return int(r) - 38
	}
	return -1
}

// read in all of the item lists and initialize the left and right sack items
func fetchInput(path string) ([]Sack, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sacks []Sack

	for scanner.Scan() {
		var s Sack
		runeArray := []rune(scanner.Text())
		s.left = runeArray[:len(runeArray)/2]
		s.right = runeArray[len(runeArray)/2:]

		sacks = append(sacks, s)
	}

	return sacks, scanner.Err()
}
