package main

import (
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	// The + after the space is very important because of the formatting of the text file
	cardregex := regexp.MustCompile(`Card +\d+: `)
	numregex := regexp.MustCompile(`\d+`)

	totalPoints := 0

	cards := make(map[int]int)

	for _, line := range lines {
		// Extract game number
		card := cardregex.FindString(line)
		cardno, _ := strconv.Atoi(numregex.FindString(card))
		line = cardregex.ReplaceAllString(line, "")

		// Get parts of card
		parts := strings.Split(line, " | ")
		drawn := strings.Split(parts[0], " ")
		winning := strings.Split(parts[1], " ")

		matches := 0

		for _, card := range drawn {
			// Some "cards" will actually be blanks
			if card == "" {
				continue
			}

			// Initialise card if not already
			if cards[cardno] == 0 {
				cards[cardno] = 1
			}

			for _, win := range winning {
				if win == "" {
					continue
				}
				// Scoring
				if card == win {
					matches++
					break
				}
			}
		}
		// Integer conversion saves us from having to handle the zero matches case
		points := int(math.Pow(2, float64(matches-1)))
		totalPoints += points

		// Handle winning additional cards
		for i := 1; i <= matches; i++ {
			if cards[cardno+i] == 0 {
				cards[cardno+i] = 1 + cards[cardno]
			} else {
				cards[cardno+i] += cards[cardno]
			}
		}
	}

	// Add up how many cards we had by the end
	totalCards := 0
	for _, count := range cards {
		totalCards += count
	}

	println(totalPoints)
	println(totalCards)
}
