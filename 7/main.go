package main

import (
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	"github.com/jack-barr3tt/gostuff/types"
)

func buildCardMap(hand string) map[int][]string {
	cardMap := make(map[string]int)

	for _, card := range hand {
		cardMap[string(card)]++
	}

	cardMapRev := make(map[int][]string)

	for card, count := range cardMap {
		cardMapRev[count] = append(cardMapRev[count], card)
	}

	return cardMapRev
}

func handType(hand string) int {
	cardMap := buildCardMap(hand)

	if cardMap[5] != nil {
		return 7
	} else if cardMap[4] != nil {
		return 6
	} else if cardMap[3] != nil {
		if cardMap[2] != nil {
			return 5
		} else {
			return 4
		}
	} else if len(cardMap[2]) == 2 {
		return 3
	} else if len(cardMap[2]) == 1 {
		return 2
	}
	return 1
}

func cardValue(card byte) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

func cardValueNew(card byte) int {
	if card == 'J' {
		return 1
	}
	return cardValue(card)
}

func optimiseHand(hand string) string {
	jokerregex := regexp.MustCompile(`J`)

	jokers := jokerregex.FindAllString(hand, -1)

	// If there are no jokers, we can't optimise
	if len(jokers) == 0 {
		return hand
	}

	cardMap := buildCardMap(hand)

	// Look for as many matching cards as possible
	for i := 5; i > 0; i-- {
		if cardMap[i] == nil {
			continue
		}

		bestCard := byte('J')

		// Find the best card we can that occurs i times
		for _, card := range cardMap[i] {
			if cardValueNew(card[0]) > cardValueNew(bestCard) {
				bestCard = card[0]
			}
		}

		// If that best card is a J, we need to keep looking
		if bestCard != 'J' {
			return jokerregex.ReplaceAllString(hand, string(bestCard))
		}
	}

	// Can't optimise
	return hand
}

func sortHands(hands *[]types.Pair[string, int]) {
	h := *hands
	sort.Slice(h, func(i, j int) bool {
		t1 := handType(h[i].First)
		t2 := handType(h[j].First)

		if t1 == t2 {
			for n := range h[i].First {
				v1 := cardValue(h[i].First[n])
				v2 := cardValue(h[j].First[n])
				if v1 != v2 {
					return v1 > v2
				}
			}
		}
		return t1 > t2
	})
}

func sortNewHands(hands *[]types.Pair[types.Pair[string, string], int]) {
	h := *hands
	sort.Slice(h, func(i, j int) bool {
		t1 := handType(h[i].First.First)
		t2 := handType(h[j].First.First)

		if t1 == t2 {
			for n := range h[i].First.First {
				v1 := cardValueNew(h[i].First.Second[n])
				v2 := cardValueNew(h[j].First.Second[n])
				if v1 != v2 {
					return v1 > v2
				}
			}
		}
		return t1 > t2
	})
}

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	pairs := slicestuff.Map(func(line string) types.Pair[string, int] {
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		return types.Pair[string, int]{First: parts[0], Second: bid}
	}, lines)

	sortHands(&pairs)

	sum := 0

	for i, pair := range pairs {
		sum += pair.Second * (len(pairs) - i)
	}

	println("part 1:", sum)

	newPairs := slicestuff.Map(func(pair types.Pair[string, int]) types.Pair[types.Pair[string, string], int] {
		return types.Pair[types.Pair[string, string], int]{
			First: types.Pair[string, string]{
				First: optimiseHand(pair.First),
				// Preserve the original hand for sorting
				Second: pair.First},
			Second: pair.Second}
	}, pairs)

	sortNewHands(&newPairs)

	newSum := 0

	for i, pair := range newPairs {
		newSum += pair.Second * (len(newPairs) - i)
	}

	println("part 2:", newSum)
}
