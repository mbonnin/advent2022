package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type round struct {
	opponent     rune
	player       rune
	neededResult rune
}

func (r round) getPlayerPoints() int {
	switch r.player {

	case 'X': // rock
		return 1
	case 'Y': // paper
		return 2
	case 'Z': // scisors
		return 3
	}

	// error case
	return 0
}

// Determine the required points based on the "cheat code" and the opponent move
func (r *round) determinePlayerMove() {
	if r.neededResult == 'X' {
		r.player = r.getLose()
	} else if r.neededResult == 'Y' {
		r.player = r.getTie()
	} else if r.neededResult == 'Z' {
		r.player = r.getWin()
	}

}

func (r round) getWin() rune {
	switch r.opponent {
	case 'A': // rock
		return 'Y' // paper
	case 'B': // paper
		return 'Z' // scissors
	case 'C': // scissors
		return 'X' // rock
	}

	return ' '
}

func (r round) getTie() rune {
	switch r.opponent {
	case 'A': // rock
		return 'X' // rock
	case 'B': // paper
		return 'Y' // paper
	case 'C': // scissors
		return 'Z' // scissors
	}

	return ' '
}

func (r round) getLose() rune {
	switch r.opponent {
	case 'A': // rock
		return 'Z' // scissors
	case 'B': // paper
		return 'X' // rock
	case 'C': // scissors
		return 'Y' // paper
	}

	return ' '
}

// returns 0 for a loss, 3 for a tie, and 6 for a win
func (r round) getScore() int {
	var result int
	// player has rock
	if r.player == 'X' {
		if r.opponent == 'B' {
			result = 0
		} else if r.opponent == 'A' {
			result = 3
		} else if r.opponent == 'C' {
			result = 6
		}
	}
	// player has paper
	if r.player == 'Y' {
		if r.opponent == 'C' {
			result = 0
		} else if r.opponent == 'B' {
			result = 3
		} else if r.opponent == 'A' {
			result = 6
		}
	}
	// player has scissors
	if r.player == 'Z' {
		if r.opponent == 'A' {
			result = 0
		} else if r.opponent == 'C' {
			result = 3
		} else if r.opponent == 'B' {
			result = 6
		}
	}
	return result
}

func main() {
	data, err := fetchRounds("input1.txt")
	if err != nil {
		log.Fatalf("fetchRounds: %s", err)
	}

	finalScore := 0
	for _, r := range data {
		r.determinePlayerMove()
		finalScore += (r.getPlayerPoints() + r.getScore())
	}

	fmt.Printf("final score: %d\n", finalScore)
}

func fetchRounds(path string) ([]round, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rounds []round

	for scanner.Scan() {
		var r round
		runeArray := []rune(scanner.Text())
		r.opponent = runeArray[0]
		r.neededResult = runeArray[2]

		rounds = append(rounds, r)
	}

	return rounds, scanner.Err()
}
