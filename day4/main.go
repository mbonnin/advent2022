package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Assignment struct {
	min int
	max int
}

type Day []Assignment
type DayList []Day

// find assignments that have total overlap
func (d Day) findTotalOverlap() bool {
	// first assignment is contained in second
	if d[0].min >= d[1].min && d[0].max <= d[1].max {
		return true
	}
	// second assignment is contained in first
	if d[1].min >= d[0].min && d[1].max <= d[0].max {
		return true
	}
	return false
}

// find assignments that have any overlap
func (d Day) findAnyOverlap() bool {
	if d[0].min > d[1].max || d[0].max < d[1].min {
		return false
	}
	return true
}

// read in all of the days and initialize the assignments
func fetchInput(path string) (DayList, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var l DayList
	for scanner.Scan() {
		var a Assignment
		var d Day
		line := scanner.Text()
		assignments := strings.Split(line, ",")
		for _, assignment := range assignments {
			sections := strings.Split(assignment, "-")
			for range sections {
				a.min, _ = strconv.Atoi(sections[0])
				a.max, _ = strconv.Atoi(sections[1])
			}
			d = append(d, a)
		}
		l = append(l, d)
	}
	return l, scanner.Err()
}

func main() {
	days, err := fetchInput("input.txt")
	if err != nil {
		log.Fatalf("fetchInput: %s", err)
	}

	start := time.Now()

	var totalOverlap int
	var anyOverlap int
	for _, day := range days {
		if day.findAnyOverlap() {
			anyOverlap++
		}
		if day.findTotalOverlap() {
			totalOverlap++
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Total overlap: %d\n", totalOverlap)
	fmt.Printf("Any overlap: %d\n", anyOverlap)
	log.Printf("took %s", elapsed)
}
