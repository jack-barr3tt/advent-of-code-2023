package main

import (
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var numregex = regexp.MustCompile(`\d+`)

func hasPart(lines []string, lineno int, partrange []int) bool {
	for y := lineno - 1; y <= lineno+1; y++ {
		for x := partrange[0] - 1; x <= partrange[1]; x++ {
			// Out of range
			if x < 0 || y < 0 || y >= len(lines) || x >= len(lines[y]) {
				continue
			}
			// The part number itself
			if (x >= partrange[0] && x < partrange[1]) && y == lineno {
				continue
			}

			if lines[y][x] != '.' {
				return true
			}
		}
	}
	return false
}

// If an array has any consecutive numbers, keep only the first
func removeConsecutive(nums []int) []int {
	sort.Ints(nums[:])

	temp := []int{}
	filtered := []int{}

	for i := range nums {
		if len(temp) == 0 || temp[len(temp)-1] == nums[i]-1 {
			temp = append(temp, nums[i])
		} else {
			filtered = append(filtered, temp[0])
			temp = []int{}
			temp = append(temp, nums[i])
		}
	}
	if len(temp) > 0 {
		filtered = append(filtered, temp[0])
	}

	return filtered
}

func getRatio(lines []string, schematic []byte, lineno int, part int) int {
	// Get all indicies of numbers around the part
	indicies := []int{}

	for y := lineno - 1; y <= lineno+1; y++ {
		for x := part - 1; x <= part+1; x++ {
			// Out of range
			if x < 0 || y < 0 || y >= len(lines) || x >= len(lines[y]) {
				continue
			}
			// The part itself
			if x == part && y == lineno {
				continue
			}

			if numregex.MatchString(string(lines[y][x])) {
				indicies = append(indicies, x+y*(len(lines[y])+1))
			}
		}
	}

	// Remove consecutive indicies because these are the same number
	indicies = removeConsecutive(indicies)

	// Gear ratios are found from parts with exactly 2 adjacent part numbers
	if len(indicies) != 2 {
		return 0
	}

	ratio := 1

	// Find all part numbers in the schematic
	nums := numregex.FindAllIndex(schematic, -1)

	// Find part numbers that match indicies
	for _, index := range indicies {
		for _, num := range nums {
			if index >= num[0] && index < num[1] {
				// Multiply the ratio by the part number
				partnum, _ := strconv.Atoi(string(schematic[num[0]:num[1]]))
				ratio *= partnum
				break
			}
		}
	}

	return ratio
}

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	partregex := regexp.MustCompile(`\*`)

	partNoTotal := 0
	ratioTotal := 0

	for linenno, line := range lines {
		nums := numregex.FindAllIndex([]byte(line), -1)
		parts := partregex.FindAllIndex([]byte(line), -1)

		for _, part := range nums {
			if hasPart(lines, linenno, part) {
				partnum, _ := strconv.Atoi(line[part[0]:part[1]])
				partNoTotal += partnum
			}
		}

		for _, part := range parts {
			ratio := getRatio(lines, data, linenno, part[0])
			ratioTotal += ratio
		}
	}

	println(partNoTotal)
	println(ratioTotal)
}
